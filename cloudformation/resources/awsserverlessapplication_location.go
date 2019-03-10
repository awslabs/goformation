package resources

import (
	"sort"

	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// AWSServerlessApplication_Location is a helper struct that can hold either a ApplicationLocation value
type AWSServerlessApplication_Location struct {
	ApplicationLocation *AWSServerlessApplication_ApplicationLocation
}

func (r AWSServerlessApplication_Location) value() interface{} {

	ret := []interface{}{}

	if r.ApplicationLocation != nil {
		ret = append(ret, *r.ApplicationLocation)
	}

	sort.Sort(byJSONLength(ret))
	if len(ret) > 0 {
		return ret[0]
	}

	return nil

}

func (r AWSServerlessApplication_Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessApplication_Location) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case map[string]interface{}:

		mapstructure.Decode(val, &r.ApplicationLocation)

	case []interface{}:

	}

	return nil
}
