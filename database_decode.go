package parser

import (
	"errors"
	"reflect"
	"strings"
)

func DecodeModel(obj interface{}, data M) error {
	v := reflect.ValueOf(obj)

	if v.Kind() != reflect.Pointer {
		return errors.New("expect a pointer to a struct")
	}

	p := v.Elem()
	for i := 0; i < p.NumField(); i++ {
		field := p.Field(i)
		if field.CanSet() && field.IsValid() {
			for _, d := range data {
				tags := strings.Split(p.Type().Field(i).Tag.Get(databaseTag), ",")
				tag := getFieldname(tags)

				if tag == d.Key {
					field.Set(reflect.ValueOf(d.Value))
				}
			}
		}
	}

	return nil
}
