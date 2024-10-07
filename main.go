package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"os/signal"
)

func GetTOML(x string) string {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
		return ""
	}
	result := config.Get(x).(string)
	return result
}

var (
	s       *discordgo.Session
	GuildID string
	err     error
)

func init() {
	botToken := util.GetTOML("Bot.token")
	GuildID = util.GetTOML("Bot.guild_id")
	s, err = discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error creating Discord session,", err.Error())
		return
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		//UTILITY COMMANDS
		{
			Name:        "rollcall",
			Description: "Start up match request and let player join off a button",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "tourney-id",
					Description: "Input ID of the tournament",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		//CREATE HANDLER
		"create": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case "tournament":
				if err := cmd.CreateTournamentCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "participant":
				if err := cmd.AddParticipantsCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			}
		},
		//REMOVE HANDLER
		"remove": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case "tournament":
				if err := cmd.RemoveTournament(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "participant":
				if err := cmd.RemoveParticipantCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			}

		},
		//UPDATE HANDLER
		"update": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case "tournament":
				if err := cmd.UpdateTournament(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "tournamentstate":
				if err := cmd.UpdateTournamentState(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "participant":
				if err := cmd.UpdateParticipantCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "match":
				if err := cmd.UpdateMatchCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			}
		},
		"show": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			switch options[0].Name {
			case "tournament":
				if err := cmd.ShowTournamentCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "match":
				if err := cmd.ShowMatchCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "participant":
				if err := cmd.ShowParticipantCMD(s, i); err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error occurred when using command. Please try again later.",
						},
					})
					fmt.Println(err)
				}
			case "all":
				options = options[0].Options
				switch options[0].Name {
				case "tournament":
					if err := cmd.ShowAllTournamentsCMD(s, i); err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: "Error occurred when using command. Please try again later.",
							},
						})
						fmt.Println(err)
					}
				case "match":
					if err := cmd.ShowAllMatchesCMD(s, i); err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: "Error occurred when using command. Please try again later.",
							},
						})
						fmt.Println(err)
					}
				case "participant":
					if err := cmd.ShowAllParticipantsCMD(s, i); err != nil {
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: "Error occurred when using command. Please try again later.",
							},
						})
						fmt.Println(err)
					}
				}
			}
		},
		"rollcall": cmd.RollCallCMD(),
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"rc_join":  cmd.RCJoinComponent(),
		"rc_close": cmd.RCCloseComponent(),
	}
)

func init() {
}

func main() {
	var RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("BracketBot 1.0 \n Logged in as: %v#%v  GuildID: %v", s.State.User.Username, s.State.User.Discriminator, GuildID)
	})
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
	// add a event handler
	if err := s.Open(); err != nil {
		fmt.Println("Error opening connection,", err.Error())
		return
	}
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	fmt.Println("Adding commands...")
	for i, v := range commands {
		fmt.Println("Added " + v.Name)
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	// keep cmd running untill there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	if *RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
	log.Println("Gracefully shutting down.")
}
