package main

import (
	"fmt"
	"ng-client/client"
)

func main() {

	//getting some player datas

	client := client.NewNetherGamesClient()

	player, _ := client.GetPlayer("ricccskn")

	fmt.Println("Username: ", player.GetName())
	fmt.Println("Xuid: ", player.GetXuid())
	fmt.Println("Ranks: ", player.GetRanks())
	fmt.Println("Level: ", player.GetLevel())

	//get guild info

	guild, _ := client.GetGuild(player.GetGuild())

	guild.GetLeader()

}
