package attrs_go

import "errors"

var (
	ErrNotStruct           = errors.New("not a struct")
	ErrFieldNotInStruct    = errors.New("field not in struct")
	ErrUnexportedField     = errors.New("field not exported")
	ErrWrongFieldValueType = errors.New("wrong field value type")
	ErrNotPointer          = errors.New("struct passed not by pointer")
)
