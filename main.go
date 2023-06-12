package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	run1()
	// run3()
}

func run3() {
	scanner := bufio.NewScanner(os.Stdin)

	type Prompt struct {
		Key       string `json:"key"`
		PromptStr string
	}

	prompts := []Prompt{
		{
			Key:       "Name",
			PromptStr: "What is your name?",
		},
		{
			Key:       "Age",
			PromptStr: "What is your age?",
		},
	}

	for _, p := range prompts {
		fmt.Println(">", p.PromptStr)
		scanner.Scan()
		line := scanner.Text()
		fmt.Printf("%s is %v\n\n", p.Key, line)
	}
}

type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func payload(payload any) *io.PipeReader {
	pr, pw := io.Pipe()
	go func() {
		json.NewEncoder(pw).Encode(payload)
		defer pw.Close()
	}()
	return pr
}

func run1() {
	resp, err := http.Post("https://httpbin.org/post", "application/json",
		payload(Payload{
			Name: "John",
			Age:  30,
		}))

	// var pl *bytes.Buffer
	// io.Writer
	// json.NewEncoder(pl).Encode(Payload{Name: "John", Age: 30})
	//  io.Reader
	// resp, err := http.Post("https://httpbin.org/post", "application/json", pl)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	res, _ := io.ReadAll(resp.Body)
	fmt.Println(string(res))
}
