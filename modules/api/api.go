package api

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	_ "github.com/Mrpye/cimpex/docs"
	"github.com/Mrpye/cimpex/modules/body_types"
	"github.com/Mrpye/cimpex/modules/helper"
	"github.com/Mrpye/cimpex/modules/registry"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var web_base_folder string
var web_ip string
var web_port string

// @Summary Import Docker Image to Registry from tar file
// @ID post-import-docker-image
// @Produce json
// @Param request body body_types.ImportExportRequest.request true true "query params"
// @Success 200 {string}  string "tar file imported"
// @Failure 404 {string}  string "error"
// @Router /import [post]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func postImport(c *gin.Context) {
	var importRequest body_types.ImportExportRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	dec := helper.DecodeB64(token)
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

// @Summary Export Docker Image from Registry to tar file
// @ID post-export-docker-image
// @Produce json
// @Param request body body_types.ImportExportRequest.request true true "query params"
// @Success 200 {string}  string "tar file Exported"
// @Failure 404 {string}  string "error"
// @Router /export [post]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func postExport(c *gin.Context) {
	var importRequest body_types.ImportExportRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	user := ""
	password := ""
	if val, ok := c.Request.Header["Authorization"]; ok {
		token := strings.Split(val[0], " ")[1]
		dec := helper.DecodeB64(token)
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

// @Summary Exports Docker Images from Registry to tar file
// @ID post-exports-docker-images
// @Produce json
// @Param request body []body_types.ImportExportRequest.request true true "query params"
// @Success 200 {string}  string "tar file Exported"
// @Failure 404 {string}  string "error"
// @Router /exports [post]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func postExports(c *gin.Context) {
	var importRequest []body_types.ImportExportRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	user := ""
	password := ""
	if val, ok := c.Request.Header["Authorization"]; ok {
		token := strings.Split(val[0], " ")[1]
		dec := helper.DecodeB64(token)
		parts := strings.Split(dec, ":")
		if parts[0] == "" || parts[1] == "" {
			c.IndentedJSON(http.StatusUnauthorized, "Missing username or password")
			return
		}
		user = parts[0]
		password = parts[1]
	}

	reg := registry.CreateDockerRegistry(user, password, importRequest[0].IgnoreSSL)
	var results []string
	for _, o := range importRequest {
		reg.IgnoreSSL = o.IgnoreSSL
		err := reg.Download(o.Target, path.Join(web_base_folder, o.TarFile))
		if err != nil {
			c.IndentedJSON(400, err)
			return
		} else {
			results = append(results, o.TarFile+" Exported")
		}
	}

	c.IndentedJSON(http.StatusCreated, results)

}

// @Summary Check API Endpoint
// @ID check-api-endpoint
// @Produce json
// @Success 200 {string}  string "ok"
// @Router / [get]
func getOK(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "OK")
}

// @Summary List the docker images tar files in the directory
// @ID post-list-docker-images
// @Produce json
// @Success 200 {object}  []body_types.PackageInfo.response
// @Failure 404 {string}  string "error"
// @Router /list [post]
func postListImages(c *gin.Context) {

	reg := registry.CreateDockerRegistry("", "", false)

	files, err := helper.WalkMatch(web_base_folder, "*.tar")
	if err != nil {
		c.IndentedJSON(http.StatusCreated, err.Error())
		return
	}
	var results []body_types.PackageInfo

	for _, o := range files {
		name_tag, _ := reg.GetImageNameTag(path.Join(web_base_folder, o))
		if name_tag != "" {
			results = append(results, body_types.PackageInfo{TarPath: o, ImageName: name_tag})
		}
	}

	c.IndentedJSON(http.StatusCreated, results)

}

// Function to start web server
func StartWebServer(ip string, port string, base_folder string) {
	//****************
	//Set the variable
	//****************
	web_ip = ip
	web_port = port
	web_base_folder = base_folder

	fmt.Println("Starting Web-Server")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/import", postImport)
	router.POST("/export", postExport)
	router.POST("/exports", postExports)
	router.POST("/list", postListImages)
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

}
