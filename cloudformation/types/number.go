package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MaybeInt int

func (i *MaybeInt) UnmarshalJSON(b []byte) error {
	// Treat this as a float if we possibly can
	var raw interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	ret, err := strconv.Atoi(fmt.Sprintf("%v", raw))
	if err != nil {
		return err
	}
	*i = MaybeInt(ret)

	return nil
}

type MaybeFloat float64

func (f *MaybeFloat) UnmarshalJSON(b []byte) error {
	// Treat this as a float if we possibly can
	var raw interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	ret, err := strconv.ParseFloat(fmt.Sprintf("%v", raw), 64)
	if err != nil {
		return err
	}
	*f = MaybeFloat(ret)

	return nil
}
