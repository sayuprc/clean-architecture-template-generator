package cmd

import (
	"cli/config"
	"cli/libs"
	"cli/meta"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	typeMap = map[string]string{
		"1":  "string",
		"2":  "int",
		"3":  "float",
		"4":  "bool",
		"5":  "array",
		"6":  "other type",
		"99": "end setting",
	}

	returnTypeMap = map[string]string{
		"1":  "string",
		"2":  "int",
		"3":  "float",
		"4":  "bool",
		"5":  "array",
		"6":  "void",
		"7":  "other type",
		"99": "end setting",
	}
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Generate a entity and repository based on input values",
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

		entityName, eParamErr := getParamString(cmd, "entity")
		if eParamErr != nil {
			fmt.Print(eParamErr)
			return
		}

		if domainName == "" {
			domainName = libs.AskedRepetition("Domain Name:")
			domainName = strings.Title(domainName)
		} else {
			domainName = strings.Title(domainName)
		}

		if entityName == "" {
			entityName = libs.AskedRepetition("Entity Name:")
			entityName = strings.Title(entityName)
		} else {
			entityName = strings.Title(entityName)
		}

		rootDir := config.Directory.Root

		metas := []meta.MetaInterface{}

		entityMeta := &meta.ClassMeta{}

		entityNamespace := config.Namespace.Entity + "\\" + domainName
		entityMeta.SetName(entityName)
		entityMeta.SetNamespace(entityNamespace)

		libs.EchoSection("Entity Property Settings")
		entityProps := propertyReader()
		libs.EchoSection("Entity Method Settings")
		entityMethods := methodReader()

		entityMeta.SetProperties(entityProps)
		entityMeta.SetMethods(entityMethods)

		metas = append(metas, entityMeta)

		if ans := libs.Ask("Generate Repository? [y/N]:"); strings.ToUpper(ans) == "Y" {
			repositoryInterfaceMeta := &meta.InterfaceMeta{}
			repositoryMeta := &meta.ClassMeta{}
			inMemoryRepositoryMeta := &meta.ClassMeta{}

			repositoryInterfaceMeta.SetName(entityName + "RepositoryInterface")
			repositoryInterfaceMeta.SetNamespace(entityNamespace)

			repositoryNamespace := config.Namespace.Repository + "\\" + domainName
			repositoryMeta.SetName(entityName + "Repository")
			repositoryMeta.SetNamespace(repositoryNamespace)
			repositoryMeta.SetInterfaces([]meta.InterfaceMeta{*repositoryInterfaceMeta})

			inMemoryRepositoryNamespace := config.Namespace.InMemoryRepository + "\\" + domainName
			inMemoryRepositoryMeta.SetName("InMemory" + entityName + "Repository")
			inMemoryRepositoryMeta.SetNamespace(inMemoryRepositoryNamespace)
			inMemoryRepositoryMeta.SetInterfaces([]meta.InterfaceMeta{*repositoryInterfaceMeta})

			repositoryMethods := methodReader()

			repositoryInterfaceMeta.SetMethods(repositoryMethods)
			repositoryMeta.SetMethods(repositoryMethods)
			inMemoryRepositoryMeta.SetMethods(repositoryMethods)

			metas = append(metas, repositoryInterfaceMeta)
			metas = append(metas, repositoryMeta)
			metas = append(metas, inMemoryRepositoryMeta)
		}

		libs.EchoSingleLine()
		libs.EchoNewLine()

		build(metas, rootDir)

		libs.EchoNewLine()
	},
}

func init() {
	entityCmd.PersistentFlags().StringP("domain", "d", "", "Input domain name")
	entityCmd.PersistentFlags().StringP("entity", "e", "", "Input entity name")

	rootCmd.AddCommand(entityCmd)
}

func build(metas []meta.MetaInterface, rootDir string) {
	for _, v := range metas {
		if err := make(v, rootDir); err != nil {
			fmt.Print(err)
		}
	}
}

