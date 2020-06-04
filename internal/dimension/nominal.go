package dimension

import (
	. "github.com/rabitzdb/rabitz/internal/data"
)

// Nominal Dimension
func GetNominalDimension(definition DimensionDefinition, data VectorReader) Dimension {
	// Get data from Database
	vectors := data.GetVectors(definition.Dataset,definition.Offset,definition.Field)
	// Get categories
	categories := vectorsToCategories(vectors)
	return Dimension{Data:categories,Field:definition.Field}
}
// In case we have vectors for the same category
func vectorsToCategories(vectors []Vector) []Category{
	categoriesMap := make(map[string]*Category)
	for _,vector := range vectors {
		category := categoriesMap[vector.Value]
		if category != nil {
			category.Docs.Or(vector.Bits)
		} else {
			newCategory := vectorToCategory(vector)
			categoriesMap[vector.Value] = &newCategory
		}
	}
	categories := make([]Category,len(categoriesMap))
	idx := 0
	for _,value := range categoriesMap {
		categories[idx] = *value
		idx++
	}
	return categories
}
func vectorToCategory(vector Vector) Category {
	return Category{Name: vector.Value, Docs: vector.Bits}
}