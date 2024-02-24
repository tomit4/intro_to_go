package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// ?
func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("======================================")
}

// main function not seen in vid
