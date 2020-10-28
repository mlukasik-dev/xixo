package marshaller

import (
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
)

func maskAppend(mask *field_mask.FieldMask, m proto.Message, item interface{}, field string) error {
	if item != nil {
		if err := mask.Append(m, field); err != nil {
			return err
		}
	}
	return nil
}
