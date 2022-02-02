package meta

import (
	"cli/libs"
	"fmt"
	"strings"
)

type InterfaceMeta struct {
	namespace   string
	useKeywords []string
	name        string
	parentClass string
	methods     []MethodMeta
}

func (meta InterfaceMeta) GetNamespace() string {
	return meta.namespace
}

func (meta *InterfaceMeta) SetNamespace(namespace string) {
	meta.namespace = namespace
}

func (meta *InterfaceMeta) SetUseKeywords(useKeywords []string) {
	meta.useKeywords = useKeywords
}

func (meta *InterfaceMeta) GetName() string {
	return meta.name
}

func (meta *InterfaceMeta) SetName(name string) {
	meta.name = name
}

func (meta *InterfaceMeta) SetMethods(methods []MethodMeta) {
	meta.methods = methods
}

func (meta InterfaceMeta) exportUseBlock() string {
	var using []string

	for i := 0; i < len(meta.useKeywords); i++ {
		trimed := strings.Trim(meta.useKeywords[i], "\\")
		tNamespace := strings.Split(trimed, "\\")
		if trimed != "" && meta.namespace != strings.Join(tNamespace[0:len(tNamespace)-1], "\\") {
			using = append(using, fmt.Sprintf("use %s;", trimed))
		}
		using = appendUsing(meta.namespace, meta.useKeywords[i], using)
	}

	for i := 0; i < len(meta.methods); i++ {
		using = appendUsing(meta.namespace, meta.methods[i].namespace, using)

		for j := 0; j < len(meta.methods[i].args); j++ {
			using = appendUsing(meta.namespace, meta.methods[i].args[j].namespace, using)
		}
	}

	using = libs.Unique(using)

	if len(using) == 0 {
		return ""
	} else {
		return "\n" + strings.Join(using, "\n") + "\n"
	}
}

func (meta InterfaceMeta) exportExtendsBlack() (extends string) {
	if meta.parentClass != "" {
		extends = " extends " + meta.parentClass
	}

	return extends
}

func (meta InterfaceMeta) exportMethodBlock() string {
	var methodList []string

	for i := 0; i < len(meta.methods); i++ {
		methodList = append(methodList, meta.createMethod(meta.methods[i]))
	}

	if len(methodList) == 0 {
		return ""
	} else {
		return "\n" + strings.Join(methodList, "\n")
	}
}

func (meta InterfaceMeta) createMethod(methodMeta MethodMeta) (method string) {
	methodType := methodMeta.returnType
	methodName := methodMeta.name
	methodArgs := methodMeta.args

	sp4 := libs.Indent(4)
	if methodType == "" {
		template := `%spublic function %s(%s);`
		method = fmt.Sprintf(template, sp4, methodName, argsToString(methodArgs))
	} else {
		template := `%spublic function %s(%s): %s;`
		method = fmt.Sprintf(template, sp4, methodName, argsToString(methodArgs), methodType)
	}

	return method
}

func (meta InterfaceMeta) ToString() string {
	template := `<?php

namespace %s;
%s
interface %s%s
{%s
}
`

	return fmt.Sprintf(template, meta.namespace, meta.exportUseBlock(), meta.name, meta.exportExtendsBlack(), meta.exportMethodBlock())
}
