package main

import "fmt"

func Login(username, password string) bool{
	if username == "admin" && password == "123" {
		fmt.Println("Login Berhasil")
		return true
	}
	fmt.Println("Login gagal")
	return false
}