package resources

import (
	"sort"

	"encoding/json"
)

// AWSServerlessFunction_Policies is a helper struct that can hold either a String, String, IAMPolicyDocument, SAMPolicyTemplate, or IAMPolicyDocument value
type AWSServerlessFunction_Policies struct {
	String *string

	StringArray *[]string

	IAMPolicyDocument *AWSServerlessFunction_IAMPolicyDocument

	IAMPolicyDocumentArray *[]AWSServerlessFunction_IAMPolicyDocument
	SAMPolicyTemplateArray *[]AWSServerlessFunction_SAMPolicyTemplate
}

func (r AWSServerlessFunction_Policies) value() interface{} {
	ret := []interface{}{}

	if r.String != nil {
		ret = append(ret, r.String)
	}

	if r.StringArray != nil {
		ret = append(ret, r.StringArray)
	}

	if r.IAMPolicyDocument != nil {
		ret = append(ret, *r.IAMPolicyDocument)
	}

	if r.IAMPolicyDocumentArray != nil {
		ret = append(ret, r.IAMPolicyDocumentArray)
	}

	if r.SAMPolicyTemplateArray != nil {
		ret = append(ret, r.SAMPolicyTemplateArray)
	}

	sort.Sort(byJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r AWSServerlessFunction_Policies) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessFunction_Policies) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case string:
		r.String = &val

	case []string:
		r.StringArray = &val

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		json.Unmarshal(b, &r.IAMPolicyDocument)

	case []interface{}:

		json.Unmarshal(b, &r.StringArray)

		json.Unmarshal(b, &r.IAMPolicyDocumentArray)

		json.Unmarshal(b, &r.SAMPolicyTemplateArray)

	}

	return nil
}
