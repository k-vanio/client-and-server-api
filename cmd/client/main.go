package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Bid string `json:"bid"`
}

func main() {
	client := http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("err: new request:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err: send request:", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: read body", err)
		return
	}

	var quote Quote
	if err := json.Unmarshal(data, &quote); err != nil {
		fmt.Println("Err: parse json", err)
		return
	}

	fmt.Printf("quote: %v\n", quote.Bid)

	f, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println("Err: save cotacao.txt:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("quote: %v", quote.Bid))
	if err != nil {
		fmt.Println("Err: writer cotacao.txt:", err)
		return
	}
}
