package cmds

import (
	"fmt"
	"io"
	"log/slog"
	"my-app/discord/utils"
	"os/exec"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

func Play(vc *discordgo.VoiceConnection, url string) {
	vc.Speaking(true)
	defer vc.Speaking(false)
	go func(link string) {

		// Lệnh yt-dlp để lấy audio từ YouTube

		cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", "-", link)

		// Pipe stdout của yt-dlp vào DCA
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error creating stdout pipe:", err)
			return
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting yt-dlp:", err)
			return
		}

		// Đọc tất cả dữ liệu
		bytes, err := io.ReadAll(stdout)
		if err != nil {
			panic(err)
		}

		// Chuyển sang string
		str := string(bytes)

		opts := dca.StdEncodeOptions
		opts.RawOutput = true
		opts.Bitrate = 128

		encodeSession, err := dca.EncodeFile(str, opts)

		if err != nil {
			slog.Error("Failed to create an encoding session", "error", err)
		}

		done := make(chan error)
		dca.NewStream(encodeSession, vc, done)

		for err := range done {
			if err != nil && err != io.EOF {
				slog.Error("An error occurred", "error", err)
			}

			// Clean up in case something happened and ffmpeg is still running
			encodeSession.Cleanup()
		}
	}(url)
}

func PlayCommand(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	voiceState := utils.FindUserVoiceState(session, message.GuildID, message.Author.ID)
	if voiceState == nil {
		session.ChannelMessageSend(message.ChannelID, "⚠️ Vui lòng vào voice channel trước.")
		return
	}

	vc := utils.JoinUserVoice(session, message, message.GuildID, voiceState.ChannelID)

	if vc == nil {
		session.ChannelMessageSend(message.ChannelID, "❌ Có lỗi xảy ra khi join voice")
		return
	}

	if len(args) == 0 {
		session.ChannelMessageSend(message.ChannelID, "⚠️ Mày có chắc chắn là muốn bố mày hát khi không có link à ?")
		return
	}

	urlString := args[0]
	isValidLink := utils.IsHTTP(urlString)

	if !isValidLink {
		session.ChannelMessageSend(message.ChannelID, "⚠️ Mày gửi cái link ĐẦU BU*I gì đấy ?")
		return
	}

	// Thực thi phát nhạc không chặn handler
	go func() {
		Play(vc, urlString)
		time.Sleep(time.Second * 30)
		vc.Disconnect()
	}()
}
