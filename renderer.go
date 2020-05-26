package veego

import (
	"net/http"

	"github.com/mattb2401/parsrus"
)

type renderer struct {
	ResponseWriter http.ResponseWriter
	ContentType    string
}

func NewRenderer(writer http.ResponseWriter, contentType string) *renderer {
	return &renderer{
		ResponseWriter: writer,
		ContentType:    contentType,
	}
}

func (r *renderer) JSON(params interface{}, httpCode ...int) {
	parser := parsrus.Parser{ResponseWriter: r.ResponseWriter, ContentType: r.ContentType}
	parser.Serialize(params, httpCode...)
}

func (r *renderer) XML(params interface{}, rootTag string, httpCode ...int) {
	parser := parsrus.Parser{ResponseWriter: r.ResponseWriter, ContentType: r.ContentType, RootTag: rootTag}
	parser.Serialize(params, httpCode...)
}
