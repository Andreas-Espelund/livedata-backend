package v1

import (
	"github.com/Andreas-Espelund/livedata-backend/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func HelloHandler() operations.GetV0HelloHandlerFunc {
	return func(params operations.GetV0HelloParams) middleware.Responder {
		return operations.NewGetV0HelloOK()
	}
}
