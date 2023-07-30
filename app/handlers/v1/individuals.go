package v1

import (
	"fmt"
	"github.com/Andreas-Espelund/livedata-backend/app/db/repository"
	"github.com/Andreas-Espelund/livedata-backend/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func IndividualsGet() operations.IndividualsGetV1HandlerFunc {
	return func(params operations.IndividualsGetV1Params) middleware.Responder {

		testIndividual, err := repository.GetIndividual(params.ID)

		if err != nil {
			fmt.Print(err)
			return nil
		}
		return operations.NewIndividualsGetV1OK().WithPayload(&testIndividual)
	}
}

func IndividualsGetAll() operations.IndividualsGetAllV1HandlerFunc {
	return func(params operations.IndividualsGetAllV1Params) middleware.Responder {

		testers, err := repository.GetAllIndividuals()

		if err != nil {
			fmt.Print(err)
			return nil
		}

		return operations.NewIndividualsGetAllV1OK().WithPayload(testers)
	}
}

func IndividualsCreateNew() operations.IndividualsCreateNewV1HandlerFunc {
	return func(params operations.IndividualsCreateNewV1Params) middleware.Responder {

		testIndividual, err := repository.CreateIndividual()

		if err != nil {
			fmt.Print(err)
			return nil
		}

		return operations.NewIndividualsCreateNewV1Created().WithPayload(&testIndividual)
	}
}

func IndividualsPatch() operations.IndividualsPatchV1HandlerFunc {
	return func(params operations.IndividualsPatchV1Params) middleware.Responder {
		testIndividual, err := repository.UpdateIndividual(params.ID, params.Patch)

		if err != nil {
			fmt.Print(err)
			return operations.NewIndividualsPatchV1BadRequest()
		}

		return operations.NewIndividualsPatchV1OK().WithPayload(&testIndividual)
	}
}

func IndividualsDelete() operations.IndividualsDeleteV1HandlerFunc {
	return func(params operations.IndividualsDeleteV1Params) middleware.Responder {

		testIndividual, err := repository.DeleteIndividual(params.ID)

		if err != nil {
			fmt.Print(err)
			return operations.NewIndividualsDeleteV1InternalServerError()
		}

		return operations.NewIndividualsDeleteV1OK().WithPayload(&testIndividual)
	}
}
