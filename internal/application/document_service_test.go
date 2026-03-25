package application

import (
	"hornet-docs/internal/infrastructure/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(doc model.Document) {
	m.Called(doc)
}

func (m *MockRepo) FindByID(id string) (model.Document, bool) {
	args := m.Called(id)
	return args.Get(0).(model.Document), args.Bool(1)
}

func (m *MockRepo) Delete(id string) {
	m.Called(id)
}

func TestDocumentService_CreateDocument(t *testing.T) {
	t.Run("should fail if ID is missing", func(t *testing.T) {
		repo := new(MockRepo)
		svc := NewDocumentService(repo)

		err := svc.CreateDocument(model.Document{Name: "Test"})

		assert.Error(t, err)
		assert.Equal(t, "document ID is required", err.Error())
	})

	t.Run("should fail if Name is missing", func(t *testing.T) {
		repo := new(MockRepo)
		svc := NewDocumentService(repo)

		err := svc.CreateDocument(model.Document{ID: "123"})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")
	})
}
