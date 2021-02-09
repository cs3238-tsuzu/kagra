package cachedanalyzer

import (
	"bytes"
	"encoding/json"
)

func keyForCache(method string, args ...interface{}) string {
	b, err := json.Marshal(
		append(
			[]interface{}{method},
			args...,
		),
	)

	if err != nil {
		// You shouldn't call this with unmarshalable objects
		panic(err)
	}

	return string(b)
}

func copyViaJSON(in, out interface{}) {
	buf := bytes.NewBuffer(nil)

	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
