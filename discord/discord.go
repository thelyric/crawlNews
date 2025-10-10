package discord

import (
	"fmt"
	configBot "my-app/discord/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func InitDiscord() {
	token := configBot.LoadToken()
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Bật intent để đọc nội dung tin nhắn
	dg.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildVoiceStates |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent
	dg.AddHandler(MessageCreate)

	dg.Open()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	dg.Close()
}
