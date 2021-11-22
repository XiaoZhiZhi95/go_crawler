package dao

import (
	"fmt"
	"testing"
)

func TestCheckUserNameAndPassword(t *testing.T) {
	userName := "lili"
	passwd := "123456"
	user,err := CheckUserNameAndPassword(userName, passwd)
	fmt.Println(user, err)
}

func TestCheckUserName(t *testing.T) {
	username := "lili"
	user, err := CheckUserName(username)
	fmt.Println(user, err)
}

func TestSaveUser(t *testing.T) {
	uname := "zhizhi"
	passwd := "fz..555FZ"
	em := "zhizhi@lili.com"
	err := SaveUser(uname, passwd, em)
	fmt.Println("保存", err)
}