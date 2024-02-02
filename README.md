---

# Discord Bot

This Discord bot is designed to provide various functionalities, including weather information, reminders, language translation, and trivia games.

## Features

### 1. Weather Command

Use `!weather [city]` to get the current weather for a specific city.

### 2. Reminder Command

Set reminders with the `!reminder` command. Example: `!reminder date="02.02.2024" time="14:36" gmt="+6" text="Meeting with John"`

### 3. Translation Command

Translate text from one language to another using the `!translate` command. Example: `!translate [en] [ru] Hello, how are you?`

### 4. Trivia Game

Start a trivia game with the `!trivia` command.

### 5. Help Command

Use `!help` to get information about available commands and their usage.

## Usage

1. Start the chat with Discord bot.
2. Use the available commands to interact with the bot.

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/6adydoping/GoBot.git
   cd GoBot
   ```

2. Install dependencies:

   ```bash
   go get -u ./...
   ```

3. Configure the bot:

   - Set up environment variables or configuration files with your Discord bot token and other required credentials.

4. Run the bot:

   ```bash
   go run main.go
   ```

## Contributing

Feel free to contribute to the development of this Discord bot. If you have suggestions or find any issues, please open an [issue](https://github.com/6a6ydoping/GoBot/issues) or submit a [pull request](https://github.com/6a6ydoping/GoBot/pulls).
