  # Simple Discord Bot Template

This is a basic template for creating Discord bots in Go.

## Setup 

1. Clone this repository:
```bash
git clone https://github.com/Luiso9/mari-go
```
   
2. Get your Discord bot token from:
```bash
https://discord.com/developers/applications
```

3. Create a file named config.json and add:
```json
{
"token": "your_bot_token_here",
"prefix": "!"  // Change the prefix if you want
}
```
4. Install the discordgo library:
```
go get -u https://github.com/bwmarrin/discordgo
```
6. Run the bot :
``` bash
go run main.go
```

## Adding Commands
Create new files in the commands directory and register them in commands/commands.go. Example 'ping' command is included.

## License
This project is licensed under the MIT License.
