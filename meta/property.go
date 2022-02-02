package meta

type PropertyMeta struct {
	namespace    string
	name         string
	propertyType string
}

func (meta *PropertyMeta) SetNamespace(namespace string) {
	meta.namespace = namespace
}

func (meta *PropertyMeta) SetName(name string) {
	meta.name = name
}

func (meta *PropertyMeta) SetPropertyType(propertyType string) {
	meta.propertyType = propertyType
}