func make(meta meta.MetaInterface, rootDir string) error {
	directory := strings.Trim(rootDir, "/") + "/" + libs.ConvertToDirectory(meta.GetNamespace())

	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}

	fileName := directory + "/" + meta.GetName() + ".php"

	fp, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer fp.Close()

	fp.WriteString(meta.ToString())

	fmt.Printf("Created %s\n", fileName)

	return nil
}

func propertyReader() (entityProps []meta.PropertyMeta) {
	propertyMapForDisplay := libs.ConvertMapForDisplay(typeMap)

	displayText := "Select the number of property type\n\n"
	for _, k := range propertyMapForDisplay {
		displayText += fmt.Sprintf("[%s] %s\n", k, typeMap[k])
	}
	displayText += "\nType number:"

	for {
		propertyTypeNumber := libs.AskedRepetition(displayText)

		if propertyTypeNumber == "99" {
			break
		}

		propertyName := libs.AskedRepetition("Property Name:")

		var propertyType string
		var propertyNamespace string

		if propertyTypeNumber == "6" {
			propertyType = libs.AskedRepetition("Property Type:")

			propertyNamespace = libs.Ask("Property Namespace:")
		} else {
			propertyType = typeMap[propertyTypeNumber]
		}

		tmpProperty := &meta.PropertyMeta{}
		tmpProperty.SetName(propertyName)
		tmpProperty.SetPropertyType(propertyType)
		if propertyNamespace != "" {
			tmpProperty.SetNamespace(propertyNamespace + "\\" + propertyType)
		}

		entityProps = append(entityProps, *tmpProperty)

		libs.EchoNewLine()
	}

	return entityProps
}

func methodReader() (methods []meta.MethodMeta) {
	returnTypeMapForDisplay := libs.ConvertMapForDisplay(returnTypeMap)

	typeMapForDisplay := libs.ConvertMapForDisplay(typeMap)

	typeDisplayText := "Select the number of argument type\n\n"
	for _, k := range typeMapForDisplay {
		typeDisplayText += fmt.Sprintf("[%s] %s\n", k, typeMap[k])
	}
	typeDisplayText += "\nType number:"

	returnTypeDisplayText := "Select the number of method return type\n\n"
	for _, k := range returnTypeMapForDisplay {
		returnTypeDisplayText += fmt.Sprintf("[%s] %s\n", k, returnTypeMap[k])
	}
	returnTypeDisplayText += "\nType number:"

	for {
		returnTypeNumber := libs.AskedRepetition(returnTypeDisplayText)

		if returnTypeNumber == "99" {
			break
		}

		methodName := libs.AskedRepetition("Method Name: ")

		var returnType string
		var returnTypeNamespace string
		if returnTypeNumber == "7" {
			returnType = libs.AskedRepetition("Return Type Name:")

			returnTypeNamespace = libs.Ask("Return Type Namespace:")
		} else {
			returnType = returnTypeMap[returnTypeNumber]
		}

		var args []meta.ArgMeta

		for {
			argTypeNumber := libs.AskedRepetition(typeDisplayText)

			if argTypeNumber == "99" {
				break
			}

			argName := libs.AskedRepetition("Argument Name:")

			var argType string
			var argNamespace string

			if argTypeNumber == "6" {
				argType = libs.AskedRepetition("Argument Type:")

				argNamespace = libs.Ask("Argument Namespace:")
			} else {
				argType = typeMap[argTypeNumber]
			}

			tmpArg := &meta.ArgMeta{}
			tmpArg.SetArgType(argType)
			tmpArg.SetName(argName)
			if argNamespace != "" {
				tmpArg.SetNamespace(argNamespace + "\\" + argType)
			}
			args = append(args, *tmpArg)
		}

		tmpMethod := &meta.MethodMeta{}

		tmpMethod.SetName(methodName)
		tmpMethod.SetReturnType(returnType)
		if returnTypeNamespace != "" {
			tmpMethod.SetNamespace(returnTypeNamespace + "\\" + returnType)
		}
		tmpMethod.SetArgs(args)

		methods = append(methods, *tmpMethod)
	}

	return methods
}
