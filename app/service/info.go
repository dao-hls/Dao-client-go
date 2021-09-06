/*
@Time : 2021/8/3 4:44 下午
@Author : fushisanlang
@File : info.go
@Software: GoLand
*/
package service

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func Info() {
	fmt.Println("XX大陆，元气消散。\n只有最优秀的人才能飞升上界，逃出生天。\n无法飞升的人只能化为虚无，归还元气给天地。\n如今，是元气复苏的第一世。\n元气剩余-170039205份。\n努力修炼，争取飞升。")

}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it

	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested

		cmd.Stdout = os.Stdout

		cmd.Run()

	}

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested

		cmd.Stdout = os.Stdout

		cmd.Run()

	}

}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.

	if ok { //if we defined a clear func for that platform:

		value() //we execute it

	} else { //unsupported platform

		panic("Your platform is unsupported! I can't clear terminal screen :(")

	}

}

func Clean() {
	time.Sleep(2 * time.Second)
	CallClear()
}
