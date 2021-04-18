package main

import (
	"fmt"
	"strings"
	"user/command"
	_ "user/init"
)

// var commands = map[string]func(){
// 	"1": usermanage.Add,
// 	"2": usermanage.Delete,
// 	"3": usermanage.Modify,
// 	"4": usermanage.Query}

var commands = command.Commands

const (
	user     = "admin"
	password = "57ef24509317ddf05654dcd597278aed" //123456
)

func Login() {

}

//获取用户输入信息
func input(msg string) (value string) {
	fmt.Print(msg)
	fmt.Scanln(&value)
	return strings.TrimSpace(value)
}

//打印提示信息
func prompt() {
	propmtInfo := command.Prompt

	fmt.Println(strings.Repeat("*", 15))
	fmt.Println("欢迎使用用户管理系统")
	fmt.Println("1 退出")
	for _, v := range propmtInfo {
		fmt.Printf("%s %s\n", v[0], v[1])
	}
}

func main() {
END:
	for {
		prompt()
		cmd := input("请输入指令:")

		if command, ok := commands[cmd]; ok {
			command()
		} else if cmd == "1" {
			break END
		} else {
			fmt.Println("指令错误")
		}
	}
}
