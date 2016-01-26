package yelp

import (
	"errors"
)

// PhoneOptions provide the ability to search for a business by phone number.
type PhoneOptions struct {
	Phone    string // Parameter that specifies the business phone number to search for. Outside of the US and Canada, include the international dialing code (e.g. +442079460000) or use the 'cc' parameter
	CC       string // ISO 3166-1 alpha-2 country code. Default country to use when parsing the phone number. United States = US, Canada = CA, United Kingdom = GB (not UK).
	Category string // Category to filter search results with. See the list of supported categories.
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
// that match the defined values.
func (o *PhoneOptions) getParameters() (params map[string]string, err error) {
	params = make(map[string]string)
	if o.Phone == "" {
		return params, errors.New("to perform a search for a business by phone number, the phone property must be specified")
	}
	params["phone"] = o.Phone

	if o.CC != "" {
		params["cc"] = o.CC
	}

	if o.Category != "" {
		params["category"] = o.Category
	}

	return params, nil
}
