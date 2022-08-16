package main

import (
	"fmt"
	"os"
	"context"
	"log"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){

	for event:= range analyticsChannel{
		fmt.Println("Command events:")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}

}



func main(){

	os.Setenv("SLACK_BOT_TOCKEN", "xoxb-3938092225443-3931627490614-swCZva0NfZwfvo2WJwgama8H")
	os.Setenv("SLACk_APP_TOCKEN", "xapp-1-A03TDH7GFEJ-3938094080739-5b4f942459a67bd25e1475f223587784805c0dc430955e1219fe8e5a58aaeeba")



	bot:= slacker.NewClient(os.Getenv("SLACK_BOT_TOCKEN"), os.Getenv("SLACK_APP_TOCKEN"))
	
	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Reply from Akshat's bot")
		},
	})

	bot.Command("Hi", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello, How you doin")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err:= bot.Listen(ctx)
	if err!= nil{
		log.Fatal(err)
	}

}