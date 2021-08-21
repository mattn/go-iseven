package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/samuelmarina/is-even/main/index.js")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	content := string(b)
	content = strings.Replace(content, `function isEven(number) {`, "package iseven\n\nfunc IsEven(number interface{}) bool {", 1)
	content = strings.Replace(content, `module.exports = isEven;`, ``, -1)
	content = regexp.MustCompile(`number === ([0-9]+|"[^"]+")`).ReplaceAllString(content, "is(number, $1)")
	content = regexp.MustCompile(`return (true|false);\n`).ReplaceAllString(content, "{\n\t\treturn $1\n\t}")
	content = strings.Replace(content, `}}`, "}\n\treturn false\n}", -1)
	f, err := os.Create("is-even.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte(content))
}
