// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bah2830/unifi_api/pkg/models"
)

// SitesListReader is a Reader for the SitesList structure.
type SitesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SitesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSitesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSitesListOK creates a SitesListOK with default headers values
func NewSitesListOK() *SitesListOK {
	return &SitesListOK{}
}

/*SitesListOK handles this case with default header values.

Login response with cookie
*/
type SitesListOK struct {
	Payload *models.SitesListResponse
}

func (o *SitesListOK) Error() string {
	return fmt.Sprintf("[GET /stat/sites][%d] sitesListOK  %+v", 200, o.Payload)
}

func (o *SitesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SitesListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}