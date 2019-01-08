package mathehappen

import (
	"encoding/json"
)

var mappings = []mapping{
	{fromKey: "originId", toKey: "originId", transform: direct},
	{fromKey: "title", toKey: "title", transform: filterSpaceBeforeColon},
	{fromKey: "description", toKey: "description", transform: directWithDefault("Keine Beschreibung")},
	{fromKey: "url", toKey: "url", transform: direct},
	{fromKey: "themeCaption", toKey: "tags", transform: toStringSlice},
	{fromKey: "thumb", toKey: "thumbnail", transform: direct},
	{toKey: "contentCategory", transform: constant("learning-object")},
	{toKey: "mimeType", transform: constant("text/html")},
	{toKey: "providerName", transform: constant("Mathehappen")},
	{toKey: "licenses", transform: constant([]string{`© Mathehappen`})},
	{toKey: "nonOer", transform: constant(true)},
}

// Transform reads the data downloaded from mathehappen.de and transforms it
// into the data (format) expected by the schul-cloud content server.
func Transform(inBytes []byte) (outBytes [][]byte) {
	dataIn := unmarshal(inBytes)
	dataOut := transformSlice(dataIn)
	return marshal(dataOut)
}

func transformSlice(data []map[string]string) (out []map[string]interface{}) {
	for _, item := range data {
		out = append(out, transform(item))
	}
	return out
}

func transform(data map[string]string) (t map[string]interface{}) {
	t = map[string]interface{}{}
	for _, m := range mappings {
		// Get the original value sent from mathehappen.de
		// and save the transformed value and its name.
		originalValue := ""
		if m.fromKey != "" {
			originalValue, _ = data[m.fromKey]
		}
		t[m.toKey] = m.transform(originalValue)
	}
	return t
}

func unmarshal(data []byte) (v []map[string]string) {
	err := json.Unmarshal(data, &v)
	if err != nil {
		panic(err)
	}
	return v
}

func marshal(data []map[string]interface{}) (b [][]byte) {
	for _, happen := range data {
		bytes, err := json.Marshal(happen)
		if err != nil {
			panic(err)
		}
		b = append(b, bytes)
	}
	return b
}
