package helpers

import "fmt"

type AttributeData struct {
	Name     string
	Comment  string
	Category string
	Value    string
}

func MakeAttributes(data []UserAttributeRaw, metadata map[string]AttributeMetadataRaw) []AttributeData {
	attrs := []AttributeData{}

	for _, raw := range data {
		attr := AttributeData{
			Category: raw.Category,
		}
		if len(raw.FloatValue) > 0 {
			attr.Value = fmt.Sprintf("%f", raw.FloatValue[0])
		} else if len(raw.KeywordValue) > 0 {
			attr.Value = raw.KeywordValue[0]
		}
		if meta, ok := metadata[raw.Key]; ok {
			attr.Name = meta.Name
			attr.Comment = meta.Comment
		}
		attrs = append(attrs, attr)
	}

	return attrs
}
