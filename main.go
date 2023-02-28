package main

import (
	"encoding/json"
	"fmt"
)

var DA = `
{
	"success": true,
	"data": [
			{
			"contract": "atomicassets",
			"collection_name": "chapeauverse",
			"name": "chapeauverse",
			"img": "QmXYDCuARTkzjx5FRobHjLWneLvUwpoAV9v9tYyPSVT2f4",
			"author": "ujro4.c.wam",
			"allow_notify": true,
			"authorized_accounts": [
				"ujro4.c.wam"
			],
			"data": {
				"img": "QmUqpRHssBFjmg1rCHK3R78DcGF8tjFK43tKKFyomzyNF6",
			"url": "https://www.spacepower.io/",
			"name": "Space Power Game",
			"images": "{\"banner_1920x500\":\"QmPG38LcSEJCXpdiARZCArxFuUiVZBcK16UVSdZA2jczkf\",\"logo_512x512\":\"QmZKb7GPpDP9YhKL2Vte8GGsd1GQsR9xHommbptda4gpFK\"}",
			"socials": "{\"twitter\":\"https://twitter.com/Spacepowerio\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"https://discord.gg/YQUVVE2yPH\",\"youtube\":\"@spacepowerio\",\"telegram\":\"https://t.me/spacepowergame\"}",
			"description": "Space Power is a NFT strategy game based on the WAX blockchain. The game has Play to Earn features, where players will be rewarded with Tokens and valuable NFTs.\n- Discord: https://discord.gg/YQUVVE2yPH\n- Twitter: https://twitter.com/Spacepowerio\n- YouTube: https://www.youtube.com/@spacepowerio\n- Telegram: https://t.me/spacepowergame\n- Instagram: https://www.instagram.com/spacepower.io/",
			"creator_info": "{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}"
			},
			"created_at_time": "1675434943500",
			"created_at_block": "227934278"
			}
	]
}
`

type MyJsonName struct {
	Data []struct {
		AllowNotify        bool     `json:"allow_notify"`
		Author             string   `json:"author"`
		AuthorizedAccounts []string `json:"authorized_accounts"`
		CollectionName     string   `json:"collection_name"`
		Contract           string   `json:"contract"`
		CreatedAtBlock     string   `json:"created_at_block"`
		CreatedAtTime      string   `json:"created_at_time"`
		Data               struct {
			CreatorInfo string `json:"creator_info"`
			Description string `json:"description"`
			Images      string `json:"images"`
			Img         string `json:"img"`
			Name        string `json:"name"`
			Socials     string `json:"socials"`
			URL         string `json:"url"`
		} `json:"data"`
		Img  string `json:"img"`
		Name string `json:"name"`
	} `json:"data"`
	Success bool `json:"success"`
}

type So struct {
	Discord  string `json:"discord"`
	
	
}

func main() {
	
	//var Otv MyJsonName
	var fin map[string]interface{}
	jsonerr := json.Unmarshal([]byte(DA), &fin)
	if jsonerr != nil {
		panic(jsonerr)
	}
	
	fmt.Println(fin)
	
	
	

}
