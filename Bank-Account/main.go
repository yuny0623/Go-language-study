package main

import (
	"fmt"

	"github.com/yuny0623/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	//err := account.Withdraw(20) // 에러를 강제하는 이러한 방식이 더 좋을 수도 있다.
	//if err != nil {
	//log.Fatalln(err) // 프로그램을 종료시켜준다. Println을 하고 종료시킨다.
	//fmt.Println(err)
	//}
	fmt.Println(account)
}

// 소문자가 private이고 대문자가 public이다.

// struct가 갖는 메소드도 있고
