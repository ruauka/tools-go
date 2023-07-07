// Package attrs_go - reflect helper
package attrs_go //nolint:revive,stylecheck

import (
	"fmt"
	"reflect"
)

// GetAttr - get struct field value.
// Struct fields can be ptr or value
// Args: obj, fieldName, newValue - value param.
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

// SetAttr - setting the structure field.
// Struct fields can be ptr or value
// Args: obj - ptr param.
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

// SetStructAttrs - updates the fields of the current structure with the values of the fields of the new structure.
// Struct fields can be ptr or value
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
