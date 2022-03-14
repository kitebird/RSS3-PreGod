// Some helper functions for json.
package json_util

import "encoding/json"

// Sorts a given json by keys.
func SortJsonByKeys(jsonBytes []byte) ([]byte, error) {
	var i interface{} // maps have their keys sorted lexicographically in golang

	err := json.Unmarshal(jsonBytes, &i)

	if err != nil {
		return []byte{}, err
	}

	output, err := json.Marshal(i)

	if err != nil {
		return []byte{}, err
	}

	return output, nil
}
