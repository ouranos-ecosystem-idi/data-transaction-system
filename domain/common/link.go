package common

import (
	"fmt"
	"net/url"
	"reflect"
)

type QueryParams interface{}

const (
	API_PATH    = "/api/v1/datatransport?dataTarget="
	LINK_FORMAT = "<%s>; rel=\"next\""
)

// CreateAfterLink
// Summary: This is function which creates the link with the after query parameter.
// input: host(string) host
// input: dataTarget(string) data target
// input: after(string) after
// input: params(QueryParams) query parameters
// output: (string) link
func CreateAfterLink(host string, dataTarget string, after string, params QueryParams) string {
	queryValues := url.Values{}
	link := host + API_PATH + dataTarget
	queryValues.Add("after", after)
	// query parameter handling
	if params != nil {
		v := reflect.ValueOf(params)
		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				field := v.Type().Field(i)
				jsonTag := field.Tag.Get("json")
				value := v.Field(i)

				// skip if json tag is not specified
				if jsonTag == "" {
					continue
				}

				// skip if the field is a pointer type and nil
				if field.Type.Kind() == reflect.Ptr && value.IsNil() {
					continue
				}

				// skip if the field is a string type and empty
				if field.Type.Kind() == reflect.String && value.String() == "" {
					continue
				}

				// If the field is a pointer type, get the value pointed to by the pointer
				if field.Type.Kind() == reflect.Ptr {
					value = value.Elem()
				}

				queryValues.Add(jsonTag, fmt.Sprintf("%v", value.Interface()))
			}
		}
	}

	link += "&" + queryValues.Encode()
	return fmt.Sprintf(LINK_FORMAT, link)
}
