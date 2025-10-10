package discord

import (
	cmds "my-app/discord/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const BotPrefix = "dmhieu!"

var Commands = map[string]func(*discordgo.Session, *discordgo.MessageCreate, []string){
	"ping": cmds.PingCommand,
	"pong": cmds.PongCommand,
	"play": cmds.PlayCommand,
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Only handle messages that start with the prefix
	if len(m.Content) < len(BotPrefix) || !strings.HasPrefix(m.Content, BotPrefix) {
		return
	}

	cmd := strings.TrimSpace(m.Content[len(BotPrefix):])
	args := strings.Fields(cmd)
	command := strings.ToLower(args[0])
	args = args[1:]

	if handler, ok := Commands[command]; ok {
		handler(s, m, args)
	}
}
