package confparams

import (
	"bytes"
	"encoding/gob"
	"errors"
)

type Serializer map[string]interface{}

type Global struct {
	I int
	S string
	F float64
	L []*Local
}
type Local struct {
	I int32
	B []byte
}

func NewSerializer() Serializer {
	M := make(map[string]interface{})
	return M
}

func NewGlobal(numLocal int) (*Global, error) {
	G := new(Global)
	if numLocal <= 0 {
		return nil, errors.New("numLocal needs to be > 0")
	}
	G.L = make([]*Local, numLocal)
	for i := range G.L {
		G.L[i] = new(Local)
	}
	return G, nil
}

func NewLocal() *Local {
	L := new(Local)
	return L
}

func (S Serializer) Serialize() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(S); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Decode(b []byte) (Serializer, error) {
	M := make(map[string]interface{})
	dec := gob.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(&M); err != nil {
		return nil, err
	}
	return M, nil
}

func (G *Global) Set(S Serializer) error {
	for key, val := range S {
		switch v := val.(type) {
		case int:
			if key == "I" {
				G.I = v
			} else {
				return errors.New("Incorrect key: " + key)
			}
		case string:
			if key == "S" {
				G.S = v
			} else {
				return errors.New("Incorrect key: " + key)
			}
		case float64:
			if key == "F" {
				G.F = v
			} else {
				return errors.New("Incorrect key: " + key)
			}
		case []*Local:
			if key == "L" {
				G.L = v
			} else {
				return errors.New("Incorrect key: " + key)
			}
		default:
			return errors.New("Incorrect type for key: " + key)
		}
	}

	return nil
}

func (L *Local) Set(S Serializer) error {
	for key, val := range S {
		switch v := val.(type) {
		case int32:
			if key == "I" {
				L.I = v
			} else {
				return errors.New("Incorrect key: " + key)
			}
		case []byte:
			if key == "B" {
				L.B = v
			} else {
				return errors.New("Incorrect key: " + key)
			}
		default:
			return errors.New("Incorrect type for key: " + key)
		}
	}

	return nil
}
