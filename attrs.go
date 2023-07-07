package attrs_go

import (
	"fmt"
	"reflect"
)

func SetAttr(obj interface{}, fieldName string, newValue interface{}) error {
	var (
		objValue = reflect.ValueOf(obj)
		field    reflect.Value
	)

	if objValue.Kind() != reflect.Ptr {
		return ErrNotPointer
	}

	if objValue.Elem().Kind() != reflect.Struct {
		return ErrNotStruct
	}

	objValue = objValue.Elem().FieldByName(fieldName)
	if objValue.Kind() == reflect.Ptr {
		field = objValue.Elem()
	} else {
		field = objValue
	}

	if !field.IsValid() {
		return ErrFieldNotInStruct
	}

	if field.Type() != reflect.TypeOf(newValue) {
		return ErrWrongFieldValueType
	}

	if !field.CanSet() {
		return ErrUnexportedField
	}

	field.Set(reflect.ValueOf(newValue))

	return nil
}

// GetAttr - get struct field value.
func GetAttr(obj interface{}, fieldName string) (interface{}, error) {
	objValue := reflect.ValueOf(obj)

	if objValue.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	field := objValue.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, ErrFieldNotInStruct
	}

	if !field.CanInterface() {
		return nil, ErrUnexportedField
	}

	if field.Kind() == reflect.Ptr {
		return field.Elem().Interface(), nil
	}

	return field.Interface(), nil
}

func SetStructAttrs(curObj, newObj interface{}) error {
	elem := reflect.ValueOf(newObj)

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		fieldName := elem.Type().Field(i).Name

		fieldValue, err := GetAttr(newObj, fieldName)
		if err != nil {
			return fmt.Errorf("err in GetAttr: %v", err)
		}

		if err := SetAttr(curObj, fieldName, fieldValue); err != nil {
			return fmt.Errorf("err in SetAttr: %v", err)
		}
	}

	return nil
}
