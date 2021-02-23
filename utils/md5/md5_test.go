package md5

import (
	"fmt"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	fmt.Println(EncryptPassword("123"))
}
