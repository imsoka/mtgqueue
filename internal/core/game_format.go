package core

type GameFormat int

const (
  Standard = iota
  Pioneer
  Modern
  Legacy
  Vintage
  Commander
  Oathbreaker
  Alchemy
  Historic
  Brawl
  Timeless
  Pauper
  Penny
)

var gameFormatNames = map[GameFormat]string {
  Standard:     "standard",
  Pioneer:      "pioneer",
  Modern:       "modern",
  Legacy:       "legacy",
  Vintage:      "vintage",
  Commander:    "commander",
  Oathbreaker:  "oathbreaker",
  Alchemy:      "alchemy",
  Historic:     "historic",
  Brawl:        "brawl",
  Timeless:     "timeless",
}

func (gf GameFormat) String() string {
  if name, exists := gameFormatNames[gf]; exists {
    return name
  }
  return "unknown"
}

