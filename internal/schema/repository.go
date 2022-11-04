package schema

type DIDSchemaRepository interface {
	GetSchema(schemaName []string) (*NDISchema, error)
	GetAllSchemas() (map[string]*NDISchema, error)
	InsertSchemas(schemas []*NDISchema) error
	DeleteSchemas(schemaNames []string) error
	UpdateSchema(schemaName string, fieldsToUpdateInto map[string]*NDIField) error
}

type SQLSchemaRepository struct {
}

func (schemaRepository *SQLSchemaRepository) InsertSchemas(schemas []*NDISchema) error {
	return nil
}
