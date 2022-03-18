// Some helper functions for json.
package json_util

import "encoding/json"

type SortOptions struct {
	// If true, signature related properties will be removed.
	//
	// This is useful for signatures that are not part of the signature verification process.
	//
	// Currently, unsignable properties are:
	//    - 'signature'
	//    - 'agents'
	//
	// Default is false.
	NoSignProperties bool
}

// Sorts a given json by keys.
func SortJsonByKeys(jsonBytes []byte, opt *SortOptions) ([]byte, error) {
	var i interface{} // maps have their keys sorted lexicographically in golang

	err := json.Unmarshal(jsonBytes, &i)

	if err != nil {
		return []byte{}, err
	}

	if opt != nil {
		if opt.NoSignProperties {
			removeSignProperties(i)
		}
	}

	output, err := json.Marshal(i)

	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

// Removes signature related properties from a given json.
func removeSignProperties(i interface{}) {
	switch v := (i).(type) {
	case map[string]interface{}:
		for k := range v {
			if k == "signature" /*|| k == "agents"*/ {
				delete(v, k)
				//} else {
				//	removeSignProperties(v[k])
			}
		}
		//case []interface{}:
		//	for _, v := range v {
		//		removeSignProperties(v)
		//	}
	}
}
