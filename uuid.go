package main

import (
	"fmt"
	"strings"

	uuid1 "github.com/pborman/uuid"
)

func main() {
	uuidWithHyphen := uuid1.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
