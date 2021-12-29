package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func getRefCode(token string, queryId string) (string, error) {

	var rcRequest string = "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.SendRequest?t=" + token + "&q=" + queryId + "&v=3"
	var refCode string = ""

	resp, err := http.Get(rcRequest)
	if err != nil {
		return "", errors.New("error while sending request: " + err.Error())
	}

	httpBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", errors.New("error while reading response body: " + err.Error())
	}

	re := regexp.MustCompilePOSIX("<ReferenceCode>(.*)</ReferenceCode>")
	submatches := re.FindSubmatch(httpBody)

	if len(submatches) > 1 {
		refCode = string(submatches[1])
	} else {
		return "", errors.New("no matching reference code in response body: " + err.Error())
	}

	return refCode, nil
}

func getStatement(token string, refCode string) (string, error) {
	var stRequest string = "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.GetStatement?t=" + token + "&&v=3&q=" + refCode

	resp, err := http.Get(stRequest)
	if err != nil {
		return "", errors.New("error while sending request: " + err.Error())
	}

	httpBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", errors.New("error while reading response body: " + err.Error())
	}

	return string(httpBody), nil
}

func main() {

	token := os.Getenv("IB_TOKEN")
	queryId := os.Getenv("IB_QUERY_ID")

	codeReq, err := getRefCode(token, queryId)
	if err != nil {
		fmt.Println("Got error while requesting reference code: " + err.Error())
		os.Exit(1)
	}

	time.Sleep(5 * time.Second)

	queryData, err := getStatement(token, codeReq)
	if err != nil {
		fmt.Println("Got error while requesting query data: " + err.Error())
	} else {
		fmt.Println(queryData)
	}
}
