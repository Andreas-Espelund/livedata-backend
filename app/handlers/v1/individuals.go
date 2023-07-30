package v1

import (
	"fmt"
	"github.com/Andreas-Espelund/livedata-backend/app"
	"github.com/Andreas-Espelund/livedata-backend/app/repository"
	"github.com/Andreas-Espelund/livedata-backend/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func IndividualsGet(api app.Api) operations.IndividualsGetV1HandlerFunc {
	return func(params operations.IndividualsGetV1Params) middleware.Responder {

		testIndividual, err := repository.GetIndividual(api.Db, params.ID)

		if err != nil {
			fmt.Print(err)
			return nil
		}
		return operations.NewIndividualsGetV1OK().WithPayload(&testIndividual)
	}
}

func IndividualsGetAll(api app.Api) operations.IndividualsGetAllV1HandlerFunc {
	return func(params operations.IndividualsGetAllV1Params) middleware.Responder {

		testers, err := repository.GetAllIndividuals(api.Db)

		if err != nil {
			fmt.Print(err)
			return nil
		}

		return operations.NewIndividualsGetAllV1OK().WithPayload(testers)
	}
}

func IndividualsCreateNew(api app.Api) operations.IndividualsCreateNewV1HandlerFunc {
	return func(params operations.IndividualsCreateNewV1Params) middleware.Responder {

		testIndividual, err := repository.CreateIndividual(api.Db)

		if err != nil {
			fmt.Print(err)
			return nil
		}

		return operations.NewIndividualsCreateNewV1Created().WithPayload(&testIndividual)
	}
}

func IndividualsPatch(api app.Api) operations.IndividualsPatchV1HandlerFunc {
	return func(params operations.IndividualsPatchV1Params) middleware.Responder {
		testIndividual, err := repository.UpdateIndividual(api.Db, params.ID, params.Patch)

		if err != nil {
			fmt.Print(err)
			return operations.NewIndividualsPatchV1BadRequest()
		}

		return operations.NewIndividualsPatchV1OK().WithPayload(&testIndividual)
	}
}

func IndividualsDelete(api app.Api) operations.IndividualsDeleteV1HandlerFunc {
	return func(params operations.IndividualsDeleteV1Params) middleware.Responder {

		testIndividual, err := repository.DeleteIndividual(api.Db, params.ID)

		if err != nil {
			fmt.Print(err)
			return operations.NewIndividualsDeleteV1InternalServerError()
		}

		return operations.NewIndividualsDeleteV1OK().WithPayload(&testIndividual)
	}
}
