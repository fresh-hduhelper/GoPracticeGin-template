package main

import (
	"fmt"
	"os"
)

var chaptersChanged = make(chan struct{}, 1)

func main() {
	op := checkOption()
	goCheck(op)
}

func goCheck(op string) {
	if checkFn, exist := availablePractices[op]; !exist {
		fmt.Println("unsupported practice")
		os.Exit(-1)
	} else {
		checkFn()
	}
}

func checkOption() string {
	if v := os.Getenv("CHECK_TARGET"); v != "" {
		return v
	}
	return os.Args[1]
}

var availablePractices = map[string]func(){
	"authority/":     check01,
	"common-params/": check02,
	"configuration/": check03,
	"custom-server/": check04,
	"database-sql/":  check05,
	"middlewire/":    check06,
	"project-layer/": check07,
	"template-mvc/":  check08,
}

// authority/
func check01() {}

// common-params/
func check02() {}

// configuration/
func check03() {}

// custom-server/
func check04() {}

// database-sql/
func check05() {}

// middlewire/
func check06() {}

// project-layer/
func check07() {}

// template-mvc/
func check08() {}
