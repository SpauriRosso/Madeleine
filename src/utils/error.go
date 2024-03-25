package utils

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NotPerms() string {
	return "Error: Missing perms"
}

func NotOwner() string {
	return "Error: This command can only be executed by the bot owner!"
}
