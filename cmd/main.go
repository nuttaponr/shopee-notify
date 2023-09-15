package main

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"shopee-notify/pkg/handler"
	"shopee-notify/pkg/notify"
	"shopee-notify/pkg/shopee"
	"syscall"
	_ "time/tzdata"
)

type Config struct {
	AppCrontab       string `env:"APP_CRONTAB,required"`
	LineChanneltoken string `env:"LINE_CHANNEL_TOKEN,required"`
	ShopeeShopId     string `env:"SHOPEE_SHOP_ID,required"`
	ShopeeItemId     string `env:"SHOPEE_ITEM_ID,required"`
	ShopeeModelId    int64  `env:"SHOPEE_MODEL_ID,required"`
	UserAgent        string `env:"USER_AGENT"`
}

func main() {

	// Loading the environment variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := Config{}       // 👈 new instance of `Config`
	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	caller := shopee.New(cfg.ShopeeShopId, cfg.ShopeeItemId, cfg.ShopeeModelId, shopee.WithUserAgent(cfg.UserAgent))
	notifier := notify.New(cfg.LineChanneltoken)
	handler := handler.New(caller, notifier)

	c := cron.New()
	c.AddFunc(cfg.AppCrontab, handler.DoIt)
	go c.Start()
	log.Println("cron job started")

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s
	c.Stop()
	log.Println("shutting down service")

}
