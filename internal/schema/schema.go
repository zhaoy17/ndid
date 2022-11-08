package schema

import datatypes "github.com/zhaoy17/ndid/internal/datatypes"

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
	DataType    datatypes.NDIDataType
	Querable    bool
}

type NDIDependency struct {
	DependencyName  string
	SchemaDependsOn *NDISchema
}
