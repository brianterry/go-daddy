package countries

import (
	"encoding/json"

	godaddy "github.com/alyx/go-daddy"
)

// Get retrieves summary country information for the provided marketID and
// filters.
func Get(c *godaddy.Client, marketID string, regionTypeID int, regionName string, sort string, order string) ([]CountrySummary, error) {
	var res []CountrySummary
	uri, err := godaddy.BuildQuery("/v1/countries", map[string]interface{}{
		"marketId":     marketID,
		"regionTypeId": regionTypeID,
		"regionName":   regionName,
		"sort":         sort,
		"order":        order,
	})
	if err != nil {
		return res, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// GetByKey retrieves country and summary state information for provided
// countryKey
func GetByKey(c *godaddy.Client, countryKey string, marketID string, sort string, order string) ([]Country, error) {
	var res []Country
	uri, err := godaddy.BuildQuery("/v1/countries/"+countryKey, map[string]interface{}{
		"marketId": marketID,
		"sort":     sort,
		"order":    order,
	})
	if err != nil {
		return res, err
	}

	data, err := c.Get(uri)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
