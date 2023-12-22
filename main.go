package main

import (
	"fmt"
	"os"
	"os/exec"
)

type User struct {
	Id       int
	Username string
	Password string
	Name     string
}

type Posts struct {
	Id     int
	UserId int
	Post   string
}

type PostComment struct {
	PostId   int
	FriendId int
	Comment  string
}

type Friends struct {
	UserId   int
	FriendId int
}

var users []User
var posts []Posts
var postComments []PostComment

var isLogin bool = false

func main() {
	for !isLogin {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			isLogin = userLogin()
		case 2:
			userRegister()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silahkan masukkan opsi yang valid.")
		}

		if isLogin {
			clearScreen()
			home()
		}
	}
}

func userRegister() {
	var newUser User

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&newUser.Username)

		if isUsernameExists(newUser.Username) {
			fmt.Println("Username sudah terpakai. Silahkan masukkan username yang lain.")
		} else {
			break
		}
	}

	fmt.Print("Masukkan password: ")
	fmt.Scan(&newUser.Password)
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&newUser.Name)

	newUser.Id = len(users) + 1
	users = append(users, newUser)
	fmt.Println("Register berhasil!")
	isLogin = true
}

func userLogin() bool {
	var username, password string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for _, user := range users {
		if user.Username == username && user.Password == password {
			fmt.Println("Login berhasil!")
			return true
		}
	}

	fmt.Println("Gagal login. Username/password anda salah.")
	return false
}

func editProfile()  {}
func addFriend()    {}
func removeFriend() {}
func friendList()   {}
func searchUser()   {}

func home() {
	fmt.Println("Halaman Home:")
	for {

	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		// For Windows
		cmd = exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func isUsernameExists(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}
