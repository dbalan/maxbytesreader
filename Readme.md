# MaxBytesReader

Throw an error if reader response is bigger than a threshould.

example:

```golang
func Get(rawurl string, maxsize int64) (r *http.Response, err error) {
	r, err = http.Get(rawurl)
	if err != nil {
		return nil, err
	}

	r.Body = NewMaxBytesReader(r.Body, maxsize)
	return r, nil
}
```

The code is from:
 - http://grokbase.com/p/gg/golang-nuts/1517ny9kdm/go-nuts-limiting-size-of-http-response-body
