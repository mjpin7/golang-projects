package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	This is the handlerFunc that will be returned as the mapHandler and used in ListenAndServe
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		/**
		* Handy syntax
		* Initializes, then checks for condition
		* When there is no corresponding value for that key, it will respond with a zero value and a false bool
		 */
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse the incoming yaml
	parsedYml, err := parseYaml(yml)

	if err != nil {
		panic(err)
	}

	// Build the map based off the parsed yaml
	pathMap := buildMap(parsedYml)

	// Return the handlerFunc that MapHandler creates
	return MapHandler(pathMap, fallback), nil
}

// buildMap will create a map out of the passed in array of YamlStuff
// It will then return that map
func buildMap(y []yamlStuff) map[string]string {
	pathMap := make(map[string]string)

	for _, entry := range y {
		pathMap[entry.Path] = entry.URL
	}

	return pathMap
}

// parseYaml will parse an inputted yaml byte array
// It will then return the decoded yaml in the form of YamlStuff array, otherwise error
func parseYaml(yml []byte) ([]yamlStuff, error) {
	var y []yamlStuff

	// Decodes the yaml into the array above
	err := yaml.Unmarshal(yml, &y)

	if err != nil {
		panic(err)
	}

	return y, nil
}

// YamlStuff is a struct for handling the yaml parsing
type yamlStuff struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
