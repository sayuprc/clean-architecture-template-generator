package meta

import (
	"fmt"
	"strings"
)

type MetaInterface interface {
	ToString() string
	GetName() string
	GetNamespace() string
}

func appendUsing(namespace, inputs string, usings []string) []string {
	trimedNamespace := strings.Trim(inputs, "\\")
	splitedNamespace := strings.Split(trimedNamespace, "\\")

	if trimedNamespace != "" && namespace != strings.Join(splitedNamespace[0:len(splitedNamespace)-1], "\\") {
		usings = append(usings, fmt.Sprintf("use %s;", trimedNamespace))
	}

	return usings
}

func argsToString(argMetas []ArgMeta) string {
	var args []string

	for i := 0; i < len(argMetas); i++ {
		propType := argMetas[i].argType
		name := argMetas[i].name
		if propType == "" {
			args = append(args, fmt.Sprintf("$%s", name))
		} else {
			args = append(args, fmt.Sprintf("%s $%s", propType, name))
		}
	}

	return strings.Join(args, ", ")
}
