package models_test

import (
	"sqli/models"
	"testing"
)

func TestVulnSQL(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {

		_, err := models.VulnGetProductsByCategory("Footwear")

		if err != nil {
			t.Fatalf("got an error, was not expecting one, %v", err)
		}
	})
	t.Run("footwear query injection", func(t *testing.T) {
		products, _ := models.VulnGetProductsByCategory(`Footwear' union select 1,@@version,3,4,5;-- -`)
		if !(len(products) > 0) {
			t.Fatal("returned table is empty")
		}
		lastProduct := products[len(products)-1]
		if lastProduct.Category == "Footwear" {
			t.Errorf("no unioned data at the end of the table: %#v", lastProduct)
		}

	})
}
