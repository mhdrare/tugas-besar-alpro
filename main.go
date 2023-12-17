package main

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

func main() {

}

func userRegister() {}
func userLogin()    {}
func editProfile()  {}
func addFriend()    {}
func removeFriend() {}
func friendList()   {}
func searchUser()   {}
