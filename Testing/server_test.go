package service_test

import (
	"testing"

	"github.com/subhammurarka/GraphQL/models"
	"github.com/subhammurarka/GraphQL/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMaterialRepository struct {
	mock.Mock
}

func (m *MockMaterialRepository) GetMaterialsByType(materialType string, price float64) ([]models.MaterialAPI, error) {
	args := m.Called(materialType, price)
	return args.Get(0).([]models.MaterialAPI), args.Error(1)
}

func (m *MockMaterialRepository) GetSuppliers() ([]models.SupplierAPI, error) {
	args := m.Called()
	return args.Get(0).([]models.SupplierAPI), args.Error(1)
}

func TestGetMaterialAndSupplier(t *testing.T) {
	mockRepo := new(MockMaterialRepository)
	service := service.NewMaterialService(mockRepo)

	materialResponse := []models.MaterialAPI{
		{
			MaterialName: "Tamko Heritage",
			MaterialType: "Cement",
			Price:        471,
			Rating:       1,
			Quality:      5,
		},
	}

	supplierResponse := []models.SupplierAPI{
		{
			SupplierName:     "Superior Roofing Supplies",
			SupplierLocation: "New York, Orlando",
			Materials: map[string][]models.MaterialItemAPI{
				"Cement": {
					{
						MaterialName:      "Tamko Heritage",
						StockAvailability: "Available",
					},
				},
			},
		},
	}

	mockRepo.On("GetMaterialsByType", "Cement", 500.0).Return(materialResponse, nil)
	mockRepo.On("GetSuppliers").Return(supplierResponse, nil)

	material, supplier, err := service.GetMaterialAndSupplier("Cement", 500.0, "New York")

	assert.Nil(t, err)
	assert.Equal(t, "Tamko Heritage", material.MaterialName)
	assert.Equal(t, "Superior Roofing Supplies", supplier.SupplierName)
	assert.Equal(t, "Available", supplier.StockAvailability)
}
