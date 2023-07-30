package repository

import (
	"fmt"
	"github.com/Andreas-Espelund/livedata-backend/generated/models"
	"github.com/go-openapi/strfmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

func getMockIndividual() models.Individual {
	testIndividual := models.Individual{
		BirthDate: strfmt.Date{},
		Father:    10020,
		Gender:    "FEMALE",
		ID:        30030,
		Mother:    20010,
		Status:    "ACTIVE",
	}

	return testIndividual
}

func individualToMap(individual models.Individual) (map[string]interface{}, error) {
	testIndividual := map[string]interface{}{
		"birth_date": "2020-04-10",
		"father":     10020,
		"gender":     "FEMALE",
		"id":         30030,
		"mother":     20010,
		"status":     "ACTIVE",
	}

	return testIndividual, nil
}

func mapToIndividual(in map[string]interface{}) (models.Individual, error) {
	var result models.Individual

	if err := mapstructure.Decode(in, &result); err != nil {
		return result, fmt.Errorf("BAD REQUEST")
	}
	return result, nil
}

func GetIndividual(id int64) (models.Individual, error) {

	return getMockIndividual(), nil
}

func GetAllIndividuals() ([]*models.Individual, error) {
	a := getMockIndividual()
	b := getMockIndividual()

	res := []*models.Individual{&a, &b}

	return res, nil
}

func UpdateIndividual(id int64, patches models.Patch) (models.Individual, error) {
	individual := getMockIndividual()

	mapIndividual, err := individualToMap(individual)

	if err != nil {
		fmt.Print(err)
	}

	for _, patch := range patches {
		if patch.Op == "replace" {
			attr := patch.Path

			if attr == "father" || attr == "mother" {
				i, err := strconv.ParseInt(patch.Value, 10, 64)
				if err != nil {
					return individual, fmt.Errorf("BAD REQUEST")
				}
				mapIndividual[attr] = i
			} else {
				mapIndividual[attr] = patch.Value
			}
		} else {
			return individual, fmt.Errorf("BAD REQUEST")
		}
	}

	result, err := mapToIndividual(mapIndividual)

	if err != nil {
		fmt.Print(err)
	}

	return result, nil

}

func DeleteIndividual(id int64) (models.Individual, error) {
	return getMockIndividual(), nil
}

func CreateIndividual() (models.Individual, error) {
	return getMockIndividual(), nil
}
