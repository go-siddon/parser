package parser

import (
	"errors"
	"reflect"
)

func EncodeModel(obj interface{}) (M, error) {
	var res M = M{}
	parsed, err := parse(databaseTag, obj)
	if err != nil {
		return nil, err
	}

	for _, p := range parsed {
		var val interface{}
		switch p.fieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = p.fieldValue.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = p.fieldValue.Uint()
		case reflect.Float32, reflect.Float64:
			val = p.fieldValue.Float()
		case reflect.Bool:
			val = p.fieldValue.Bool()
		case reflect.String:
			val = p.fieldValue.String()
		default:
			return nil, errors.New("case not provided")
		}

		if !p.fieldValue.IsZero() {
			res = append(res, E{Key: getFieldname(p.fieldTag), Value: val,
				Required: checkTag(p.fieldTag, propertyRequired),
				Index:    checkTag(p.fieldTag, propertyIndex),
				Unique:   checkTag(p.fieldTag, propertyUnique),
			})
		}

	}
	return res, nil
}
