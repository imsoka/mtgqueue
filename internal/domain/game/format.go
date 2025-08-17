package game

type Format int

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

var gameFormatNames = map[Format]string {
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
  Pauper:       "pauper",
  Penny:        "penny",
}

func (gf Format) String() string {
  if name, exists := gameFormatNames[gf]; exists {
    return name
  }
  return "unknown"
}

