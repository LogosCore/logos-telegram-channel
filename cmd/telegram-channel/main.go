package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/logoscore/logos-telegram-channel/internal/config"
	"github.com/logoscore/logos-telegram-channel/internal/transport/telegram"
)

func main() {
	cfg, err := config.Load("configs/channel.example.yaml")
	if err != nil {
		log.Fatal(err)
	}
	profiles, err := config.LoadProfiles(cfg.ProfilesFile)
	if err != nil {
		log.Fatal(err)
	}
	ch := telegram.New(cfg.ChannelID, cfg.BotToken, cfg.LogosSyncBaseURL, cfg.PollTimeout, profiles)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("starting telegram channel: channel_id=%s profiles=%d", cfg.ChannelID, len(profiles))
	if err := ch.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
