package settings

import (
	"fmt"
	"os"
)

type setting interface {
	value() string
}

var envNames = map[string]string{
	"BasicAuthUser":     "BASIC_AUTH_USER",
	"BasicAuthPassword": "BASIC_AUTH_PASSWORD",
	"User":              "MH_USERNAME",
	"Password":          "MH_PASSWORD",
	"TargetURL":         "TARGET_URL",
}

type hardcodedSetting string

func (h hardcodedSetting) value() string {
	return string(h)
}

type envSetting struct {
	envKey       string
	mustBeGiven  bool
	defaultValue string
}

func (s envSetting) value() string {
	val, exists := os.LookupEnv(s.envKey)
	if !exists {
		return s.fallbackToDefaultSetting()
	}
	return val
}

func (s envSetting) fallbackToDefaultSetting() string {
	if s.mustBeGiven {
		panic(fmt.Sprintf("environment variable %q must be given", s.envKey))
	}
	return s.defaultValue
}
