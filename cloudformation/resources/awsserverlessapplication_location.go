package resources

import (
	"sort"

	"encoding/json"
)

// AWSServerlessApplication_Location is a helper struct that can hold either a String or ApplicationLocation value
type AWSServerlessApplication_Location struct {
	String *string

	ApplicationLocation *AWSServerlessApplication_ApplicationLocation
}

func (r AWSServerlessApplication_Location) value() interface{} {
	ret := []interface{}{}

	if r.String != nil {
		ret = append(ret, r.String)
	}

	if r.ApplicationLocation != nil {
		ret = append(ret, *r.ApplicationLocation)
	}

	sort.Sort(byJSONLength(ret)) // Heuristic to select best attribute
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

	case string:
		r.String = &val

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		json.Unmarshal(b, &r.ApplicationLocation)

	case []interface{}:

	}

	return nil
}
