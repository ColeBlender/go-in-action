package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func (u user) notify() {
	fmt.Println(u.name)
}

func main() {

	lisa := user{
		name:       "Lisa",
		email:      "dank@email.com",
		ext:        23,
		privileged: false,
	}

	lisa.notify()

	cole := user{"Cole", "coleblender@gmail.com", 23, true}

	fmt.Println(cole)
}
