/*
Package for the error log

Usage

When calling the Error function, this will print the error in the console and it will also record the error in an error.log file

	package main

	import (
		"errors"

		"github.com/mvochoa/logger"
	)

	func main() {
		logger.Error("TAG", errors.New("error"))
	}

Output console:

	========

	[2019/04/21 11:15:31] ERROR/TAG: error

	========

Contents of the file log/error_0000_00_00.log

	[2019/04/21 11:15:31] ERROR/TAG: error

*/
package logger
