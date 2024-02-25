package main

import (
	"fmt"
)

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	// NOTE: We cannot just delete here without the `defer` keyword,
	// because this will delete the user prior to these checks taking place
	// we want go through these conditionals, and AFTERWARDS, ALWAYS delete the user by name
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

type user struct {
	name   string
	number int
	admin  bool
}
