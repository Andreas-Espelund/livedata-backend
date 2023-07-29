// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetV0HelloOKCode is the HTTP code returned for type GetV0HelloOK
const GetV0HelloOKCode int = 200

/*
GetV0HelloOK Successful response

swagger:response getV0HelloOK
*/
type GetV0HelloOK struct {
}

// NewGetV0HelloOK creates GetV0HelloOK with default headers values
func NewGetV0HelloOK() *GetV0HelloOK {

	return &GetV0HelloOK{}
}

// WriteResponse to the client
func (o *GetV0HelloOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}