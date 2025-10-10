package cmds

import (
	"github.com/bwmarrin/discordgo"
)

func PongCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Gọi cái thằng bố mày")
}
