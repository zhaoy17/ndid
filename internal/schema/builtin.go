package schema

var ndiDocumentSchema = &NDISchema{
	SchemaName:  "ndi-document",
	Description: "The base properties all the DID Documents built upon",
	SchemaFields: []*NDIField{
		{
			FieldName:   "id",
			Description: "Unique identification of a document",
			DataType:    &NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "session_id",
			Description: "",
			DataType:    &NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "name",
			Description: "",
			DataType:    &NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "type",
			Description: "",
			DataType:    &NDIString{},
			Querable:    true,
		},
		{
			FieldName:   "datestamp",
			Description: "",
			DataType:    &NDIString{},
		},
		{
			FieldName:   "full_content",
			Description: "The content of the document including all the dependencies",
			DataType:    &NDIString{},
			Querable:    true,
		},
	},
}
