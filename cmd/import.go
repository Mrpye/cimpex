package cmd

import (
	"errors"
	"fmt"

	"github.com/Mrpye/cimpex/modules/registry"
	"github.com/spf13/cobra"
)

func Cmd_Import() *cobra.Command {
	var target_User string
	var target_Password string
	var target_IgnoreSSL bool
	var cmd = &cobra.Command{
		Use:   "import [tar] [target]",
		Short: "Imports a container image into a registry",
		Long:  `Imports a container image into a registry`,
		RunE: func(cmd *cobra.Command, args []string) error {
			//****************
			//Validate Command
			//****************
			if len(args) < 2 {
				return errors.New("missing tar and target")
			}
			reg := registry.CreateDockerRegistry(target_User, target_Password, target_IgnoreSSL)

			reg.Upload(args[0], args[1])
			fmt.Println("import called")
			return nil
		},
	}
	cmd.Flags().StringVarP(&target_User, "user", "u", "", "Registry User")
	cmd.Flags().StringVarP(&target_Password, "password", "p", "", "Registry User")
	cmd.Flags().BoolVarP(&target_IgnoreSSL, "ignore_ssl", "i", false, "Ignore SSL")
	return cmd
}

func init() {
	rootCmd.AddCommand(Cmd_Import())
}
