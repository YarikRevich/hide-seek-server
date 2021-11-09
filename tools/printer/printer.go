package printer

import (
	"fmt"
	"strings"

	"github.com/mbndr/figlet4go"
)

//Prints welcome message when server is run
func PrintWelcomeMessage() {
	renderer := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}

	for _, v := range strings.Split("HideSeek\nServer!", "\n") {
		text, err := renderer.RenderOpts(v, options)
		if err != nil {
			panic(err)
		}

		fmt.Print(text)

	}
}
