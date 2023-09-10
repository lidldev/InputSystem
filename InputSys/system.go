package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	username string
	fullname string
	age      int
	mail     string
}

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func (u User) UsernameUpdate() {
	fmt.Print("Enter Your Username: ")
	scanner.Scan()
	u.username = scanner.Text()
	if strings.Contains(u.username, "@") {
		panic("This Is Not Email!!")
	}
}
func (u User) FullnameUpdate() {
	arr := []string{"*", "?", "=", ")", "\",", ",", ":", ";", "+", ",", "@", "€", "ƒ", "{", ".", "]", "[", "}", "§", "½", "$", "#", "&", "£", ">", "∑", "≥", "≤", "µ", "~", "∫", "√", "≈", "Ω", "`", "æ", "´", "¬", "¨", "∆", "^", "ƒ", "∂", "ß", "æ", "~", "π", "¥", "₺", "®", "∑"}
	fmt.Print("Enter Your Full Name: ")
	scanner.Scan()
	u.fullname = scanner.Text()
	for _, e := range arr {
		if strings.Contains(u.fullname, e) {
			panic("Your Real Name Cannot Contain These Characters (Unless You're Elon Musk's Kid)")
		}
	}
}
func (u User) AgeUpdate() {
	fmt.Print("Enter Your Age: ")
	scanner.Scan()
	age, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic("hm")
	}
	u.age = int(age)
	if u.age < 18 {
		panic("Too Little")
	}
	if u.age > 100 {
		panic("Congratulations But Too Old")
	}
}
func (u User) MailUpdate() {
	fmt.Print("Enter Your Email: ")
	scanner.Scan()
	u.mail = scanner.Text()
	if strings.Contains(u.mail, "@") == false {
		panic("This Is Email Please Use '@' Sign")
	}
}
func main() {
	User.UsernameUpdate(User{})
	User.FullnameUpdate(User{})
	User.AgeUpdate(User{})
	User.MailUpdate(User{})
	fmt.Println("Welcome...")
}
