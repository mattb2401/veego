package veego

import (
	"net/http"

	"github.com/mattb2401/parsrus"
)

type Response struct {
	ResponseWriter http.ResponseWriter
	ContentType    string
}

func NewResponse(writer http.ResponseWriter, contentType string) *Response {
	return &Response{
		ResponseWriter: writer,
		ContentType:    contentType,
	}
}

func (r *Response) JSON(params interface{}) {
	parser := parsrus.Parser{ResponseWriter: r.ResponseWriter, ContentType: r.ContentType}
	parser.Serialize(params)
}

func (r *Response) XML(params interface{}, rootTag string) {
	parser := parsrus.Parser{ResponseWriter: r.ResponseWriter, ContentType: r.ContentType, RootTag: rootTag}
	parser.Serialize(params)
}
