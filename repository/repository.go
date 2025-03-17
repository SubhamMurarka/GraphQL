package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	config "github.com/subhammurarka/GraphQL/Config"
	"github.com/subhammurarka/GraphQL/models"
)

// MaterialRepository defines the interface for material data operations
type MaterialRepository interface {
	GetMaterialsByType(materialType string, price float64) ([]models.MaterialAPI, error)
	GetSuppliers() ([]models.SupplierAPI, error)
}

type materialRepository struct {
	client *http.Client
}

// NewMaterialRepository creates a new material repository instance
func NewMaterialRepository() MaterialRepository {
	repo := &materialRepository{
		client: &http.Client{},
	}
	return repo
}

// GetMaterialsByType retrieves materials of a specific type within a price range
func (r *materialRepository) GetMaterialsByType(materialType string, price float64) ([]models.MaterialAPI, error) {
	resp, err := r.client.Get(config.Config.MaterialAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch materials: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("materials API returned status: %d", resp.StatusCode)
	}

	var materials []models.MaterialAPI
	if err := json.NewDecoder(resp.Body).Decode(&materials); err != nil {
		return nil, fmt.Errorf("failed to decode materials response: %v", err)
	}

	fmt.Println(materials, "\n\n\n")

	var filteredMaterials []models.MaterialAPI
	for _, material := range materials {
		if material.MaterialType == materialType && material.Price <= price {
			filteredMaterials = append(filteredMaterials, material)
		}
	}

	sort.Slice(filteredMaterials, func(i, j int) bool {
		if filteredMaterials[i].Rating != filteredMaterials[j].Rating {
			return filteredMaterials[i].Rating > filteredMaterials[j].Rating
		}
		return filteredMaterials[i].Quality > filteredMaterials[j].Quality
	})

	fmt.Println(filteredMaterials, "\n\n\n")
	return filteredMaterials, nil
}

// GetSuppliers retrieves all suppliers
func (r *materialRepository) GetSuppliers() ([]models.SupplierAPI, error) {
	resp, err := r.client.Get(config.Config.SupplierAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch suppliers: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("supplier API returned status: %d", resp.StatusCode)
	}

	var suppliers []models.SupplierAPI
	if err := json.NewDecoder(resp.Body).Decode(&suppliers); err != nil {
		return nil, fmt.Errorf("failed to decode supplier response: %v", err)
	}

	fmt.Println(suppliers, "\n\n\n")

	return suppliers, nil
}
