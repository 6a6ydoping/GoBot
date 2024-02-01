package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func listCommands(s *discordgo.Session) {
	commands, err := s.ApplicationCommands("1202473503962566666", "")
	if err != nil {
		log.Println("Error retrieving commands:", err)
		return
	}

	fmt.Println("List of slash commands:")
	for _, command := range commands {
		fmt.Printf("%s - %s\n", command.ID, command.Name)
	}
}
