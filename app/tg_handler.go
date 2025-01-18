package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/l4m3r/serverless-guard/app/apigw"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type app struct {
	// logger *slog.Logger
	Bot *tgbotapi.BotAPI
}

func NewApp() (*app, error) {
	botAPIToken := os.Getenv("TG_TOKEN")
	if botAPIToken == "" {
		return nil, fmt.Errorf("No token")
	}
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		return nil, err
	}
	return &app{
		Bot: bot,
	}, nil
}

func TgHandler(ctx context.Context, req *apigw.APIGatewayRequest) (*apigw.APIGatewayResponse, error) {
	app, err := NewApp()
	if err != nil {
		return nil, err
	}
	var bodyBytes []byte
	if req.IsBase64Encoded {
		bodyBytes, err = base64.StdEncoding.DecodeString(string(req.Body))
		if err != nil {
			return nil, err
		}
	} else {
		bodyBytes = []byte(req.Body)
	}
	var body = apigw.Request{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, err
	}
	message := body.Message
	reply := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err = app.Bot.Send(reply)
	if err != nil {
		return nil, err
	}
	return &apigw.APIGatewayResponse{
		StatusCode: 200,
	}, nil
}
