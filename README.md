# 1000aaas.io

## Compiling

```bash
CGO_ENABLED=0 go build -a -installsuffix cgo main.go
```

This will give you the statically compiled binary `main`. The binary being statically compiled is important to keep the container... contained.


## Docker

The Dockerfile uses the base image [scratch](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)

It can be built as per:

```bash
docker build -t quay.io/jspc/1000aaas:latest .
```

And, obviously, lives on [quay](http://quay.io/jspc/1000aaas).

## Running

```bash
docker run -p 8000:8000 quay.io/jspc/1000aaas
```

## Licence

The MIT License (MIT)

Copyright (c) 2015, jspc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
