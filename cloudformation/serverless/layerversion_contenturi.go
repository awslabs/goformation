package serverless

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"

	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

// LayerVersion_ContentUri is a helper struct that can hold either a String or S3Location value
type LayerVersion_ContentUri struct {
	String *string

	S3Location *LayerVersion_S3Location
}

func (r LayerVersion_ContentUri) value() interface{} {
	ret := []interface{}{}

	if r.String != nil {
		ret = append(ret, r.String)
	}

	if r.S3Location != nil {
		ret = append(ret, *r.S3Location)
	}

	sort.Sort(utils.ByJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r LayerVersion_ContentUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *LayerVersion_ContentUri) UnmarshalJSON(b []byte) error {

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

		if err := decoder.Decode(&r.S3Location); err != nil {
			r.S3Location = nil
		}
		reader.Seek(0, io.SeekStart)

	case []interface{}:

	}

	return nil
}
