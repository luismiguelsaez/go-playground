package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func getRefCode(token string, queryId string) (string, error) {

	var rcRequest string = "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.SendRequest?t=" + token + "&q=" + queryId + "&v=3"
	var refCode string = ""

	resp, err := http.Get(rcRequest)
	if err != nil {
		return "", errors.New("Error while sending request")
	}

	httpBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", errors.New("Error while reading response body")
	}

	re := regexp.MustCompilePOSIX("<ReferenceCode>(.*)</ReferenceCode>")
	submatches := re.FindSubmatch(httpBody)

	refCode = string(submatches[1])

	return refCode, nil
}

func main() {

	token := os.Getenv("IB_TOKEN")

	codeReq, err := getRefCode(token, "479814")
	if err != nil {
		fmt.Println("Got error while requesting reference code")
	} else {
		fmt.Println(codeReq)
	}

	//time.Sleep(5 * time.Second)
	//
	//stRequest := "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.GetStatement?t=" + token + "&&v=3&q=" + string(refCode)
	//fmt.Printf("Getting flex query: %s\n", stRequest)
	//respst, errst := http.Get(stRequest)
	//if errst != nil {
	//	fmt.Printf("Got error: %v\n", errst)
	//}
	//
	//fmt.Printf("Got HTTP status code %v\n", respst.StatusCode)
	//
	//sthttpBody, err := io.ReadAll(respst.Body)
	//resp.Body.Close()
	//if err != nil {
	//	fmt.Printf("Got error while ready body contents: %s", err)
	//}
	//
	//fmt.Printf("Query body: %s", sthttpBody)
}
