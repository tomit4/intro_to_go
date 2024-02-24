package main

import "fmt"

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println("something went wrong when getting the user :=>", err)
		return
	}
	profile, err := getProfile(user.ID)
	if err != nil {
		fmt.Println("something went wrong when getting profile :=>", err)
		return
	}

}

func getUser() (User, error) {
	// do some get user logic here
}
