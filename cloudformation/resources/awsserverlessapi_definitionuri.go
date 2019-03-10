package resources

import (
	"sort"

	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// AWSServerlessApi_DefinitionUri is a helper struct that can hold either a String or S3Location value
type AWSServerlessApi_DefinitionUri struct {
	String *string

	S3Location *AWSServerlessApi_S3Location
}

func (r AWSServerlessApi_DefinitionUri) value() interface{} {

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

func (r AWSServerlessApi_DefinitionUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessApi_DefinitionUri) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case string:
		r.String = &val

	case map[string]interface{}:

		mapstructure.Decode(val, &r.S3Location)

	case []interface{}:

	}

	return nil
}
