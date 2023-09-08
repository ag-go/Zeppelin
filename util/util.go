package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func HasArg(arg string) bool {
	for _, s := range os.Args {
		if s == arg {
			return true
		}
	}
	return false
}

func GetArg(name string, def string) string {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, name+"=") {
			if s := strings.TrimPrefix(arg, name+"="); s != "" {
				return s
			}
		}
	}
	return def
}

type Player struct {
	UUID string `json:"id"`
	Name string `json:"name"`
}

func AddDashesToUUID(uuid string) string {
	str := ""
	for i, char := range strings.Split(uuid, "") {
		str += char
		if i == 7 || i == 11 || i == 15 || i == 19 {
			str += "-"
		}
	}
	return str
}

func FetchUsername(username string) (bool, Player) {
	resp, err := http.Get(fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username))
	var player Player
	if err != nil {
		return false, player
	}
	body, _ := io.ReadAll(resp.Body)
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, player
	}
	if data["errorMessage"] != "" {
		return false, player
	}
	player.UUID = AddDashesToUUID(data["id"])
	player.Name = data["name"]
	return true, player
}

func FetchUUID(uuid string) (bool, Player) {
	resp, err := http.Get(fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", uuid))
	var player Player
	if err != nil {
		return false, player
	}
	body, _ := io.ReadAll(resp.Body)
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, player
	}
	if data["errorMessage"] != "" {
		return false, player
	}
	player.UUID = AddDashesToUUID(data["id"])
	player.Name = data["name"]
	return true, player
}

type Placeholders struct {
	PlayerName   string
	Message      string
	PlayerGroup  string
	PlayerPrefix string
	PlayerSuffix string
}

func ParsePlaceholders(str string, placeholders Placeholders) string {
	str = strings.ReplaceAll(str, "%player%", placeholders.PlayerName)
	str = strings.ReplaceAll(str, "%message%", placeholders.Message)
	str = strings.ReplaceAll(str, "%player_prefix%", placeholders.PlayerPrefix)
	str = strings.ReplaceAll(str, "%player_suffix%", placeholders.PlayerSuffix)
	str = strings.ReplaceAll(str, "%player_group%", placeholders.PlayerGroup)
	str = strings.TrimSpace(str)
	return str
}
