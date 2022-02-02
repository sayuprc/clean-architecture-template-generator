package meta

type MethodMeta struct {
	namespace  string
	name       string
	returnType string
	args       []ArgMeta
}

func (meta *MethodMeta) SetNamespace(namespace string) {
	meta.namespace = namespace
}

func (meta *MethodMeta) SetName(name string) {
	meta.name = name
}

func (meta *MethodMeta) SetReturnType(returnType string) {
	meta.returnType = returnType
}

func (meta *MethodMeta) SetArgs(args []ArgMeta) {
	meta.args = args
}
