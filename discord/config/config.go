package configBot

import "os"

func LoadToken() string {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		panic("DISCORD_TOKEN is not set")
	}
	return "Bot " + token

}
