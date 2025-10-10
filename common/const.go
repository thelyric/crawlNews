package common

import "log"

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}

const CurrentUser = "currentuser"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
