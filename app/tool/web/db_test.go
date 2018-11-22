package web

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDB_CreateOrModify(t *testing.T) {

	ip := "1.1.1.1"
	reg := regexp.MustCompile(`((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))`)
	fmt.Println(reg.MatchString(ip))

}
