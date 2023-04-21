package main

import (
	"fmt"
)

type errKind int

const (
	_ errKind = iota
	noHeader
	cantReadHeader
	invalidChunkLen
)

type WaveError struct {
	kind  errKind
	value int
	err   error
}

var (
	HeaderMissing      = WaveError{kind: noHeader}
	HeaderReadFailed   = WaveError{kind: cantReadHeader}
	InvalidChunkLength = WaveError{kind: invalidChunkLen}
)

func (e WaveError) Error() string {
	switch e.kind {
	case noHeader:
		return "no header"
	case cantReadHeader:
		return fmt.Sprintf("cant read header [%d]: %s", e.value, e.err.Error())
	case invalidChunkLen:
		return fmt.Sprintf("invalid chunk length: %d", e.value)
	}
	return fmt.Sprintf("unknown error [%d]: %s", e.kind, e.err.Error())
}

func (e WaveError) Unwrap() error {
	return e.err
}

func (e WaveError) with(val int) WaveError {
	return WaveError{
		kind:  e.kind,
		value: val,
		err:   e.err,
	}
}

func (e WaveError) from(pos int, err error) WaveError {
	return WaveError{
		kind:  e.kind,
		value: pos,
		err:   err,
	}
}

func main() {

	// lvl0Error := HeaderMissing
	// lvl1Error := lvl0Error.from(2, fmt.Errorf("err lvl 1"))
	// lvl2Error := lvl1Error.from(2, fmt.Errorf("err lvl 2"))

	// fmt.Println(lvl0Error)
	// fmt.Println(lvl1Error)
	// fmt.Println(lvl2Error)

	// fmt.Println(errors.Is(lvl0Error, HeaderMissing))
	// fmt.Println(errors.Is(lvl1Error, HeaderMissing))

	// y := WaveError{
	// 	kind:  cantReadHeader,
	// 	value: 0,
	// 	err:   nil,
	// }
	// z := WaveError{
	// 	kind:  invalidHeaderType,
	// 	value: 0,
	// 	err:   nil,
	// }
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(cantReadHeader)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from: ", err)
		}
	}()
	abc()
}

func abc() {
	panic("*_*")
}
