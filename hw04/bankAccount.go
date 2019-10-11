package main

import "fmt"

// in order to pass a function as an argument, define a specific type of function
type wd func(int, string) int

// static return type -> communications are printed, not returned
func makeWithdraw(balance int, password string) wd {
	var attempts []string
	fmt.Printf("Withdraw account with balance %d is created.\n", balance)
	return func(amount int, inp_pwd string) int {
		if len(attempts) >= 3 {
			fmt.Printf("Your account is locked. Attempts: %s\n", attempts)
			return -1
		}
		if inp_pwd != password {
			attempts = append(attempts, inp_pwd)
			fmt.Println("Incorrect password")
			return -1
		}
		if amount > balance {
			fmt.Println("Insufficient funds")
			return -1
		}
		balance -= amount
		fmt.Println(balance)
		return balance
	}
}

func makeJoint(withdraw wd, oldPassword string, newPassword string) wd {
	info := withdraw(0, oldPassword)
	if info == -1 {
		return withdraw
	}
	fmt.Printf("Joint account with balance %d is created.\n", info)
	return func(amount int, password string) int {
		if password == newPassword {
			return withdraw(amount, oldPassword)
		}
		return withdraw(amount, password)
	}
}

func main() {
	w := makeWithdraw(100, "hax0r")
	// Withdraw account with balance 100 is created.
	w(25, "hax0r")
	// 75
	w(90, "hax0r")
	// Insufficient funds
	w(25, "hwat")
	// Incorrect password
	w(25, "hax0r")
	// 50
	w(75, "a")
	// Incorrect password
	w(10, "hax0r")
	// 40
	w(20, "n00b")
	// Incorrect password
	w(10, "hax0r")
	// Your account is locked. Attempts: [hwat a n00b]
	w(10, "l33t")
	// Your account is locked. Attempts: [hwat a n00b]
	fmt.Print("\n")

	w = makeWithdraw(100, "hax0r")
	// Withdraw account with balance 100 is created.
	w(25, "hax0r")
	// 75
	makeJoint(w, "my", "secret")
	// Incorrect password
	j := makeJoint(w, "hax0r", "secret")
	// 75
	// Joint account with balance 75 is created.
	w(25, "secret")
	// Incorrect password
	j(25, "secret")
	// 50
	j(25, "hax0r")
	// 25
	j(100, "secret")
	// Insufficient funds
	j2 := makeJoint(j, "secret", "code")
	// 25
	// Joint account with balance 25 is created.
	j2(5, "code")
	// 20
	j2(5, "secret")
	// 15
	j2(5, "hax0r")
	// 10
	j2(25, "password")
	// Incorrect password
	j2(5, "secret")
	// Your account is locked. Attempts: [my secret password]
	j(5, "secret")
	// Your account is locked. Attempts: [my secret password]
	w(5, "hax0r")
	// Your account is locked. Attempts: [my secret password]
	makeJoint(w, "hax0r", "hello")
	// Your account is locked. Attempts: [my secret password]

}
