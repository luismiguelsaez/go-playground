package main

import (
	"fmt"
	"os"
	"time"
)

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
