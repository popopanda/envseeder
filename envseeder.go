package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type UserData struct {
	Environment map[string]interface{} `json:"environment"`
}

func main() {
	url := "http://169.254.169.254/latest/user-data"
	filePath := "/etc/environment"
	var buffer bytes.Buffer

	result := parseUserData(url)

	for a, b := range result.Environment {
		outputString := fmt.Sprintf("%v=\"%v\"\n", strings.ToUpper(a), b)
		buffer.WriteString(outputString)
	}

	response := buffer.String()
	writeEnvFile(filePath, response)
}

func parseUserData(urlPath string) UserData {
	var userData UserData
	resp, err := http.Get(urlPath)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	yaml.Unmarshal(body, &userData)

	return userData
}

func writeEnvFile(path, response string) {
	stringByte := []byte(response)
	err := ioutil.WriteFile(path, stringByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
