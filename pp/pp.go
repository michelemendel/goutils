package pp

import (
	"encoding/json"
	"fmt"
	"log"
)

func PP(s any) {
	res, err := PrettyStruct(s)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
