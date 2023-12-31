// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Andreas-Espelund/livedata-backend/generated/models"
)

// IndividualsCreateNewV1CreatedCode is the HTTP code returned for type IndividualsCreateNewV1Created
const IndividualsCreateNewV1CreatedCode int = 201

/*
IndividualsCreateNewV1Created Individual successfully created

swagger:response individualsCreateNewV1Created
*/
type IndividualsCreateNewV1Created struct {

	/*
	  In: Body
	*/
	Payload *models.Individual `json:"body,omitempty"`
}

// NewIndividualsCreateNewV1Created creates IndividualsCreateNewV1Created with default headers values
func NewIndividualsCreateNewV1Created() *IndividualsCreateNewV1Created {

	return &IndividualsCreateNewV1Created{}
}

// WithPayload adds the payload to the individuals create new v1 created response
func (o *IndividualsCreateNewV1Created) WithPayload(payload *models.Individual) *IndividualsCreateNewV1Created {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the individuals create new v1 created response
func (o *IndividualsCreateNewV1Created) SetPayload(payload *models.Individual) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndividualsCreateNewV1Created) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
