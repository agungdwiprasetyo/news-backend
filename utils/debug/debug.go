package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func PrettyJSON(data interface{}) string {
	buff, _ := json.Marshal(data)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, buff, "", "\t")
	// log.Printf("%s: %s", name, string(prettyJSON.Bytes()))
	return fmt.Sprintf("\x1b[32;1m%s\x1b[0m", string(prettyJSON.Bytes()))
}

func PrintError(data ...interface{}) {
	var arr []string
	for _, val := range data {
		arr = append(arr, fmt.Sprint(val))
	}
	str := strings.Join(arr, " >>> ")
	fmt.Printf("\x1b[31;1m%v\x1b[0m\n", str)
}

func PrintRed(data ...interface{}) {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	fmt.Printf("\x1b[31;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func PrintGreen(data ...interface{}) {
	var str []string
	for _, val := range data {
		str = append(str, fmt.Sprint(val))
	}
	fmt.Printf("\x1b[32;1m%v\x1b[0m\n", strings.Join(str, " >> "))
}

func PrintYellow(data interface{}) {
	fmt.Printf("\x1b[33;1m%v\x1b[0m\n", data)
}
