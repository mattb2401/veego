package veego

import (
	"net/http"

	"github.com/mattb2401/parsrus"
)

type response struct {
	ResponseWriter http.ResponseWriter
	ContentType    string
}

func NewResponse(writer http.ResponseWriter, contentType string) *response {
	return &response{
		ResponseWriter: writer,
		ContentType:    contentType,
	}
}

func (r *response) JSON(params interface{}, httpCode ...int) {
	parser := parsrus.Parser{ResponseWriter: r.ResponseWriter, ContentType: r.ContentType}
	parser.Serialize(params, httpCode...)
}

func (r *response) XML(params interface{}, rootTag string, httpCode ...int) {
	parser := parsrus.Parser{ResponseWriter: r.ResponseWriter, ContentType: r.ContentType, RootTag: rootTag}
	parser.Serialize(params, httpCode...)
}
