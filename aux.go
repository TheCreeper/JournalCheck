package main

import (
	"os"
)

func GetHostName() string {

	n, e := os.Hostname()
	if e != nil {

		n = "unknown"
	}
	return n
}
