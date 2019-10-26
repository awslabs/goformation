package cloudformation

import (
	"encoding/json"
)

type byJSONLength []interface{}

func (s byJSONLength) Len() int {
	return len(s)
}

func (s byJSONLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byJSONLength) Less(i, j int) bool {
	// Nil is always at the end
	if s[i] == nil {
		return false
	}
	if s[j] == nil {
		return true
	}
	jsoni, _ := json.Marshal(s[i])
	jsonj, _ := json.Marshal(s[j])

	if jsoni == nil {
		return false
	}
	if jsonj == nil {
		return true
	}

	return len(jsoni) > len(jsonj)
}
