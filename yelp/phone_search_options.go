package yelp

import (
	"reflect"
)

// PhoneSearchOptions are the top level search parameters used for performing phone searches.
// You can define multiple sets of options, and use them together.
type PhoneSearchOptions struct {
	PhoneOptions *PhoneOptions // phone search options
}

// Generate a map that contains the querystring parameters for
// all of the defined options.
func (o *PhoneSearchOptions) getParameters() (params map[string]string, err error) {
	// create an empty map of options
	params = make(map[string]string)

	// reflect over the properties in o, adding parameters to the global map
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		if !val.Field(i).IsNil() {
			o := val.Field(i).Interface().(OptionProvider)
			fieldParams, err := o.getParameters()
			if err != nil {
				return params, err
			}
			for k, v := range fieldParams {
				params[k] = v
			}
		}
	}
	return params, nil
}
