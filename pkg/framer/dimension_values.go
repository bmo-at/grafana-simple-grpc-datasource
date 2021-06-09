package framer

import (
	"bitbucket.org/innius/grafana-simple-grpc-datasource/pkg/framer/fields"
	pb "bitbucket.org/innius/grafana-simple-grpc-datasource/pkg/proto"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type DimensionValues struct {
	pb.ListDimensionValuesResponse
}

func (p DimensionValues) Frames() (data.Frames, error) {
	length := len(p.Results)
	labelField := fields.NewFieldWithName("label", data.FieldTypeString, length)
	nameField := fields.NewFieldWithName("value", data.FieldTypeString, length)
	descriptionField := fields.NewFieldWithName("description", data.FieldTypeString, length)

	frame := data.NewFrame("dimensions", labelField, nameField, descriptionField)

	for i, d := range p.Results {
		labelField.Set(i, d.Value)
		nameField.Set(i, d.Value)
		descriptionField.Set(i, d.Description)
	}

	return data.Frames{frame}, nil
}
