package domain


type Identifier interface {
  String() string
  IsEmpty() bool
  IsNil() bool
  Equals(other Identifier) bool
  MarshalJSON() ([]byte, error)
  UnmarshalJSON([]byte) error
}
