package gobot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tb "gopkg.in/telebot.v3"
)

var (
	appVersion = ""
	TeleToken  = os.Getenv("TELE_TOKEN")
)

var rootCmd = &cobra.Command{
	Use:     "bot",
	Aliases: []string{"start"},
	Short:   "Telegram bot",
	Long:    `This is a sample Telegram bot using Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bot started...%s", appVersion)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  TeleToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
		return
	}

	b.Handle("/start", func(m tb.Context) error {
		return m.Send("Hi, I`m first GO bot by Dexjke!")
	})

	b.Handle(tb.OnText, func(m tb.Context) error {
		text := m.Text()

		switch text {
		case "hello":
			err = m.Send("Hello, I`m echo-bot")
		default:
			err = m.Send(fmt.Sprintf("U say: %s", text))
		}

		return err
	})
	b.Start()
}
