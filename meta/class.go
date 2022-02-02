package meta

import (
	"cli/libs"
	"fmt"
	"strings"
)

type ClassMeta struct {
	namespace   string
	useKeywords []string
	name        string
	parentClass string
	interfaces  []InterfaceMeta
	properties  []PropertyMeta
	methods     []MethodMeta
}

func (meta ClassMeta) GetNamespace() string {
	return meta.namespace
}

func (meta *ClassMeta) SetNamespace(namespace string) {
	meta.namespace = namespace
}

func (meta *ClassMeta) SetUseKeywords(useKeywords []string) {
	meta.useKeywords = useKeywords
}

func (meta ClassMeta) GetName() string {
	return meta.name
}

func (meta *ClassMeta) SetName(name string) {
	meta.name = name
}

func (meta *ClassMeta) SetInterfaces(interfaces []InterfaceMeta) {
	meta.interfaces = interfaces
}

func (meta *ClassMeta) SetProperties(properties []PropertyMeta) {
	meta.properties = properties
}

func (meta *ClassMeta) SetMethods(methods []MethodMeta) {
	meta.methods = methods
}

func (meta ClassMeta) exportUseBlock() string {
	var using []string

	for _, v := range meta.useKeywords {
		using = appendUsing(meta.namespace, v, using)
	}

	for _, v := range meta.interfaces {
		using = appendUsing(meta.namespace, v.namespace+"\\"+v.name, using)
	}

	for _, v := range meta.properties {
		using = appendUsing(meta.namespace, v.namespace, using)
	}

	for _, v := range meta.methods {
		using = appendUsing(meta.namespace, v.namespace, using)

		for _, a := range v.args {
			using = appendUsing(meta.namespace, a.namespace, using)
		}
	}

	using = libs.Unique(using)

	if len(using) == 0 {
		return ""
	} else {
		return "\n" + strings.Join(using, "\n") + "\n"
	}
}

func (meta ClassMeta) exportExtendsBlack() (extends string) {
	if meta.parentClass != "" {
		extends = " extends " + meta.parentClass
	}

	return extends
}

func (meta ClassMeta) exportImplementsBlock() (implements string) {
	for _, v := range meta.interfaces {
		implements += v.name + ","
	}

	if implements == "" {
		return implements
	} else {
		return " implements " + strings.Trim(implements, ",")
	}
}

func (meta ClassMeta) exportPropertyBlock() (fieldProps string) {
	sp4 := libs.Indent(4)

	for _, v := range meta.properties {
		propType := v.propertyType
		name := v.name

		if propType == "" {
			fieldProps += fmt.Sprintf("%sprivate $%s;\n", sp4, name)
		} else {
			fieldProps += fmt.Sprintf("%sprivate %s $%s;\n", sp4, propType, name)
		}
	}

	if len(fieldProps) == 0 {
		return fieldProps
	} else {
		return "\n" + fieldProps
	}
}

func (meta ClassMeta) exportMethodBlock() string {
	var methodList []string

	methodList = append(methodList, meta.createConstructor())

	if len(meta.properties) != 0 {
		methodList = append(methodList, meta.createGetter())
	}

	for _, v := range meta.methods {
		methodList = append(methodList, meta.createMethod(v))
	}

	if len(methodList) == 0 {
		return ""
	} else {
		return "\n" + strings.Join(methodList, "\n\n")
	}
}

func (meta ClassMeta) createMethod(methodMeta MethodMeta) (method string) {
	methodType := methodMeta.returnType
	methodName := methodMeta.name
	methodArgs := methodMeta.args

	sp4 := libs.Indent(4)

	if methodType == "" {
		template := `%spublic function %s(%s)
%s{
%s}`
		method = fmt.Sprintf(template, sp4, methodName, argsToString(methodArgs), sp4, sp4)
	} else {
		template := `%spublic function %s(%s): %s
%s{
%s}`
		method = fmt.Sprintf(template, sp4, methodName, argsToString(methodArgs), methodType, sp4, sp4)
	}

	return method
}

func (classMeta ClassMeta) createConstructor() string {
	var method string

	sp4 := libs.Indent(4)

	var args []ArgMeta

	for i := 0; i < len(classMeta.properties); i++ {
		arg := &ArgMeta{}
		arg.name = classMeta.properties[i].name
		arg.argType = classMeta.properties[i].propertyType
		arg.namespace = classMeta.properties[i].namespace
		args = append(args, *arg)
	}

	template := `%spublic function __construct(%s)
%s{%s
%s}`
	method = fmt.Sprintf(template, sp4, argsToString(args), sp4, classMeta.bindProps(), sp4)

	return method
}

func (meta ClassMeta) bindProps() string {
	var bindedProps []string

	sp8 := libs.Indent(8)

	for i := 0; i < len(meta.properties); i++ {
		name := meta.properties[i].name
		s := fmt.Sprintf("%s$this->%s = $%s;", sp8, name, name)
		bindedProps = append(bindedProps, s)
	}

	if len(bindedProps) == 0 {
		return ""
	} else {
		return "\n" + strings.Join(bindedProps, "\n")
	}
}

func (meta ClassMeta) createGetter() string {
	var getter []string
	sp4 := libs.Indent(4)
	sp8 := libs.Indent(8)

	for i := 0; i < len(meta.properties); i++ {
		propType := meta.properties[i].propertyType
		name := meta.properties[i].name

		value := fmt.Sprintf("return $this->%s;", name)

		if propType == "" {
			method := `%spublic function get%s()
%s{
%s%s
%s}`
			getter = append(getter, fmt.Sprintf(method, sp4, strings.Title(name), sp4, sp8, value, sp4))
		} else {

			method := `%spublic function get%s(): %s
%s{
%s%s
%s}`
			getter = append(getter, fmt.Sprintf(method, sp4, strings.Title(name), propType, sp4, sp8, value, sp4))
		}
	}

	return strings.Join(getter, "\n\n")
}

func (meta ClassMeta) ToString() string {
	template := `<?php

namespace %s;
%s
class %s%s%s
{%s%s
}
`

	return fmt.Sprintf(template, meta.namespace, meta.exportUseBlock(), meta.name, meta.exportExtendsBlack(), meta.exportImplementsBlock(), meta.exportPropertyBlock(), meta.exportMethodBlock())
}
