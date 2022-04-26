// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"

	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

// Function_RequestParameters is a helper struct that can hold either a String or RequestParameter value
type Function_RequestParameters struct {
	StringArray *[]string

	RequestParameterArray *[]Function_RequestParameter
}

func (r Function_RequestParameters) value() interface{} {
	ret := []interface{}{}

	if r.StringArray != nil {
		ret = append(ret, r.StringArray)
	}

	if r.RequestParameterArray != nil {
		ret = append(ret, r.RequestParameterArray)
	}

	sort.Sort(utils.ByJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r Function_RequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *Function_RequestParameters) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case []string:
		r.StringArray = &val

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		reader := bytes.NewReader(b)
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		reader.Seek(0, io.SeekStart)

	case []interface{}:

		json.Unmarshal(b, &r.StringArray)

		json.Unmarshal(b, &r.RequestParameterArray)

	}

	return nil
}
