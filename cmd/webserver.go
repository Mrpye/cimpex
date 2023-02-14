/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Mrpye/cimpex/modules/api"
	"github.com/spf13/cobra"
)

func Cmd_WebServer() *cobra.Command {
	var web_port string
	var web_ip string
	var web_base_folder string

	var cmd = &cobra.Command{
		Use:   "web",
		Short: "Start a API Web-Server",
		Long:  `Start a API Web-Server`,
		Run: func(cmd *cobra.Command, args []string) {
			api.StartWebServer(web_ip, web_port, web_base_folder)
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
