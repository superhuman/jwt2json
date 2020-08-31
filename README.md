jwt2json takes a JWT in stdin, and outputs the claims as JSON on stdout.

It does not validate the header, or check the signature.

### Usage

If you have a functional go environment, you can use jwt2json like so:

```
$ pbpaste | go run github.com/superhuman/jwt2json
```

To save some typing you may find it better to first install `jwt2json` once, and then use it:

```
$ go get github.com/superhuman/jwt2json
```

From then on, 
```
$ pbpaste | jwt2json
```

### License (MIT)

Copyright 2020 Conrad Irwin <conrad@superhuman.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
