package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func BotRun() {
	BotToken := GetEnv("APP_USER_TOKEN")
	var Discord, err = discordgo.New("Bot " + BotToken)
	CheckError(err)
	Discord.Identify.Intents = discordgo.IntentsAll
	Discord.Open()
	defer Discord.Close()
	fmt.Println("connected!")
	PreventBotOffline()
}

func GetEnv(key string) string {
	CheckError(godotenv.Load("./.env"))
	return os.Getenv(key)
}

func GetData(key string) string {
	// define config file path
	viper.AddConfigPath("./configs")
	// define config file(s)
	viper.SetConfigName("config")
	// define config file to be read
	viper.SetConfigType("json")
	CheckError(viper.ReadInConfig())

	return viper.GetString(key)
}

func IsAdmin(uid int) bool {
	id := strconv.Itoa(uid)
	admins := viper.GetStringSlice("Admins")

	for _, v := range admins {
		if id == v {
			return true
		}
	}
	return false
}

func IsAuthorizedUserRole(rid int) bool {
	id := strconv.Itoa(rid)
	roles := viper.GetStringSlice("Chatbot.AuthorizedUser")

	for _, v := range roles {
		if id == v {
			return true
		}
	}

	return false
}

// PreventBotOffline : Prevent bot from going offline one main.go is executed
func PreventBotOffline() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
