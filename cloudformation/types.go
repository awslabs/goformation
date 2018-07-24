package cloudformation

import (
	"encoding/json"
)

type stringIntrinsic struct {
	stringValue    *string
	intrinsicValue *intrinsic
}

type intrinsic struct {
	self map[string]interface{}
}

func NewString(value string) *stringIntrinsic {
	return &stringIntrinsic{
		stringValue: &value,
	}
}

func NewStringIntrinsic(key string, value interface{}) *stringIntrinsic {
	return &stringIntrinsic{
		intrinsicValue: &intrinsic{
			self: map[string]interface{}{
				key: value,
			},
		},
	}
}

func NewStringRef(value string) *stringIntrinsic {
	return NewStringIntrinsic("Ref", value)
}

func (r stringIntrinsic) value() interface{} {

	if r.stringValue != nil {
		return r.stringValue
	}

	if r.intrinsicValue != nil {
		return r.intrinsicValue.self
	}

	return nil

}

func (r *stringIntrinsic) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

func (r *stringIntrinsic) UnmarshalJSON(b []byte) error {
	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {
	case string:
		r.stringValue = &val
	case map[string]interface{}:
		r.intrinsicValue = &intrinsic{self: val}

	}

	return nil
}
