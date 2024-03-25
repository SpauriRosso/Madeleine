package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func BotRun() {
	BotToken := GetEnv("APP_USER_TOKEN")
	discord, err := discordgo.New("Bot " + BotToken)
	CheckError(err)

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		msg := strings.ToLower(m.Content)
		if strings.Contains(msg, "aide") {
			s.ChannelMessageSend(m.ChannelID, "import (\"Cherche toi même\")")
		}
	})
	discord.Identify.Intents = discordgo.IntentsAll
	discord.Open()
	defer discord.Close()
	fmt.Println("connected!")
	PreventBotOffline()
}

func GetEnv(key string) string {
	CheckError(godotenv.Load("./.env"))
	return os.Getenv(key)
}

// PreventBotOffline : Prevent bot from going offline one main.go is executed
func PreventBotOffline() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
