# logger
Package for the error log

## Get Started

```sh
go get -u github.com/mvochoa/logger
```

- [API docs](https://godoc.org/github.com/mvochoa/logger)

You can change the folder where the records are stored, declared the environment variable `DIR_LOG`.

## Usage

When calling the Error function, this will print the error in the console and it will also record the error in an error.log file

```golang
package main

import (
    "errors"

    "github.com/mvochoa/logger"
)

func main() {
    logger.Error("TAG", errors.New("error"))
}
```

Output console:

```
========

[2019/04/21 11:15:31] ERROR/TAG: error

========
```

Contents of the file `log/error_0000_00_00.log`

```
[2019/04/21 11:15:31] ERROR/TAG: error
```