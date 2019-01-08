package mathehappen

import "strings"

type mapping struct {
	fromKey   string        // the key the data is originally retrieved by
	toKey     string        // the key the data is mapped to
	transform transformFunc // a transformation of the data
}

type transformFunc func(data string) (transfromedData interface{})

func direct(data string) interface{} {
	return data
}

func directWithDefault(defaultData string) transformFunc {
	return func(data string) interface{} {
		if data == "" {
			return defaultData
		}
		return data
	}
}

func toStringSlice(data string) interface{} {
	if data == "" {
		return []string{}
	}
	return []string{data}
}

func constant(c interface{}) transformFunc {
	return func(_ string) interface{} {
		return c
	}
}

// filterSpaceBeforeColon is a transformFunc which will filter out the
// space immediately preceding a colon, if there is such a space. Everyting
// else will be mapped directly.
//	k, v := filterSpaceBeforeColon("Bruchzahlen : Bruchdarstellungen")
//	// k == "title"
//	// v == "Bruchzahlen: Bruchdarstellungen"
func filterSpaceBeforeColon(s string) interface{} {
	oneTime := 1
	s = strings.Replace(s, " :", ":", oneTime)
	return s
}
