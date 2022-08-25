package client

import (
	"encoding/json"
	"net/http"
	"ng-client/response"
)

type NetherGamesClient struct {
}

const (
	player_url = "https://api.ngmc.co/v1/players/"
	guild_url  = "https://api.ngmc.co/v1/players/"
)

func NewNetherGamesClient() *NetherGamesClient {
	return &NetherGamesClient{}
}

func (ng *NetherGamesClient) GetPlayer(username string) (*response.PlayerResponse, error) {

	res, err := http.Get(player_url + username)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var player response.PlayerResponse

	err = json.NewDecoder(res.Body).Decode(&player)

	return &player, err
}

func (ng NetherGamesClient) GetGuild(guildName string) (*response.GuildResponse, error) {

	res, err := http.Get(guild_url + guildName)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var guild response.GuildResponse

	err = json.NewDecoder(res.Body).Decode(&guild)

	return &guild, err
}
