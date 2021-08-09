package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

// POST Body  json
func ParseData(r *http.Request, dest interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = UnMarshal(data, dest.(proto.Message))
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	return nil
}

// GET url params
func ParseParams(r *http.Request, dest interface{}) error {
	v := reflect.ValueOf(dest)
	t := reflect.TypeOf(dest)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("dest must is point type")
	}
	v = v.Elem()
	t = t.Elem()
	urlValues := r.URL.Query()
	for i := 0; i < v.NumField(); i++ {
		vfield := v.Field(i)
		tag := t.Field(i).Tag.Get("json")
		value := urlValues.Get(tag)
		switch vfield.Kind() {
		case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint, reflect.Uint8:
			valueuint64, err := strconv.ParseUint(value, 0, 10)
			if err != nil {
				return err
			}
			vfield.SetUint(valueuint64)
		case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int, reflect.Int8:
			valueuint64, err := strconv.ParseInt(value, 0, 10)
			if err != nil {
				return err
			}
			vfield.SetInt(valueuint64)
		case reflect.Float64, reflect.Float32:
			valueuint64, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			vfield.SetFloat(valueuint64)
		case reflect.String:
			vfield.SetString(value)
		case reflect.Bool:
			m := map[string]bool{
				"1":    true,
				"true": true,
			}
			vfield.SetBool(m[strings.ToLower(value)])
		default:
			log.Printf("unsupport type %T\n", vfield.Kind())
		}
	}
	return nil
}
