package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetHostName() string {

	n, e := os.Hostname()
	if e != nil {

		n = "unknown"
	}
	return n
}
