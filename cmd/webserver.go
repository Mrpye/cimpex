/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Mrpye/cimpex/registry"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var web_base_folder string

type ImportExportRequest struct {
	TarFile   string `json:"tar"`
	Target    string `json:"target"`
	IgnoreSSL bool   `json:"ignore_ssl"`
}

func DecodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(base64Text, []byte(message))
	fmt.Printf("base64: %s\n", base64Text)
	return string(base64Text)
}

func postImport(c *gin.Context) {
	var importRequest ImportExportRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	dec := DecodeB64(token)
	parts := strings.Split(dec, ":")
	if parts[0] == "" || parts[1] == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Missing username or password")
		return
	}

	reg := registry.CreateDockerRegistry(parts[0], parts[1], importRequest.IgnoreSSL)
	err := reg.Upload(importRequest.Target, path.Join(web_base_folder, importRequest.TarFile))
	if err != nil {
		c.IndentedJSON(400, err)
	} else {
		c.IndentedJSON(http.StatusCreated, importRequest.TarFile+" Imported")
	}

}

func postExport(c *gin.Context) {
	var importRequest ImportExportRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	user := ""
	password := ""
	if val, ok := c.Request.Header["Authorization"]; ok {
		token := strings.Split(val[0], " ")[1]
		dec := DecodeB64(token)
		parts := strings.Split(dec, ":")
		if parts[0] == "" || parts[1] == "" {
			c.IndentedJSON(http.StatusUnauthorized, "Missing username or password")
			return
		}
		user = parts[0]
		password = parts[1]
	}

	reg := registry.CreateDockerRegistry(user, password, importRequest.IgnoreSSL)
	err := reg.Download(importRequest.Target, path.Join(web_base_folder, importRequest.TarFile))
	if err != nil {
		c.IndentedJSON(400, err)
	} else {
		c.IndentedJSON(http.StatusCreated, importRequest.TarFile+" Exported")
	}

}

func getOK(c *gin.Context) {

	c.IndentedJSON(http.StatusCreated, "OK")

}

func Cmd_WebServer() *cobra.Command {
	var web_port string
	var web_ip string

	var cmd = &cobra.Command{
		Use:   "web",
		Short: "Start a API Web-Server",
		Long:  `Start a API Web-Server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Web-Server")
			gin.SetMode(gin.ReleaseMode)
			router := gin.Default()

			router.POST("/import", postImport)
			router.POST("/export", postExport)
			router.GET("/", getOK)

			//**********************************
			//Set up the environmental variables
			//**********************************
			if os.Getenv("WEB_IP") != "" {
				web_ip = os.Getenv("WEB_IP")
			}
			if os.Getenv("WEB_PORT") != "" {
				web_port = os.Getenv("WEB_PORT")
			}
			if os.Getenv("BASE_FOLDER") != "" {
				web_base_folder = os.Getenv("BASE_FOLDER")
			}

			router.Run(web_ip + ":" + web_port)
		},
	}
	cmd.Flags().StringVarP(&web_port, "port", "p", "8080", "Listen on Port")
	cmd.Flags().StringVarP(&web_ip, "ip", "i", "localhost", "Listen on Ip")
	cmd.Flags().StringVarP(&web_base_folder, "folder", "f", "", "base export import folder")

	return cmd
}
func init() {
	rootCmd.AddCommand(Cmd_WebServer())
}
