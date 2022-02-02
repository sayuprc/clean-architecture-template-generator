package meta

type ArgMeta struct {
	namespace string
	name      string
	argType   string
}

func (meta *ArgMeta) SetNamespace(namespace string) {
	meta.namespace = namespace
}

func (meta *ArgMeta) SetName(name string) {
	meta.name = name
}

func (meta *ArgMeta) SetArgType(argType string) {
	meta.argType = argType
}
