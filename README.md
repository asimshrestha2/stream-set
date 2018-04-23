# Stream Set
A tool that helps Twitch Streamers run their streams.

## Currently Supported
* Auto Updating Games
* Reseting to Default Game

### Auto Updating Games
When you are logged in with Twitch, the App detects the Active window. If the active window matches the a game name from Twitch's top game list, then it updates on the logged in Twitch Channel. 

### Default settings.ini

```ini
[list]
# List of Apps Titles to Avoid
ignore = Firefox, Google Chrome, Discord, Steam, Blizzard Battle.net, Epic Games Launcher, Stream Set

[twitch]
# Game that the App sets when not in you are not playing any game
defaultGame = IRL
# The total time the app waits for to go back to the Default Game (in Sec)
waitToReset = 300
```