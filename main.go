package main

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type PlayState struct {
	Format string `json:"format"`
	State  string `json:"state"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
}

type PlayHolder struct {
	Spotify  PlayState `json:"spotify"`
	Swinsian PlayState `json:"swinsian"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		state, err := GetState()
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Nope",
			})
		}
		c.JSON(200, state)
	})
	r.Run()
}

func GetState() (PlayHolder, error) {

	// 	command := `tell application "Swinsian"
	// 	set playstate to player state
	// 	set fileformat to kind of current track
	// 	set trackname to name of current track
	// 	set trackartist to artist of current track
	// 	set trackalbum to album of current track
	// 	set info to "{\"format\": \"" & fileformat & "\",\"state\": \"" & playstate & "\",\"song\": \"" & trackname & "\",\"artist\": \"" & trackartist & "\",\"album\": \"" & trackalbum & "\"}"
	// end tell`

	cmd := exec.Command("osascript", "as.scpt")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return PlayHolder{}, err
	}

	var state PlayHolder

	err = json.Unmarshal(output, &state)

	state.Swinsian.Format = strings.ToUpper(state.Swinsian.Format)

	return state, err

}
