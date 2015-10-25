# jdog

dig through piles of json with ease
[jdog GoDoc](https://godoc.org/github.com/i/jdog)

## usage

```go
var m interface{}
json.Unmarshal([]byte(`{
  "foo": "bar",
  "wow": [
    "cool",
    {
      "sweet": 9
    }
  ]
}`), &m)

fmt.Println(jdog.Get(m, "foo"))          // bar <nil>
fmt.Println(jdog.Get(m, "wow[1].sweet")) // 9 <nil>
```

## Contributing

Feel free to contribute. Be sure to `gofmt`, `golint`, and `govet` your code and add tests.

## License

Copyright (c) Ian Lozinski.
