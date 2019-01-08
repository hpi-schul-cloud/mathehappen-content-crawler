package mathehappen

import (
	"encoding/json"
	"reflect"
	"testing"
)

const data = `[{"coursesId": "Klasse_5_6",
"coursesCaption": "Klasse 5 und 6",
"themeId": "Bruchzahlen",
"themeCaption": "Bruchzahlen",
"originId": "Klasse_5_6-Bruchzahlen-Umformungen",
"title": "Bruchzahlen : Bruch angeben",
"description": "",
"url": "https://example.com/aufgaben/Bruchzahlen/xyz/",
"image": "https://example.com/aufgaben/Bruchzahlen/xyz/img",
"thumb": "https://example.com/aufgaben/Bruchzahlen/xyz/thumbnail1"}]`

const expected = `{"originId": "Klasse_5_6-Bruchzahlen-Umformungen",
"title": "Bruchzahlen: Bruch angeben",
"tags": ["Bruchzahlen"],
"thumbnail": "https://example.com/aufgaben/Bruchzahlen/xyz/thumbnail1",
"url": "https://example.com/aufgaben/Bruchzahlen/xyz/",
"providerName": "Mathehappen",
"contentCategory": "learning-object",
"description": "Keine Beschreibung",
"licenses": ["© Mathehappen"],
"mimeType": "text/html",
"nonOer": true}`

func TestTransform(t *testing.T) {
	happen := Transform([]byte(data))
	if len(happen) != 1 {
		t.Errorf("expected only one result, got %d", len(happen))
		if len(happen) < 1 {
			return
		}
	}

	actual := string(happen[0])
	if !equalJSONs(expected, actual) {
		t.Errorf("\n%s\n\n", actual)
		t.Errorf("actual json does not conform to the expected outcome")
	}
}

func equalJSONs(expected, actual string) bool {
	var parsedExpected interface{}
	var parsedActual interface{}

	err := json.Unmarshal([]byte(expected), &parsedExpected)
	if err != nil {
		panic("cannot unmarshal expected: " + err.Error())
	}
	err = json.Unmarshal([]byte(actual), &parsedActual)
	if err != nil {
		panic("cannot unmarshal actual: " + err.Error())
	}

	return reflect.DeepEqual(parsedExpected, parsedActual)
}
