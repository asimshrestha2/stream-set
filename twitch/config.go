package twitch

const (
	ClientID        string = "1jcuu1fyzg8nabsmoplijb826zoyte0"
	RedirectURI     string = "http://localhost:8000/twitch/token/"
	RequestTokenURL string = "" +
		"https://id.twitch.tv/oauth2/authorize" +
		"?response_type=token" +
		"&client_id=" + ClientID +
		"&redirect_uri=" + RedirectURI +
		"&force_verify=true" +
		"&scope=channel_editor+channel_read"

	TwitchAPIURL string = "https://api.twitch.tv/kraken"
)

var (
	Token       = ""
	GameList    TopGamesResponse
	GameDB      []DBGame
	UserChannel Channel
)

type (
	DBGame struct {
		TwitchName       string   `json:"twitchName"`
		FileName         string   `json:"fileName"`
		AlternativeNames []string `json:"alternativeName"`
	}
)
