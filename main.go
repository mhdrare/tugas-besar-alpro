package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
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
var friends []Friends

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

func home() {
	for {
		fmt.Println("1. Lihat Postingan teman")
		fmt.Println("2. Buat Postingan")
		fmt.Println("3. Daftar teman")
		fmt.Println("4. Cari pengguna")
		fmt.Println("5. Edit Profil")
		fmt.Println("6. Logout")

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewPosts()
		case 2:
			createPost()
		case 3:
			friendList()
		case 4:
			searchUser()
		case 5:
			editProfile()
		case 6:
			isLogin = false
			return
		default:
			fmt.Println("Pilihan tidak valid. Silahkan masukkan opsi yang valid.")
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

func editProfile() {
	fmt.Println("Edit Profil:")
	fmt.Println("1. Ubah Password")
	fmt.Println("2. Ubah Nama")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		changePassword()
	case 2:
		changeName()
	default:
		fmt.Println("Pilihan tidak valid. Silahkan masukkan opsi yang valid.")
	}
}

func changePassword() {
	var newPassword string
	fmt.Print("Masukkan password baru Anda: ")
	fmt.Scan(&newPassword)

	users[0].Password = newPassword

	fmt.Println("Password berhasil diubah.")
}

func changeName() {
	var newName string
	fmt.Print("Masukkan nama baru Anda: ")
	fmt.Scan(&newName)

	users[0].Name = newName

	fmt.Println("Nama berhasil diubah.")
}

func addFriend() {
	var friendUsername string
	fmt.Print("Masukkan username yang dicari: ")
	fmt.Scan(&friendUsername)

	var friend User
	found := false
	for _, user := range users {
		if user.Username == friendUsername {
			friend = user
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	for _, existingFriendship := range friends {
		if existingFriendship.UserId == users[0].Id && existingFriendship.FriendId == friend.Id {
			fmt.Println("Kamu telah berteman dengan", friend.Name)
			return
		}
	}

	friends = append(friends, Friends{
		UserId:   users[0].Id,
		FriendId: friend.Id,
	})

	fmt.Println("Berhasil menambahkan teman.")
}

func removeFriend() {
	var friendUsername string
	fmt.Print("Masukkan username pengguna yang akan dihapus: ")
	fmt.Scan(&friendUsername)

	var friend User
	found := false
	for _, user := range users {
		if user.Username == friendUsername {
			friend = user
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pengguna tidak ditemukan.")
		return
	}

	foundFriendshipIndex := -1
	for i, existingFriendship := range friends {
		if existingFriendship.UserId == users[0].Id && existingFriendship.FriendId == friend.Id {
			foundFriendshipIndex = i
			break
		}
	}

	if foundFriendshipIndex == -1 {
		fmt.Println("Kamu tidak berteman dengan", friend.Name)
		return
	}

	friends = append(friends[:foundFriendshipIndex], friends[foundFriendshipIndex+1:]...)

	fmt.Println("Berhasil menghapus teman.")
}

func friendList() {
	for {
		fmt.Println("1. Lihat Teman")
		fmt.Println("2. Tambah Friend")
		fmt.Println("3. Hapus Friend")
		fmt.Println("4. Kembali ke Home")

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewFriends()
		case 2:
			addFriend()
		case 3:
			removeFriend()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silahkan masukkan opsi yang valid.")
		}
	}
}

func viewFriends() {
	fmt.Println("Daftar teman:")
	fmt.Println("ID\tNama\tUsername")

	sortedFriends := sortFriendsByName(friends)

	for _, friend := range sortedFriends {
		user := getUserByID(friend.FriendId)
		fmt.Printf("%d\t%s\t%s\n", user.Id, user.Name, user.Username)
	}
}

func searchUser() {
	var searchQuery string
	fmt.Print("Masukkan username yang dicari: ")
	fmt.Scan(&searchQuery)

	fmt.Println("Hasil pencarian:")
	fmt.Println("ID\tNama\tUsername\tTeman?")

	for _, user := range users {
		if user.Username == searchQuery {
			isFriend := isUserFriend(user.Id)
			fmt.Printf("%d\t%s\t%s\t%t\n", user.Id, user.Name, user.Username, isFriend)
		}
	}

	if !isUserExist(searchQuery) {
		fmt.Println("Pengguna tidak ditemukan.")
	}
}

func createPost() {
	var postText string
	fmt.Print("Masukkan postingan: ")
	fmt.Scan(&postText)

	post := Posts{
		Id:     len(posts) + 1,
		UserId: users[0].Id,
		Post:   postText,
	}

	posts = append(posts, post)

	fmt.Println("Berhasil membuat postingan.")
}

func viewPosts() {
	fmt.Println("Postingan:")
	fmt.Println("ID\tUser\tPostingan")

	// Display posts
	for _, post := range posts {
		user := getUserByID(post.UserId)
		fmt.Printf("%d\t%s\t%s\n", post.Id, user.Name, post.Post)
	}
}

func sortFriendsByName(friends []Friends) []Friends {
	var friendData []struct {
		Friend Friends
		User   User
	}

	for _, friend := range friends {
		user := getUserByID(friend.FriendId)
		friendData = append(friendData, struct {
			Friend Friends
			User   User
		}{Friend: friend, User: user})
	}

	sort.Slice(friendData, func(i, j int) bool {
		return friendData[i].User.Name < friendData[j].User.Name
	})

	var sortedFriends []Friends
	for _, data := range friendData {
		sortedFriends = append(sortedFriends, data.Friend)
	}

	return sortedFriends
}

func getUserByID(userID int) User {
	for _, user := range users {
		if user.Id == userID {
			return user
		}
	}
	return User{}
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

func isUserExist(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func isUserFriend(userID int) bool {
	for _, friend := range friends {
		if friend.FriendId == userID && friend.UserId == users[0].Id {
			return true
		}
	}
	return false
}
