package riot

import "testing"

func TestURL(t *testing.T) {
	cases := []struct {
		region     Region
		static 	   bool
		endpoint   string
		version    string
		key        string
		path 	[]string
		parameters map[string]string
		want       string
	}{
		{"euw", false, "champion", "1.2", "1234", make([]string, 0, 0), make(map[string]string), "https://euw.api.pvp.net/api/lol/euw/v1.2/champion?api_key=1234"},
		{"euw", false, "champion", "1.2", "1234", []string{"20"}, make(map[string]string), "https://euw.api.pvp.net/api/lol/euw/v1.2/champion/20?api_key=1234"},
		{"euw", false, "champion", "1.2", "1234", []string{"20"}, map[string]string{"foo": "bar"}, "https://euw.api.pvp.net/api/lol/euw/v1.2/champion/20?api_key=1234&foo=bar"},
		{"euw", false, "champion", "1.2", "1234", []string{"20", "name"}, make(map[string]string), "https://euw.api.pvp.net/api/lol/euw/v1.2/champion/20/name?api_key=1234"},
	}

	for _, c := range cases {
		url, err := URL(c.region, c.static, c.endpoint, c.version, c.key, c.path, c.parameters)

		if err != nil {
			t.Errorf("URL(%q, %q, %q, %q, %q, %q, %q) == %q, error %q", c.region, c.static, c.endpoint, c.version, c.key, c.path, c.parameters, url, err)
		}

		if url != c.want {
			t.Errorf("URL(%q, %q, %q, %q, %q, %q, %q) == %q, want %q", c.region, c.static, c.endpoint, c.version, c.key, c.path, c.parameters, url, c.want)
		}
	}
}
