package service

import (
	"fmt"
	"strings"

	"github.com/subhammurarka/GraphQL/models"
	"github.com/subhammurarka/GraphQL/repository"
)

type MaterialService interface {
	GetMaterialAndSupplier(materialType string, price float64, location string) (models.MaterialAPI, models.Supplier, error)
}

type materialService struct {
	repo repository.MaterialRepository
}

func NewMaterialService(repo repository.MaterialRepository) MaterialService {
	return &materialService{
		repo: repo,
	}
}

func (s *materialService) GetMaterialAndSupplier(materialType string, price float64, location string) (models.MaterialAPI, models.Supplier, error) {
	materialsChan := make(chan struct {
		materials []models.MaterialAPI
		err       error
	}, 1)

	suppliersChan := make(chan struct {
		suppliers []models.SupplierAPI
		err       error
	}, 1)

	go func() {
		materials, err := s.repo.GetMaterialsByType(materialType, price)
		materialsChan <- struct {
			materials []models.MaterialAPI
			err       error
		}{materials, err}
	}()

	go func() {
		suppliers, err := s.repo.GetSuppliers()
		suppliersChan <- struct {
			suppliers []models.SupplierAPI
			err       error
		}{suppliers, err}
	}()

	materialsResult := <-materialsChan
	suppliersResult := <-suppliersChan

	if materialsResult.err != nil {
		return models.MaterialAPI{}, models.Supplier{}, fmt.Errorf("failed to fetch materials: %v", materialsResult.err)
	}

	if suppliersResult.err != nil {
		return models.MaterialAPI{}, models.Supplier{}, fmt.Errorf("failed to fetch suppliers: %v", suppliersResult.err)
	}

	filteredMaterials := materialsResult.materials
	suppliers := suppliersResult.suppliers

	for _, material := range filteredMaterials {
		for _, supplier := range suppliers {
			if !strings.Contains(strings.ToLower(supplier.SupplierLocation), strings.ToLower(location)) {
				continue
			}

			materialItem, exists := supplier.Materials[material.MaterialType]
			if !exists {
				continue
			}

			for _, materialDet := range materialItem {
				if materialDet.MaterialName != material.MaterialName || materialDet.StockAvailability != "Available" {
					continue
				}

				supplierResp := models.Supplier{
					SupplierName:      supplier.SupplierName,
					SupplierLocation:  supplier.SupplierLocation,
					StockAvailability: "Available",
				}

				return material, supplierResp, nil
			}
		}
	}

	return models.MaterialAPI{}, models.Supplier{}, fmt.Errorf("no matching material found with availability in location: %s", location)
}
