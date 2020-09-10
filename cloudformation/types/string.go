package types

import (
	"encoding/json"
	"fmt"
)

/*
	CloudFormation is a lot more accepting of mistmatched types than the
	default Go JSON decoder.

	Here we try and match this behaviour.
*/

type MaybeString string

func (s *MaybeString) UnmarshalJSON(b []byte) error {
	// Treat this as a string, even if it's really an int or whatever
	var raw interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	*s = MaybeString(fmt.Sprintf("%v", raw))

	return nil
}
