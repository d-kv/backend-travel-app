package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
)

func TestLatLngParsing(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		input        string
		outputLatLng *util.LatLng
	}{
		{"37.7749,-122.4194", &util.LatLng{37.7749, -122.4194}},
		{"39.9526,-75.1652", &util.LatLng{39.9526, -75.1652}},
		{"54.5260,-105.2551", &util.LatLng{54.5260, -105.2551}},
		{"51.5074,-0.1278", &util.LatLng{51.5074, -0.1278}},
	}

	for _, tc := range testCases {
		ll, err := util.NewLatLngFromString(tc.input)
		if err != nil {
			t.Errorf("Test case (%s) failed with error (%v)", tc.input, err)
			continue
		}
		assert.Equal(tc.outputLatLng, ll, "must be the same")
	}
}
