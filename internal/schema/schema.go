package schema

type NDISchema struct {
	SchemaName   string
	Description  string
	SchemaFields []*NDIField
	Dependencies []*NDIDependency
	Superclasses []*NDISchema
}

type NDIField struct {
	FieldName   string
	Description string
	DataType    NDIDataType
	Querable    bool
}

type NDIDependency struct {
	DependencyName  string
	SchemaDependsOn *NDISchema
}
