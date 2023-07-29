// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Andreas-Espelund/livedata-backend/generated/models"
)

// IndividualsPatchV1OKCode is the HTTP code returned for type IndividualsPatchV1OK
const IndividualsPatchV1OKCode int = 200

/*
IndividualsPatchV1OK Success

swagger:response individualsPatchV1OK
*/
type IndividualsPatchV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.Individual `json:"body,omitempty"`
}

// NewIndividualsPatchV1OK creates IndividualsPatchV1OK with default headers values
func NewIndividualsPatchV1OK() *IndividualsPatchV1OK {

	return &IndividualsPatchV1OK{}
}

// WithPayload adds the payload to the individuals patch v1 o k response
func (o *IndividualsPatchV1OK) WithPayload(payload *models.Individual) *IndividualsPatchV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the individuals patch v1 o k response
func (o *IndividualsPatchV1OK) SetPayload(payload *models.Individual) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndividualsPatchV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// IndividualsPatchV1BadRequestCode is the HTTP code returned for type IndividualsPatchV1BadRequest
const IndividualsPatchV1BadRequestCode int = 400

/*
IndividualsPatchV1BadRequest Bad request

swagger:response individualsPatchV1BadRequest
*/
type IndividualsPatchV1BadRequest struct {
}

// NewIndividualsPatchV1BadRequest creates IndividualsPatchV1BadRequest with default headers values
func NewIndividualsPatchV1BadRequest() *IndividualsPatchV1BadRequest {

	return &IndividualsPatchV1BadRequest{}
}

// WriteResponse to the client
func (o *IndividualsPatchV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndividualsPatchV1NotFoundCode is the HTTP code returned for type IndividualsPatchV1NotFound
const IndividualsPatchV1NotFoundCode int = 404

/*
IndividualsPatchV1NotFound Not found

swagger:response individualsPatchV1NotFound
*/
type IndividualsPatchV1NotFound struct {
}

// NewIndividualsPatchV1NotFound creates IndividualsPatchV1NotFound with default headers values
func NewIndividualsPatchV1NotFound() *IndividualsPatchV1NotFound {

	return &IndividualsPatchV1NotFound{}
}

// WriteResponse to the client
func (o *IndividualsPatchV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// IndividualsPatchV1InternalServerErrorCode is the HTTP code returned for type IndividualsPatchV1InternalServerError
const IndividualsPatchV1InternalServerErrorCode int = 500

/*
IndividualsPatchV1InternalServerError Internal server error

swagger:response individualsPatchV1InternalServerError
*/
type IndividualsPatchV1InternalServerError struct {
}

// NewIndividualsPatchV1InternalServerError creates IndividualsPatchV1InternalServerError with default headers values
func NewIndividualsPatchV1InternalServerError() *IndividualsPatchV1InternalServerError {

	return &IndividualsPatchV1InternalServerError{}
}

// WriteResponse to the client
func (o *IndividualsPatchV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
