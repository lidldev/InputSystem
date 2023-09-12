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
	tries := 0
	for {
		fmt.Print("Enter Your Username: ")
		scanner.Scan()
		u.username = scanner.Text()
		if strings.Contains(u.username, "@") {
			tries++
			if tries == 3 {
				panic("This Is Not An Email")
			}
			if tries == 1 {
				fmt.Println("Wrong Entry You Have '2' Tries Left Please Dont Use '@' Sign")
			}
			if tries == 2 {
				fmt.Println("Wrong Entry You Have '1' Try Left")
			}
		}
	}
}
func (u User) FullnameUpdate() {
	tries := 0
	for {
		arr := []string{"*", "?", "=", ")", "\",", ",", ":", ";", "+", ",", "@", "€", "ƒ", "{", ".", "]", "[", "}", "§", "½", "$", "#", "&", "£", ">", "∑", "≥", "≤", "µ", "~", "∫", "√", "≈", "Ω", "`", "æ", "´", "¬", "¨", "∆", "^", "ƒ", "∂", "ß", "æ", "~", "π", "¥", "₺", "®", "∑"}
		fmt.Print("Enter Your Full Name: ")
		scanner.Scan()
		u.fullname = scanner.Text()
		for _, e := range arr {
			if strings.Contains(u.fullname, e) {
				tries++
				if tries == 3 {
					panic("Your Real Name Cannot Contain These Characters (Unless You're Elon Musk's Kid)")
				}
				if tries == 2 {
					fmt.Println("You Have '1' Tries Left")
				}
				if tries == 1 {
					fmt.Println("You Have '2' Tries Left")
				}
			}
		}
	}
}
func (u User) AgeUpdate() {
	c := 0
	for {
		fmt.Print("Enter Your Age: ")
		scanner.Scan()
		age, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			c++
			if c == 3 {
				panic("Are You Sure You Entered A Number??")
			}
			if c == 2 {
				fmt.Println("You Have '1' Tries Left")
			}
			if c == 3 {
				fmt.Println("You Have '2' Tries Left")
			}
		}
		u.age = int(age)
		if u.age < 18 {
			panic("Too Little")
		}
		if u.age > 100 {
			panic("Congratulations But Too Old")
		}
	}
}
func (u User) MailUpdate() {
	c := 0
	for {
		fmt.Print("Enter Your Email: ")
		scanner.Scan()
		u.mail = scanner.Text()
		if strings.Contains(u.mail, "@") == false {
			c++
			if c == 3 {
				panic("This Is Email Please Use '@' Sign")
			}
			if c == 2 {
				fmt.Println("You Have '1' Tries Left")
			}
			if c == 1 {
				fmt.Println("You Have '2' Tries Left")
			}
		}
	}
}
func main() {
	User.UsernameUpdate(User{})
	User.FullnameUpdate(User{})
	User.AgeUpdate(User{})
	User.MailUpdate(User{})
	fmt.Println("Welcome...")
}
