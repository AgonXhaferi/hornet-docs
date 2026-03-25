package ports

import "hornet-docs/internal/infrastructure/model"

type DocumentRepository interface {
	Save(doc model.Document)
	FindByID(id string) (model.Document, bool)
	Delete(id string)
}
