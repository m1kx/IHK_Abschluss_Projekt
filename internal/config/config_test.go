package config

import (
	"reflect"
	"testing"
)

func TestAllFieldsSet(t *testing.T) {
	s, err := New("../../config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	assertSet(t, "Config", reflect.ValueOf(s.Get()))
}

func assertSet(t *testing.T, path string, v reflect.Value) {
	t.Helper()
	if v.Kind() == reflect.Struct {
		for i := range v.NumField() {
			assertSet(t, path+"."+v.Type().Field(i).Name, v.Field(i))
		}
		return
	}
	if v.IsZero() {
		t.Errorf("%s is not set", path)
	}
}
