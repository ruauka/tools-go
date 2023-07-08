// Package attrs_go - reflect helper.
package attrs_go //nolint:revive,stylecheck

import (
	"fmt"
	"math"
	"reflect"
)

// GetAttr - get struct field value.
// Args: obj, fieldName - value param. Struct fields can be ptr or value.
func GetAttr(obj interface{}, fieldName string) (interface{}, error) {
	// to reflect value
	objValue := reflect.ValueOf(obj)
	// is struct check
	if objValue.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}
	// get field value
	field := objValue.FieldByName(fieldName)
	// is field in struct
	if !field.IsValid() {
		return nil, ErrFieldNotInStruct
	}
	// is field exported
	if !field.CanInterface() {
		return nil, ErrUnexportedField
	}
	// field  ptr check
	if field.Kind() == reflect.Ptr {
		return field.Elem().Interface(), nil
	}

	return field.Interface(), nil
}

// SetAttr - set new value on structure field.
// Args: obj - ptr param. Struct fields can be ptr or value.
// Args: fieldName, newValue - value param.
func SetAttr(obj interface{}, fieldName string, newValue interface{}) error {
	var (
		// to reflect value
		objValue = reflect.ValueOf(obj)
		field    reflect.Value
	)
	// struct ptr check
	if objValue.Kind() != reflect.Ptr {
		return ErrNotPointer
	}
	// is struct check
	if objValue.Elem().Kind() != reflect.Struct {
		return ErrNotStruct
	}
	// get field value
	objValue = objValue.Elem().FieldByName(fieldName)
	// field  ptr check
	if objValue.Kind() == reflect.Ptr {
		field = objValue.Elem()
	} else {
		field = objValue
	}
	// is field in struct
	if !field.IsValid() {
		return ErrFieldNotInStruct
	}
	// types check
	if field.Type() != reflect.TypeOf(newValue) {
		return ErrWrongFieldValueType
	}
	// is field exported
	if !field.CanSet() {
		return ErrUnexportedField
	}
	// set value
	field.Set(reflect.ValueOf(newValue))

	return nil
}

// SetStructAttrs - updates current structure fields with the values of the new structure fields.
// Args: curObj - ptr param. Struct fields can be ptr or value.
// Args: newObj - value param. Struct fields can be ptr or value.
func SetStructAttrs(curObj, newObj interface{}) error {
	// to reflect value
	elem := reflect.ValueOf(newObj)
	// is struct check
	if elem.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	for i := 0; i < elem.NumField(); i++ {
		// get field reflect value
		field := elem.Field(i)
		// is ptr && is ptr nil check
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}
		// get newObj field name
		fieldName := elem.Type().Field(i).Name
		// get newObj field value
		fieldValue, err := GetAttr(newObj, fieldName)
		if err != nil {
			return fmt.Errorf("err in GetAttr: %w", err)
		}
		// get curObj field value
		if err := SetAttr(curObj, fieldName, fieldValue); err != nil {
			return fmt.Errorf("err in SetAttr: %w", err)
		}
	}

	return nil
}

func RoundUp(value float64, precision int) float64 {
	return math.Ceil(value*(math.Pow10(precision))) / math.Pow10(precision)
}

func iterRound(field reflect.Value, precision int, bitSize int) {
	for j := 0; j < field.Len(); j++ {
		field := field.Index(j)

		if bitSize == 64 {
			field.Set(reflect.ValueOf(RoundUp(field.Float(), precision)))
			continue
		}

		field.Set(reflect.ValueOf(float32(RoundUp(field.Float(), precision))))
	}
}

func RoundUpFloatFields(obj interface{}, precision int) error { //nolint:funlen
	objValue := reflect.ValueOf(obj)
	// struct ptr check
	if objValue.Kind() != reflect.Ptr {
		return ErrNotPointer
	}
	// is struct check
	if objValue.Elem().Kind() != reflect.Struct {
		return ErrNotStruct
	}

	objValue = objValue.Elem()

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		// is field exported
		if !field.CanSet() {
			return ErrUnexportedField
		}

		switch field.Kind() {
		case reflect.Float64, reflect.Float32:
			if field.Kind() == reflect.Float64 {
				field.Set(reflect.ValueOf(RoundUp(field.Float(), precision)))
				break
			}
			field.Set(reflect.ValueOf(float32(RoundUp(field.Float(), precision))))
		case reflect.Array, reflect.Slice:
			if field.Len() == 0 {
				break
			}
			if field.Index(0).Kind() == reflect.Float64 {
				iterRound(field, precision, 64)
				break
			}
			iterRound(field, precision, 32)
		}
	}

	return nil
}
