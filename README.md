# tern

[![Build Status](https://travis-ci.org/arteev/tern.svg?branch=master)](https://travis-ci.org/arteev/tern)
[![Coverage Status](https://coveralls.io/repos/arteev/tern/badge.svg?branch=master&service=github)](https://coveralls.io/github/arteev/tern?branch=master)

Description
-----------

Package Golang with ternary operator

Installation
------------

This package can be installed with the go get command:

    go get github.com/arteev/tern

Documentation
-------------
Example:

```go
   package main
   import (
   	"fmt"
   	"github.com/arteev/tern"
   )
   func main() {
   	fmt.Println(tern.Op(true,"It's good:)","Oh how bad"))
   }
```

License
-------

  MIT


Author
------

Arteev Aleksey