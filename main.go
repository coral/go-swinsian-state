package main

import (
	"encoding/json"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type PlayState struct {
	Format string `json:"format"`
	State  string `json:"state"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
}

func main() {
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

func GetState() (PlayState, error) {

	command := `tell application "Swinsian"
	set playstate to player state
	set fileformat to kind of current track
	set trackname to name of current track
	set trackartist to artist of current track
	set trackalbum to album of current track
	set info to "{\"format\": \"" & fileformat & "\",\"state\": \"" & playstate & "\",\"song\": \"" & trackname & "\",\"artist\": \"" & trackartist & "\",\"album\": \"" & trackalbum & "\"}"
end tell`

	cmd := exec.Command("osascript", "-e", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return PlayState{}, err
	}

	var state PlayState

	err = json.Unmarshal(output, &state)

	return state, err

}
