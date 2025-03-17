package models

type MaterialAPI struct {
	ID           int32   `json:"id"`
	MaterialName string  `json:"materialName"`
	MaterialType string  `json:"materialType"`
	Price        float64 `json:"price"`
	Rating       int32   `json:"rating"`
	Quality      int32   `json:"Quality"`
}

type MaterialItemAPI struct {
	MaterialName      string `json:"materialName"`
	StockAvailability string `json:"stockAvailability"`
}

type SupplierAPI struct {
	ID               int                          `json:"id"`
	SupplierName     string                       `json:"supplierName"`
	SupplierLocation string                       `json:"supplierLocation"`
	Materials        map[string][]MaterialItemAPI `json:"materials"`
}

type Supplier struct {
	ID                int
	SupplierName      string
	SupplierLocation  string
	StockAvailability string
}
