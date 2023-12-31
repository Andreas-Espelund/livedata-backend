// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/Andreas-Espelund/livedata-backend/app"
	"github.com/Andreas-Espelund/livedata-backend/app/handlers/v1"
	"net/http"

	"github.com/Andreas-Espelund/livedata-backend/generated/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ../../generated --name LivedataBackend --spec ../../manifests/openApi/swagger.yaml --principal interface{}

func configureFlags(api *operations.LivedataBackendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LivedataBackendAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// v1 handlers

	ap := app.New()

	// PATH: /individuals
	api.IndividualsGetAllV1Handler = v1.IndividualsGetAll(ap)
	api.IndividualsCreateNewV1Handler = v1.IndividualsCreateNew(ap)

	// PATH: /individuals/{id}
	api.IndividualsPatchV1Handler = v1.IndividualsPatch(ap)
	api.IndividualsGetV1Handler = v1.IndividualsGet(ap)
	api.IndividualsDeleteV1Handler = v1.IndividualsDelete(ap)

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
