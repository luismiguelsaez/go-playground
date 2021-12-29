package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {

	token := os.Getenv("IB_TOKEN")

	var rcRequest string = "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.SendRequest?t=" + token + "&q=479814&v=3"
	var refCode string = ""
	fmt.Printf("Obtaining reference code from %s\n", rcRequest)

	resp, err := http.Get(rcRequest)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	}

	fmt.Printf("Got HTTP status code %v\n", resp.StatusCode)

	httpBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Got error while ready body contents: %s", err)
	}

	re := regexp.MustCompilePOSIX("<ReferenceCode>(.*)</ReferenceCode>")
	submatches := re.FindSubmatch(httpBody)

	refCode = string(submatches[1])

	time.Sleep(5 * time.Second)

	stRequest := "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.GetStatement?t=" + token + "&&v=3&q=" + string(refCode)
	fmt.Printf("Getting flex query: %s\n", stRequest)
	respst, errst := http.Get(stRequest)
	if errst != nil {
		fmt.Printf("Got error: %v\n", errst)
	}

	fmt.Printf("Got HTTP status code %v\n", respst.StatusCode)

	sthttpBody, err := io.ReadAll(respst.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Got error while ready body contents: %s", err)
	}

	fmt.Printf("Query body: %s", sthttpBody)
}
