package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedValue, cacheHit := c.cache.Get(url)
	locationsResp := RespShallowLocations{}

	if cacheHit {
		err := json.Unmarshal([]byte(cachedValue), &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}

func (c *Client) Explore(locationName string) (RespLocation, error) {
	url := baseURL + "/location-area/" + locationName

	cachedValue, cacheHit := c.cache.Get(url)

	result := RespLocation{}

	if cacheHit {
		err := json.Unmarshal([]byte(cachedValue), &result)
		if err != nil {
			return RespLocation{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return RespLocation{}, err
	}
	c.cache.Add(url, data)

	return result, nil
}
