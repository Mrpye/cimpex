/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/Mrpye/cimpex/modules/registry"
	"github.com/spf13/cobra"
)

func Cmd_Export() *cobra.Command {
	var target_User string
	var target_Password string
	var target_IgnoreSSL bool
	var target_SaveName string
	// exportCmd represents the export command
	var cmd = &cobra.Command{
		Use:   "export [target]",
		Short: "Exports a container image from a registry",
		Long:  `Exports a container image from a registry`,
		RunE: func(cmd *cobra.Command, args []string) error {
			//****************
			//Validate Command
			//****************
			if len(args) < 1 {
				return errors.New("missing target")
			}
			reg := registry.CreateDockerRegistry(target_User, target_Password, target_IgnoreSSL)
			reg.Download(args[0], target_SaveName)
			fmt.Println("export called")
			return nil
		},
	}
	cmd.Flags().StringVarP(&target_SaveName, "save", "s", "", "Tar Save Path")
	cmd.Flags().StringVarP(&target_User, "user", "u", "", "Registry User")
	cmd.Flags().StringVarP(&target_Password, "password", "p", "", "Registry User")
	cmd.Flags().BoolVarP(&target_IgnoreSSL, "ignore_ssl", "i", false, "Ignore SSL")
	return cmd
}
func init() {
	rootCmd.AddCommand(Cmd_Export())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
