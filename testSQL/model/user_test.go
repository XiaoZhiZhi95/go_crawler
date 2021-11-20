package model

import (
	"fmt"
	"testing"
)

//单元测试，使用go test
//go test -v，查看更详细的测试信息
//需要与测试的韩式放在同一个model下，且需要以TestXxx命名
func TestUser_AddUser1(t *testing.T) {
	fmt.Println("测试添加用户")
	//user := &User{}

	//调用添加用户的方法
	//user.AddUser1()
	//user.AddUser2()
	fmt.Println("添加")
}

//用子测试函数进行测试
func TestUseSun(t *testing.T)  {
	fmt.Println("开始测试子测试函数")
	t.Run("测试子函数", testSun)
}

func testSun(t *testing.T)  {
	fmt.Println("这是子函数")
}

//testMain函数可以在测试函数执行前做一些其他的操作
func TestMain(m *testing.M)  {
	fmt.Println("这是Beginning")
	//通过m.Run来开始执行后面的测试函数
	m.Run()
}