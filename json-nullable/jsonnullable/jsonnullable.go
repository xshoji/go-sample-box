package jsonnullable

import (
	"database/sql"
	"encoding/json"
)

// > One simple way to handle NULL database value in Golang ｜ by Raymond Hartoyo ｜ Medium
// > https://medium.com/@raymondhartoyo/one-simple-way-to-handle-null-database-value-in-golang-86437ec75089
// --
// > How I handled the null possible value in a sql database row in golang?
// > https://gist.github.com/rsudip90/45fad7d8959c58bcc91d464873b50013
type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	err := json.Unmarshal(data, &s.String)
	s.Valid = err == nil
	return err
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

func (f *NullFloat64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Float64, f.Valid = 0.0, false
		return nil
	}
	err := json.Unmarshal(data, &f.Float64)
	f.Valid = err == nil
	return err
}

type NullBool struct {
	sql.NullBool
}

func (b NullBool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

func (b *NullBool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Bool, b.Valid = false, false
		return nil
	}
	err := json.Unmarshal(data, &b.Bool)
	b.Valid = err == nil
	return err
}
