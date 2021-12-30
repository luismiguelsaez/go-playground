package main

import (
	"errors"
	"io"
	"net/http"
	"regexp"
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
	submatches := re.FindStringSubmatch(string(httpBody))

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
