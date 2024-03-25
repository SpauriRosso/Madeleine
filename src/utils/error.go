package utils

import "log"

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func NotPerms() string {
	return "Error: Missing perms"
}

func NotOwner() string {
	return "Error: This command can only be executed by the bot owner!"
}
