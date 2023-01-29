package registry

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/layout"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type DockerRegistry struct {
	Host      string `json:"host" yaml:"host"`
	User      string `json:"user" yaml:"user"`
	Password  string `json:"password" yaml:"password"`
	IgnoreSSL bool   `json:"ignore_ssl" yaml:"ignore_ssl"`
}

//Function to create the DockerRegistry struct
func CreateDockerRegistry(User string, Password string, IgnoreSSL bool) *DockerRegistry {
	return &DockerRegistry{
		User:      User,
		Password:  Password,
		IgnoreSSL: IgnoreSSL,
	}
}

func (c *DockerRegistry) loadImage(path string, index bool) (partial.WithRawManifest, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		img, err := crane.Load(path)
		if err != nil {
			return nil, fmt.Errorf("loading %s as tarball: %w", path, err)
		}
		return img, nil
	}

	l, err := layout.ImageIndexFromPath(path)
	if err != nil {
		return nil, fmt.Errorf("loading %s as OCI layout: %w", path, err)
	}

	if index {
		return l, nil
	}

	m, err := l.IndexManifest()
	if err != nil {
		return nil, err
	}
	if len(m.Manifests) != 1 {
		return nil, fmt.Errorf("layout contains %d entries, consider --index", len(m.Manifests))
	}

	desc := m.Manifests[0]
	if desc.MediaType.IsImage() {
		return l.Image(desc.Digest)
	} else if desc.MediaType.IsIndex() {
		return l.ImageIndex(desc.Digest)
	}

	return nil, fmt.Errorf("layout contains non-image (mediaType: %q), consider --index", desc.MediaType)
}

//Upload the image to docker Registry
func (m *DockerRegistry) Upload(config string, image_path string) error {

	basicAuthn := &authn.Basic{
		Username: m.User,
		Password: m.Password,
	}

	var options []remote.Option
	if m.User != "" && m.Password != "" {
		withAuthOption := remote.WithAuth(basicAuthn)
		//--insecure
		options = []remote.Option{withAuthOption}
	} else {
		options = []remote.Option{}
	}

	//**********************
	//Check if to ignore ssl
	//**********************
	if m.IgnoreSSL {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		options = append(options, remote.WithTransport(transport))
	}
	fmt.Printf("Importing Tar %s to %s \n", image_path, config)
	img, err := m.loadImage(image_path, false)
	if err != nil {
		return err
	}

	ref, err := name.ParseReference(config)
	if err != nil {
		return err
	}
	//var h v1.Hash
	switch t := img.(type) {
	case v1.Image:
		if err := remote.Write(ref, t, options...); err != nil {
			return err
		}
		if _, err = t.Digest(); err != nil {
			return err
		}
	case v1.ImageIndex:
		if err := remote.WriteIndex(ref, t, options...); err != nil {
			return err
		}
		if _, err = t.Digest(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot push type (%T) to registry", img)
	}
	fmt.Printf("Finished Importing Tar %s to %s\n", image_path, config)
	return nil
}

//Upload the image from the docker Registry
func (m *DockerRegistry) Download(config string, save_name string) error {

	basicAuthn := &authn.Basic{
		Username: m.User,
		Password: m.Password,
	}

	imageMap := map[string]v1.Image{}
	var options []remote.Option
	if m.User != "" && m.Password != "" {
		withAuthOption := remote.WithAuth(basicAuthn)
		options = []remote.Option{withAuthOption}
	} else {
		options = []remote.Option{}
	}
	//**********************
	//Check if to ignore ssl
	//**********************
	if m.IgnoreSSL {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		options = append(options, remote.WithTransport(transport))

	}

	imageName := config

	ref, err := name.ParseReference(imageName)
	if err != nil {
		log.Fatalf("cannot parse reference of the image %s , detail: %v", imageName, err)
	}

	rmt, err := remote.Get(ref, options...)
	if err != nil {
		return err
	}

	img, err := rmt.Image()
	if err != nil {
		return err
	}

	imageMap[imageName] = img

	//**************
	//Save the image
	//**************
	full_img_name := ref.Context().Name()
	parts := strings.Split(full_img_name, "/")
	if len(parts) < 1 {
		return errors.New("error parsing image name")
	}
	img_name := parts[len(parts)-1]
	tag := "latest"
	if len(parts) > 3 {
		tag = parts[3]
	}
	full_path := img_name + "-" + tag + ".tar"
	if save_name != "" {
		full_path = save_name
	}

	file := path.Dir(full_path)
	err = os.MkdirAll(file, os.ModePerm)
	if err != nil {
		return err
	}

	for k := range imageMap {
		fmt.Printf("Downloading Image %v Please Wait!\n", k)
	}
	if err := crane.MultiSave(imageMap, full_path); err != nil {
		return fmt.Errorf("saving tarball %s: %w", full_path, err)
	}
	for k := range imageMap {
		fmt.Printf("Finished Downloading Image %v \n", k)
	}

	return nil
}
