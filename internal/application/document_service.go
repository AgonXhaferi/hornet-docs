package application

import (
	"errors"
	"hornet-docs/internal/application/ports"
	"hornet-docs/internal/infrastructure/model"
)

type DocumentService struct {
	repo ports.DocumentRepository
}

func NewDocumentService(r ports.DocumentRepository) *DocumentService {
	return &DocumentService{repo: r}
}

func (s *DocumentService) CreateDocument(doc model.Document) error {
	if doc.ID == "" {
		return errors.New("document ID is required")
	}

	if doc.Name == "" {
		return errors.New("document name is required")
	}

	if _, exists := s.repo.FindByID(doc.ID); exists {
		return errors.New("document with this ID already exists")
	}

	s.repo.Save(doc)
	return nil
}

func (s *DocumentService) GetDocument(id string) (model.Document, error) {
	doc, ok := s.repo.FindByID(id)
	if !ok {
		return model.Document{}, errors.New("document not found")
	}
	return doc, nil
}

func (s *DocumentService) DeleteDocument(id string) error {
	if _, ok := s.repo.FindByID(id); !ok {
		return errors.New("cannot delete: document not found")
	}

	s.repo.Delete(id)
	return nil
}
