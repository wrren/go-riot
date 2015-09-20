package riot

import "testing"

func TestURL(t *testing.T) {
	cases := []struct {
		region Region
		endpoint string
		version string
		key string
		parameters map[string]string
		want string
	}{
		{ "euw", "champion", "1.2", "1234", make(map[string]string), "https://euw.api.pvp.net/api/lol/euw/v1.2/champion?api_key=1234"},
	}

	for _, c := range cases {
		url, err := URL( c.region, c.endpoint, c.version, c.key, c.parameters )

		if err != nil {
			t.Errorf( "URL(%q, %q, %q, %q, %q) == %q, error %q", c.region, c.endpoint, c.version, c.key, c.parameters, url, err )
		}

		if url != c.want {
			t.Errorf( "URL(%q, %q, %q, %q, %q) == %q, want %q", c.region, c.endpoint, c.version, c.key, c.parameters, url, c.want )
		}
	}
}