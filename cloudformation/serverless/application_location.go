package serverless

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"

	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

// Application_Location is a helper struct that can hold either a String or ApplicationLocation value
type Application_Location struct {
	String *string

	ApplicationLocation *Application_ApplicationLocation
}

func (r Application_Location) value() interface{} {
	ret := []interface{}{}

	if r.String != nil {
		ret = append(ret, r.String)
	}

	if r.ApplicationLocation != nil {
		ret = append(ret, *r.ApplicationLocation)
	}

	sort.Sort(utils.ByJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r Application_Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *Application_Location) UnmarshalJSON(b []byte) error {

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

		reader := bytes.NewReader(b)
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.ApplicationLocation); err != nil {
			r.ApplicationLocation = nil
		}
		reader.Seek(0, io.SeekStart)

	case []interface{}:

	}

	return nil
}
