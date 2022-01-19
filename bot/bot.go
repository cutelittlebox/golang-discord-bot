package bot

import (
	"fmt" //errors
	"golang-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

var (
	//BotID ID number of bot
	BotID string
	goBot *discordgo.Session
)

func errorHandler(err error) error {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

//Start starts the bot
func Start() {
	//create bot session
	fmt.Println("creating bot session..")
	goBot, err := discordgo.New("Bot " + config.Token)

	if errorHandler(err) != nil {
		return
	}

	fmt.Println("Opening connection")
	err = goBot.Open()
	if errorHandler(err) != nil {
		return
	}

	//make the bot a user
	fmt.Println("creating bot user..")
	u, err := goBot.User("@me")

	if errorHandler(err) != nil {
		return
	}

	BotID = u.ID

	//add a handler to handle messages
	fmt.Println("creating message handler..")
	goBot.AddHandler(messageHandler)

	//if everything has worked so far the bot is ready to start running
	fmt.Println("Bot is now running.")
}

func messageHandler(session *discordgo.Session, mess *discordgo.MessageCreate) {
	//prevent the bot from replying to it's own messages
	if mess.Author.ID == BotID {
		return
	}
	//say "pong" in response to "ping"
	if mess.Content == "ping" {
		_, _ = session.ChannelMessageSend(mess.ChannelID, "pong")
	}
	//say polo in response to marco
	if mess.Content == "marco" {
		_, _ = session.ChannelMessageSend(mess.ChannelID, "polo")
	}
}
