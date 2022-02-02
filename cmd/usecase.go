package cmd

import (
	"cli/config"
	"cli/libs"
	"cli/meta"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var usecaseCmd = &cobra.Command{
	Use:   "usecase",
	Short: "Generate a usecase and interactor and input/output data struct",
	Run: func(cmd *cobra.Command, args []string) {
		config, configErr := config.Load()

		if configErr != nil {
			fmt.Print(configErr)
			return
		}

		domainName, dParamErr := getParamString(cmd, "domain")
		if dParamErr != nil {
			fmt.Print(dParamErr)
			return
		}

		useCaseName, uParamErr := getParamString(cmd, "usecase")
		if uParamErr != nil {
			fmt.Print(uParamErr)
			return
		}

		if domainName == "" {
			domainName = libs.AskedRepetition("Domain Name:")
			domainName = strings.Title(domainName)
		} else {
			domainName = strings.Title(domainName)
		}

		if useCaseName == "" {
			useCaseName = libs.AskedRepetition("UseCase Name:")
			useCaseName = strings.Title(useCaseName)
		} else {
			useCaseName = strings.Title(useCaseName)
		}

		rootDir := config.Directory.Root

		metas := []meta.MetaInterface{}

		useCaseNamespace := config.Namespace.UseCase + "\\" + domainName + "\\" + useCaseName

		libs.EchoSection("Input Data Struct Property Settings")
		inputDataStructProps := propertyReader()
		libs.EchoSection("Input Data Struct Method Settings")
		inputDataStructMethods := methodReader()

		inputDataStructMeta := &meta.ClassMeta{}
		inputDataStructMeta.SetNamespace(useCaseNamespace)
		inputDataStructName := domainName + useCaseName + "Request"
		inputDataStructMeta.SetName(inputDataStructName)
		inputDataStructMeta.SetProperties(inputDataStructProps)
		inputDataStructMeta.SetMethods(inputDataStructMethods)

		libs.EchoSection("Output Data Struct Property Settings")
		outputDataStructProps := propertyReader()
		libs.EchoSection("Output Data Struct Method Settings")
		outputDataStructMethods := methodReader()
		outputDataStructMeta := &meta.ClassMeta{}
		outputDataStructMeta.SetNamespace(useCaseNamespace)
		outputDataStructName := domainName + useCaseName + "Response"
		outputDataStructMeta.SetName(outputDataStructName)
		outputDataStructMeta.SetProperties(outputDataStructProps)
		outputDataStructMeta.SetMethods(outputDataStructMethods)

		useCaseMeta := &meta.InterfaceMeta{}
		useCaseMeta.SetNamespace(useCaseNamespace)
		useCaseMeta.SetName(domainName + useCaseName + "UseCaseInterface")
		methodArg := meta.ArgMeta{}
		methodArg.SetArgType(inputDataStructName)
		methodArg.SetName("request")
		method := meta.MethodMeta{}
		method.SetName("handle")
		method.SetReturnType(outputDataStructName)
		method.SetArgs([]meta.ArgMeta{methodArg})
		useCaseMeta.SetMethods([]meta.MethodMeta{method})

		interactorMeta := &meta.ClassMeta{}
		interactorMeta.SetNamespace(config.Namespace.Interactor + "\\" + domainName)
		interactorMeta.SetName(domainName + useCaseName + "Interactor")
		interactorMeta.SetInterfaces([]meta.InterfaceMeta{*useCaseMeta})

		mockInteractorMeta := &meta.ClassMeta{}
		mockInteractorMeta.SetNamespace(config.Namespace.MockInteractor + "\\" + domainName)
		mockInteractorMeta.SetName("Mock" + domainName + useCaseName + "Interactor")
		mockInteractorMeta.SetInterfaces([]meta.InterfaceMeta{*useCaseMeta})

		metas = append(metas, inputDataStructMeta)
		metas = append(metas, outputDataStructMeta)
		metas = append(metas, useCaseMeta)
		metas = append(metas, interactorMeta)
		metas = append(metas, mockInteractorMeta)

		libs.EchoSingleLine()
		libs.EchoNewLine()

		build(metas, rootDir)

		libs.EchoNewLine()
	},
}

func init() {
	usecaseCmd.PersistentFlags().StringP("domain", "d", "", "Input domain name")
	usecaseCmd.PersistentFlags().StringP("usecase", "u", "", "Input usecase name")

	rootCmd.AddCommand(usecaseCmd)
}
