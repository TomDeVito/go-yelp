package yelp

import (
	"testing"
)

// TestPhoneOptions will get a business by a valid phone number
func TestPhoneOptions(t *testing.T) {
	client := getClient(t)
	options := PhoneSearchOptions{
		PhoneOptions: &PhoneOptions{
			phone: "+12012221998",
		},
	}

	result, err := client.GetBusinessByPhoneSearch(options)
	check(t, err)

	assert(t, len(result.Businesses) > 0, containsResults)
}
