package validate

import (
	"errors"
	"strings"
)

// ValidateMaterialType checks if the provided locality is valid
func ValidateMaterialType(materialType string) error {
	validMaterialTypes := []string{
		"Shingles", "Tiles", "Brick", "Cement", "Wood",
		"Metal Sheets", "Glass Panels", "Plywood", "Insulation",
	}

	for _, validType := range validMaterialTypes {
		if strings.EqualFold(validType, materialType) {
			return nil
		}
	}

	return errors.New("invalid material type: choose from Shingles, Tiles, Brick, Cement, Wood, Metal Sheets, Glass Panels, Plywood, Insulation")
}

// ValidateLocality checks if the provided locality is valid
func ValidateLocality(locality string) error {
	validLocalities := []string{
		"New York", "Los Angeles", "Chicago", "Houston", "Phoenix",
		"Philadelphia", "San Antonio", "San Diego", "Dallas",
		"San Jose", "Austin", "Seattle", "Denver", "Miami",
		"Boston", "Las Vegas", "San Francisco", "Atlanta",
		"Minneapolis", "Detroit",
	}

	for _, validLocality := range validLocalities {
		if strings.EqualFold(locality, validLocality) {
			return nil
		}
	}

	return errors.New("invalid locality: please choose a valid city")
}
