package schema

import datatypes "github.com/zhaoy17/ndid/internal/datatypes"

var ndiDocumentSchema = &NDISchema{
	SchemaName:  "ndi-document",
	Description: "The base properties all the DID Documents built upon",
	SchemaFields: []*NDIField{
		{
			FieldName:   "id",
			Description: "Unique identification of a document",
			DataType:    &datatypes.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "session_id",
			Description: "",
			DataType:    &datatypes.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "name",
			Description: "",
			DataType:    &datatypes.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "type",
			Description: "",
			DataType:    &datatypes.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "datestamp",
			Description: "",
			DataType:    &datatypes.NDIString{},
		},
		{
			FieldName:   "full_content",
			Description: "The content of the document including all the dependencies",
			DataType:    &datatypes.NDIString{},
			Querable:    true,
		},
	},
}
