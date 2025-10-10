package utils

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func JoinUserVoice(session *discordgo.Session, message *discordgo.MessageCreate, guildID, ChannelID string) *discordgo.VoiceConnection {
	voiceChannel, err := session.ChannelVoiceJoin(guildID, ChannelID, false, true)
	if err != nil {
		fmt.Println("ChannelVoiceJoin error: %w", err)
		return nil
	}
	return voiceChannel
}

func FindUserVoiceState(session *discordgo.Session, guildId, userID string) *discordgo.VoiceState {
	guild, err := session.State.Guild(guildId)
	if err != nil {
		return nil
	}

	for _, state := range guild.VoiceStates {
		if state.UserID == userID {
			return state
		}

	}
	return nil
}

func CheckUrl(text string) (isUrl bool) {
	if text == "" || len(text) < 3 {
		return false
	}
	u, err := url.ParseRequestURI(text)
	if err != nil {
		return false
	}

	return u.Host != ""
}

func CheckYoutubeLink(text string) (isYoutube bool) {
	host := strings.ToLower(text)
	return host == "www.youtube.com" ||
		host == "youtube.com" ||
		host == "m.youtube.com" ||
		host == "youtu.be"
}

func IsHTTP(s string) bool {
	if !CheckUrl(s) {
		return false
	}
	u, _ := url.Parse(s)

	if isYoutube := CheckYoutubeLink(u.Host); !isYoutube {
		return false
	}

	scheme := strings.ToLower(u.Scheme)
	return scheme == "http" || scheme == "https"
}
