package resources

import (
	"sort"

	"encoding/json"
)

// AWSServerlessFunction_CodeUri is a helper struct that can hold either a String or S3Location value
type AWSServerlessFunction_CodeUri struct {
	String *string

	S3Location *AWSServerlessFunction_S3Location
}

func (r AWSServerlessFunction_CodeUri) value() interface{} {

	if r.String != nil {
		return r.String
	}

	ret := []interface{}{}

	if r.S3Location != nil {
		ret = append(ret, *r.S3Location)
	}

	sort.Sort(byJSONLength(ret))
	if len(ret) > 0 {
		return ret[0]
	}

	return nil

}

func (r AWSServerlessFunction_CodeUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessFunction_CodeUri) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case string:
		r.String = &val

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		json.Unmarshal(b, &r.S3Location)

	case []interface{}:

	}

	return nil
}
