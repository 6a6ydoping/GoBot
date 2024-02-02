package bot

import "github.com/bwmarrin/discordgo"

func (mg Manager) help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, `
	**Available Commands:**
1. `+"`!weather [city]`"+` - Get the current weather for a specific city.
2. `+"`!reminder date=\"date\" time=\"time\" gmt=\"gmt\" text=\"message\"`"+` - Set a reminder for a specific date and time.
3. `+"`!translate [fromLang] [toLang] text`"+` - Translate text from one language to another.
4. `+"`!trivia`"+` - Start a trivia game.

**Example Usage:**
- `+"`!weather Almaty`"+`
- `+"`!reminder date=\"02.02.2024\" time=\"14:36\" gmt=\"+6\" text=\"Meeting with John\"`"+`
- `+"`!translate [en] [ru] Hello, how are you?`"+`
- `+"`!trivia`"+`
it 
Feel free to use these commands to interact with the bot!
`)
}
