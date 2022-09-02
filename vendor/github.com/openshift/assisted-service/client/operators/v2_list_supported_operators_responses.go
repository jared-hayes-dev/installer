// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListSupportedOperatorsReader is a Reader for the V2ListSupportedOperators structure.
type V2ListSupportedOperatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListSupportedOperatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListSupportedOperatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListSupportedOperatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListSupportedOperatorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListSupportedOperatorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListSupportedOperatorsOK creates a V2ListSupportedOperatorsOK with default headers values
func NewV2ListSupportedOperatorsOK() *V2ListSupportedOperatorsOK {
	return &V2ListSupportedOperatorsOK{}
}

/* V2ListSupportedOperatorsOK describes a response with status code 200, with default header values.

Success.
*/
type V2ListSupportedOperatorsOK struct {
	Payload []string
}

func (o *V2ListSupportedOperatorsOK) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators][%d] v2ListSupportedOperatorsOK  %+v", 200, o.Payload)
}
func (o *V2ListSupportedOperatorsOK) GetPayload() []string {
	return o.Payload
}

func (o *V2ListSupportedOperatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListSupportedOperatorsUnauthorized creates a V2ListSupportedOperatorsUnauthorized with default headers values
func NewV2ListSupportedOperatorsUnauthorized() *V2ListSupportedOperatorsUnauthorized {
	return &V2ListSupportedOperatorsUnauthorized{}
}

/* V2ListSupportedOperatorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2ListSupportedOperatorsUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2ListSupportedOperatorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators][%d] v2ListSupportedOperatorsUnauthorized  %+v", 401, o.Payload)
}
func (o *V2ListSupportedOperatorsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListSupportedOperatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListSupportedOperatorsForbidden creates a V2ListSupportedOperatorsForbidden with default headers values
func NewV2ListSupportedOperatorsForbidden() *V2ListSupportedOperatorsForbidden {
	return &V2ListSupportedOperatorsForbidden{}
}

/* V2ListSupportedOperatorsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2ListSupportedOperatorsForbidden struct {
	Payload *models.InfraError
}

func (o *V2ListSupportedOperatorsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators][%d] v2ListSupportedOperatorsForbidden  %+v", 403, o.Payload)
}
func (o *V2ListSupportedOperatorsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListSupportedOperatorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListSupportedOperatorsInternalServerError creates a V2ListSupportedOperatorsInternalServerError with default headers values
func NewV2ListSupportedOperatorsInternalServerError() *V2ListSupportedOperatorsInternalServerError {
	return &V2ListSupportedOperatorsInternalServerError{}
}

/* V2ListSupportedOperatorsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ListSupportedOperatorsInternalServerError struct {
	Payload *models.Error
}

func (o *V2ListSupportedOperatorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators][%d] v2ListSupportedOperatorsInternalServerError  %+v", 500, o.Payload)
}
func (o *V2ListSupportedOperatorsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListSupportedOperatorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
