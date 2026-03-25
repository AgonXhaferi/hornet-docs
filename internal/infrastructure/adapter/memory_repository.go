package adapter

import (
	"hornet-docs/internal/application/ports"
	"hornet-docs/internal/infrastructure/model"
	"sync"
)

func NewInMemoryRepository() ports.DocumentRepository {
	return &inMemoryRepo{
		docs: make(map[string]model.Document),
	}
}

type inMemoryRepo struct {
	mu   sync.RWMutex
	docs map[string]model.Document
}

func (r *inMemoryRepo) Save(doc model.Document) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.docs[doc.ID] = doc
}

func (r *inMemoryRepo) FindByID(id string) (model.Document, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	doc, ok := r.docs[id]
	return doc, ok
}

func (r *inMemoryRepo) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.docs, id)
}
