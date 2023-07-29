package v1

import (
	"github.com/Andreas-Espelund/livedata-backend/generated/models"
	"github.com/Andreas-Espelund/livedata-backend/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

func IndividualsGet() operations.IndividualsGetV1HandlerFunc {
	return func(params operations.IndividualsGetV1Params) middleware.Responder {
		response := operations.NewIndividualsGetV1OK()

		ind := models.Individual{
			BirthDate: strfmt.Date{},
			Father:    10020,
			Gender:    "EWE",
			ID:        params.ID,
			Mother:    20010,
			Status:    "ACTIVE",
		}
		response.WithPayload(&ind)
		return response
	}
}

func IndividualsGetAll() operations.IndividualsGetAllV1HandlerFunc {
	return func(params operations.IndividualsGetAllV1Params) middleware.Responder {
		response := operations.NewIndividualsGetAllV1OK()

		testIndividualA := models.Individual{
			BirthDate: strfmt.Date{},
			Father:    10020,
			Gender:    "FEMALE",
			ID:        30030,
			Mother:    20010,
			Status:    "ACTIVE",
		}
		testIndividualB := models.Individual{
			BirthDate: strfmt.Date{},
			Father:    10020,
			Gender:    "MALE",
			ID:        30031,
			Mother:    20010,
			Status:    "ACTIVE",
		}

		testIndividuals := []*models.Individual{
			&testIndividualA,
			&testIndividualB,
		}

		response.WithPayload(testIndividuals)

		return response

	}
}

func IndividualsCreateNew() operations.IndividualsCreateNewV1HandlerFunc {
	return func(params operations.IndividualsCreateNewV1Params) middleware.Responder {
		response := operations.NewIndividualsCreateNewV1Created()

		testIndividual := models.Individual{
			BirthDate: strfmt.Date{},
			Father:    10020,
			Gender:    "FEMALE",
			ID:        30030,
			Mother:    20010,
			Status:    "ACTIVE",
		}

		response.WithPayload(&testIndividual)

		return response
	}
}

type IndividualPatch struct {
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`
	Status    *string      `json:"status,omitempty"`
	Gender    *string      `json:"gender,omitempty"`
	Mother    *int64       `json:"mother,omitempty"`
	Father    *int64       `json:"father,omitempty"`
}

func IndividualsPatch() operations.IndividualsPatchV1HandlerFunc {
	return func(params operations.IndividualsPatchV1Params) middleware.Responder {
		testIndividual := map[string]interface{}{
			"birth_date": "2020-04-10",
			"father":     10020,
			"gender":     "FEMALE",
			"id":         30030,
			"mother":     20010,
			"status":     "ACTIVE",
		}

		patches := params.Patch

		for _, patch := range patches {
			if patch.Op == "replace" {
				attr := patch.Path

				if attr == "father" || attr == "mother" {
					i, err := strconv.ParseInt(patch.Value, 10, 64)
					if err != nil {
						return operations.NewIndividualsPatchV1BadRequest()
					}
					testIndividual[attr] = i
				} else {
					testIndividual[attr] = patch.Value
				}
			} else {
				return operations.NewIndividualsPatchV1BadRequest()
			}
		}

		var result models.Individual

		if err := mapstructure.Decode(testIndividual, &result); err != nil {
			return operations.NewIndividualsPatchV1BadRequest()
		}

		return operations.NewIndividualsPatchV1OK().WithPayload(&result)
	}
}
