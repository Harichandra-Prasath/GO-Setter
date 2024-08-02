package main

import "fmt"

func getRContent(title string) []byte {
	return []byte(fmt.Sprintf("# %s", title))
}

func getMContent(title string) []byte {
	return []byte(fmt.Sprintf(
		`run: build
	@./bin/%s

build:
	@go build -o ./bin/%s`,

		title, title))
}

func getIcontent() []byte {
	return []byte("bin\n")
}
