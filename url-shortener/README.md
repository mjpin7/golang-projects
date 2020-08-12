# URL Shortener

This program is a very simple url shortener/redirector. It will run on localhost:8080, and will listen for any url paths.

When a new path is recieved in the url It will first take in a map of strings, with the keys corresponding to the paths and the values corresponding to the redirect url. For example

```
{
    "google": "https://google.ca",
    "go-docs": "https://godoc.org"
}
```

It will then create a new http.HandlerFunc called `mapHandler` that checks if the requested path is in the created map and redirect to the value (for example localhost:8080/google will redirect to https://google.ca).

The program will then create some yaml string in the format

```
- path: /google
  url: https://google.ca
- path: /go-docs
  url: https://godoc.org
```

and create a http.HandlerFunc "yamlHandler" to try and do the same thing but with yaml text. The yamlHandler will first parse/decode the yaml, build a map of the decoded yaml and then return a new mapHandler.


If it could not find any matches in the yaml text, it will fallback to the mapHandler originally created that corresponds to the created map called `pathsToUrls` in main.go. If all else fails, it will fallback on the default mux and show the main page.

## Future Enhancements
* Allow for addition of new url redirects
* Read from yaml file instead of variable (from flag)

