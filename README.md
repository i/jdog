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

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
