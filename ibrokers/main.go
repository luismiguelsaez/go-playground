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

	refCode = string(submatches[1])

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

	codeReq, err := getRefCode(token, "479814")
	if err != nil {
		fmt.Println("Got error while requesting reference code: " + err.Error())
	} else {
		fmt.Println(codeReq)
	}

	time.Sleep(5 * time.Second)

	queryData, err := getStatement(token, codeReq)
	if err != nil {
		fmt.Println("Got error while requesting query data: " + err.Error())
	} else {
		fmt.Println(queryData)
	}
}
