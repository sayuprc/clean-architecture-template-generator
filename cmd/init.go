package cmd

import (
	"cli/libs"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var configTemplate = `{
	"directory": {
		"root": "."
	},
	"namespace": {
		"entity": "packages\\Domain\\Domain",
		"repository": "packages\\Infrastructure",
		"inmemoryrepository": "packages\\InMemoryInfrastructure",
		"usecase": "packages\\UseCase",
		"interactor": "packages\\Domain\\Application",
		"mockinteractor": "packages\\MockInteractor"
	}
}`

var initCmd = &cobra.Command{
	Use:   "init",
	Long:  "Create a configuration file",
	Short: "Create a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := "./.cagt.json"

		isForce, err := getParamBool(cmd, "force")

		if err != nil {
			fmt.Print(err)
			return
		}

		if !isForce {
			if libs.FileExists(filePath) {
				var yn string
				fmt.Print("Configuration file is already exists.\nOverwrite?: [y/N] ")
				fmt.Scanln(&yn)

				if strings.ToUpper(yn) != "Y" {
					return
				}
			}
		}

		fp, err := os.Create(filePath)

		if err != nil {
			fmt.Print(err)
			return
		}

		defer fp.Close()

		fp.WriteString(configTemplate)

		fmt.Printf("Created %s\n", filePath)
	},
}

func init() {
	initCmd.PersistentFlags().BoolP("force", "f", false, "Forcibly create a configuration file.")

	rootCmd.AddCommand(initCmd)
}
