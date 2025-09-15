package game

import (
	"errors"
	"strings"
)

type Format int

const (
	Commander Format = iota
	CompetitiveCommander
	Pauper
	PauperCommander
	Modern
	Standard
	Legacy
	Vintage
)

var formatStrings = map[Format]string{
	Commander:            "EDH",
	CompetitiveCommander: "CEDH",
	Pauper:               "PAUPER",
	PauperCommander:      "PAUPER_EDH",
	Standard:             "STANDARD",
	Modern:               "MODERN",
	Legacy:               "LEGACY",
	Vintage:              "VINTAGE",
}

var stringToFormat = map[string]Format{
	"EDH":        Commander,
	"CEDH":       CompetitiveCommander,
	"PAUPER":     Pauper,
	"PAUPER_EDH": PauperCommander,
	"STANDARD":   Standard,
	"MODERN":     Modern,
	"LEGACY":     Legacy,
	"VINTAGE":    Vintage,
}

func NewFormatFromString(s string) (Format, error) {
	if s == "" {
		return Format(0), errors.New("Empty string cannot create a game format")
	}

	upperS := strings.ToUpper(s)

	if format, exists := stringToFormat[upperS]; exists {
		return format, nil
	}

	return Format(0), errors.New("Invalid game format string:" + s)
}

func (f Format) String() string {
	if str, exists := formatStrings[f]; exists {
		return str
	}

	return "UNKNOWN"
}

func (f Format) Equals(other Format) bool {
	return f == other
}

func (f Format) IsValid() bool {
	if _, exists := formatStrings[f]; exists {
		return true
	}

	return false
}
