package main

var users []string

func main() {
	AddUser("Alice")
	AddUser("Bob")
}

func AddUser(a string) {
	users = append(users, a)
}
