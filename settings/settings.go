// Package settings exposes the settings the program is started with. The
// settings values are prepared in this packages init() function, so no further
// setup is needed to use them, except importing this package. The package panics
// if required settings values can not be obtained.
//
// Access the settings values as follows
//	import "github.com/schul-cloud/mathehappen-crawler/settings"
//
//	...
//
//	func someFunc() {
//		var scUserName = settings.Settings[settings.BasicAuthUser]
//		...
//	}
package settings

// String constants that are the keys in the settings map.
const (
	BasicAuthUser     = "BasicAuthUser"
	BasicAuthPassword = "BasicAuthPassword"
	ContentLocation   = "ContentLocation"
	TargetURL         = "TargetURL"
	MHappenUser       = "MHappenUser"
	MHappenPassword   = "MHappenPassoword"
)

// Settings exposes the settings given to the program as key-value pairs.
var Settings = map[string]string{}

var settings = map[string]setting{
	BasicAuthUser:     envSetting{envKey: "BASIC_AUTH_USER", defaultValue: "schulcloud-content-1"},
	BasicAuthPassword: envSetting{envKey: "BASIC_AUTH_PASSWORD", defaultValue: "content-1"},
	ContentLocation:   hardcodedSetting("https://mathehappen.de/api/materials/"),
	TargetURL:         envSetting{envKey: "TARGET_URL", defaultValue: "http://localhost:4040/resources"},
	MHappenUser:       envSetting{envKey: "MH_USER", mustBeGiven: true},
	MHappenPassword:   envSetting{envKey: "MH_PASSWORD", mustBeGiven: true},
}

func init() {
	for name, s := range settings {
		Settings[name] = s.value()
	}
}
