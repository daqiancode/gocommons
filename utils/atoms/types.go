package atoms

import "time"

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type Float interface {
	~float32 | ~float64
}
type String interface {
	~string
}
type IntUint interface {
	Int | Uint
}

type IntUintFloat interface {
	Int | Uint | Float
}

type NumStr interface {
	Int | Uint | String
}

type Atom interface {
	~bool | Int | Uint | String
}

type Basic interface {
	Atom | time.Time
}

type JSON interface {
	Atom | time.Time
}
