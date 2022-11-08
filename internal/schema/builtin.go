package schema

import "github.com/zhaoy17/ndid/internal/validator"

var ndiDocumentSchema = &NDISchema{
	SchemaName:  "ndi-document",
	Description: "The base properties all the DID Documents built upon",
	SchemaFields: []*NDIField{
		{
			FieldName:   "id",
			Description: "Unique identification of a document",
			DataType:    &validator.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "session_id",
			Description: "",
			DataType:    &validator.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "name",
			Description: "",
			DataType:    &validator.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "type",
			Description: "",
			DataType:    &validator.NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "datestamp",
			Description: "",
			DataType:    &validator.NDIString{},
		},
		{
			FieldName:   "full_content",
			Description: "The content of the document including all the dependencies",
			DataType:    &validator.NDIString{},
			Querable:    true,
		},
	},
}
