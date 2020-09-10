package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MaybeBool bool

func (o *MaybeBool) UnmarshalJSON(b []byte) error {
	// CFN accepts the strings 'true' and 'false' as boolean
	var raw interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	var l bool
	switch v := raw.(type) {
	case bool:
		l = v
	case string:
		lowerV := strings.ToLower(v)
		if lowerV == "false" {
			l = false
		}
		if lowerV == "true" {
			l = true
		}
	case nil:
		l = false
	default:
		return fmt.Errorf("invalid type %T - expected bool or bool-like-string", raw)
	}
	*o = MaybeBool(l)
	return nil
}
