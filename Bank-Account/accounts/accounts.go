package accounts

import (
	"errors"
	"fmt"
)

// Account stuct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw") // 에러를 변수에 저장할 수 있는데 이렇게 하려면 errFOO와 같은 네이밍을 줘라.

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 새로운 object를 만들어서 리턴한다.
} // go의 흔한 패턴, constructor가 없기에 function으로 만든다.

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // 이게 메소드다. 이렇게 함으로써 Account가 Deposit이라는 메소드를 갖게 되었다.(a Account)를 Receiver라고 부른다. receiver작성 규칙은 항상 struct의 첫글자를 따서지어야 한다는거다. 그래서 지금 a가 있는거다.
	a.balance += amount
	// (a *Account) 가 있지? 복사본을 받아오는게 아니라 실제로 Deposit을 호출한 Account를 사용하라는 뜻이다.
	// 이걸 pointer Receiver라고 말한다.
}

// Balance of your acount
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil // error에는 두개의 상태가 있다. null과 같은 nil과 if문에서 리턴시키는 error가 있다.
}

// go에는 try except 뭐 이런건 없다. catch도 없다. error를 그냥 return해서 거기서 처리해야 한다.
// error를 다룰 코드를 우리가 직접 만들어야 한다.

// ChangeOwner of the acount
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string { // account를 그냥 출력시키려고 하면 이런 String같은 메소드들이 자동으로 출력된다. 파이썬의 class를 print할경우 __str__출력과 비슷함.
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}

// method를 struct에만 추가할 수 있는게 아니다. type에도 추가가능하다? 호오
