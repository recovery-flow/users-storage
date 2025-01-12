package repositories

import "go.mongodb.org/mongo-driver/bson"

// Вспомогательные функции
func copyFilters(filters bson.M) bson.M {
	newFilters := bson.M{}
	for k, v := range filters {
		newFilters[k] = v
	}
	return newFilters
}

func copySort(sort bson.D) bson.D {
	newSort := make(bson.D, len(sort))
	copy(newSort, sort)
	return newSort
}
