// @title cimpex
// @version 1.0
// @description cimpex is a CLI application written in Golang that gives the ability import and export docker images from a repository. GitHub repository at https://github.com/Mrpye/compex

// @contact.url https://github.com/Mrpye/cimpex

// @license.name Apache 2.0 licensed
// @license.url https://github.com/Mrpye/cimpex/blob/main/LICENSE

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
package main

import "github.com/Mrpye/cimpex/cmd"

func main() {
	cmd.Execute()
}
