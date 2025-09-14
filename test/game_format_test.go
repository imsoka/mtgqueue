package test

import (
	"soka/mtgqueue/internal/domain/game"
	"testing"
)


func TestGameFormat_Constructor(t *testing.T) {
  tests := []struct {
    name            string
    format          string
    expected        game.Format
    expectedError   bool
  }{
    {"Create Pauper format", "PAUPER", game.Pauper, false},
    {"Create EDH format", "EDH", game.Commander, false},
    {"CreateCCEDH format", "CEDH", game.CompetitiveCommander, false},
    {"Create PauperEDH format", "PAUPER_EDH", game.PauperCommander, false},
    {"Create Standard format", "STANDARD", game.Standard, false},
    {"Create Modern format", "MODERN", game.Modern, false},
    {"Create Vintage format", "VINTAGE", game.Vintage, false},
    {"Create Legacy format", "LEGACY", game.Legacy, false},
    {"Invalid format", "INVALID", game.Format(0), true},
    {"Invalid number", "123", game.Format(0), true},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result, err := game.NewFormatFromString(tt.format)

      if tt.expectedError {
        if err == nil {
          t.Errorf("Expeceted error for input %s, but got none", tt.format)
        }
      } else {
        if err != nil {
          t.Errorf("Unexpected error for input %s: %v", tt.format, err)
        }

        if result != tt.expected {
          t.Errorf("Expected %s, got %s", tt.expected.String(), result.String())
        }
      }
    })
  }
}

func TestGameFormat_String(t *testing.T) {
  tests := []struct {
    name          string
    format        game.Format
    expected      string
  }{
    {"Pauper to string", game.Pauper, "PAUPER"},
    {"EDH to string", game.Commander, "EDH"},
    {"CEDH to string", game.CompetitiveCommander, "CEDH"},
    {"PauperEDH to string", game.PauperCommander, "PAUPER_EDH"},
    {"Standard to string", game.Standard, "STANDARD"},
    {"Modern to string", game.Modern, "MODERN"},
    {"Vintage to string", game.Vintage, "VINTAGE"},
    {"Legacy to string", game.Legacy, "LEGACY"},
    {"InvalidFormat", game.Format(69), "UNKNOWN"},
  }

  for _,tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := tt.format.String()
      if result != tt.expected {
        t.Errorf("Expected %s but got %s", tt.expected, result)
      }
    })
  }
}

func TestGameFormat_Equals(t *testing.T) {
  tests := []struct {
    name          string
    format1       game.Format
    format2       game.Format
    expected      bool
  }{
    {"Pauper equals Pauper", game.Pauper, game.Pauper, true},
    {"Commander equals Commander", game.Commander, game.Commander, true},
    {"CCommander equals CCommander", game.CompetitiveCommander, game.CompetitiveCommander, true},
    {"PauperCommander equals PauperCommander", game.PauperCommander, game.PauperCommander, true},
    {"Modern equals Modern", game.Modern, game.Modern, true},
    {"Standard equals Standard", game.Standard, game.Standard, true},
    {"Legacy equals Legacy", game.Legacy, game.Legacy, true},
    {"Vintage equals Vintage", game.Vintage, game.Vintage, true},
    {"Pauper not equals Commander", game.Pauper, game.Commander, false},
    {"Pauper not equals Invalid format", game.Pauper, game.Format(69), false},
    {"CCommander not equals Commander", game.CompetitiveCommander, game.Commander, false},
  }

  for _,tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := tt.format1.Equals(tt.format2)

      if result != tt.expected {
        t.Errorf("Expected %v, got %v", tt.expected, result)
      }
    })
  }
}

func TestGameFormat_IsValid(t *testing.T) {
  tests := []struct {
    name          string
    format        game.Format
    expected      bool
  }{
    {"Pauper should be valid", game.Pauper, true},
    {"Commander should be valid", game.Commander, true},
    {"CEDH should be valid", game.CompetitiveCommander, true},
    {"PauperEDH should be valid", game.PauperCommander, true},
    {"Standard should be valid", game.Standard, true},
    {"Modern should be valid", game.Modern, true},
    {"Vintage should be valid", game.Vintage, true},
    {"Legacy should be valid", game.Legacy, true},
    {"Unkown format shouldn't be valid", game.Format(69), false},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := tt.format.IsValid()

      if result != tt.expected {
        t.Errorf("Expected %v, got %v", tt.expected, result)
      }
    })
  }
}
