/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package argon2

import (
	"echo-starter/internal/utils"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"

	"fmt"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var passwordsCmd = &cobra.Command{
	Use:   "argon2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		type output struct {
			Secret string
			Hash   string
		}

		hash, _ := utils.GeneratePasswordHash(secret)
		fmt.Println(core_utils.PrettyJSON(output{
			Secret: secret,
			Hash:   hash,
		}))
	},
}
var secret string

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(passwordsCmd)

	passwordsCmd.Flags().StringVarP(&secret, "secret", "s", utils.GeneratePassword(), "--secret=mypassword")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
