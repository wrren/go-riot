package riot

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
	"encoding/json"
)

//
//	Make a HTTP GET request to the Riot API with the specified parameters. See the documentation for riot.URL to understand how
//	the given parameters affect the target URL. If the request is successful, the response is interpreted as JSON and unmarshaled
//	into the given value, v, as per the behaviour of the Go json.Unmarshal function.
//	
//	Return an error on failure or nil otherwise
//
func GetAndUnmarshal(region Region, static bool, endpoint string, version string, key string, path []string, parameters map[string]string, v interface{} ) error {
	response, err := Get( region, static, endpoint, version, key, pathm parameters )

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll( response.Body )

	if err != nil {
		return err
	}

	return json.Unmarshal( body, &v )
}

//
//	Make a HTTP GET request to the Riot API with the specified parameters. See the documentation for riot.URL to understand how
//	the given parameters affect the target URL.
//	
//	Returns the response to the GET request on success, with a non-nil error on failure
//
func Get(region Region, static bool, endpoint string, version string, key string, path []string, parameters map[string]string) (*http.Response, error) {
	Url, err := URL(region, static, endpoint, version, key, path, parameters)

	if err != nil {
		return nil, err
	}

	return http.Get(Url)
}

//
//	Form a URL for a Riot API endpoint based on the given parameters and options
//	region - Target API Region, see the constants defined in riot.go
//	static - Indicates whether data should be retrieved from the static data API, doesn't count towards API request limits if true
//	endpoint - Target endpoint name
//	version - Endpoint target version
//	key - API Key
//	path - Additional path variables for the request
//	parameters - GET parameters added to the request
//	
//	Returns a string containing the requested URL. Error will be non-nil if the URL could not be formed correctly with the given
//	parameters
//
func URL(region Region, static bool, endpoint string, version string, key string, path []string, parameters map[string]string) (string, error) {
	rawUrl := fmt.Sprintf("https://%s.api.pvp.net/api/lol/%s/v%s/%s", region, region, version, endpoint)

	if static {
		rawUrl = strings.Replace( rawUrl, "/api/lol/", "/api/lol/static-data/", 1 )
	}

	for _, part := range path {
		rawUrl += "/" + part
	}

	URL, err := url.Parse(rawUrl)

	if err != nil {
		return "", err
	}

	params := url.Values{}

	for key, value := range parameters {
		params.Add(key, value)
	}

	params.Add("api_key", key)
	URL.RawQuery = params.Encode()
	return URL.String(), nil
}
