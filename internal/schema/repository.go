package schema

import "context"

type DIDSchemaRepository interface {
	GetSchema(schemaName []string, ctx context.Context) (*NDISchema, error)
	GetAllSchemas(ctx context.Context) (map[string]*NDISchema, error)
	InsertSchemas(schemas []*NDISchema, ctx context.Context) error
	DeleteSchemas(schemaNames []string, ctx context.Context) error
	UpdateSchema(schemaName string, fieldsToUpdateInto map[string]*NDIField, ctx context.Context) error
	Setup(context.Context) error
}
