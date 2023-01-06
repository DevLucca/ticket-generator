package main

import (
	_ "embed"
	"os"
	"fmt"
	"io/ioutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/akamensky/argparse"
	"github.com/tdewolff/canvas"
)

//go:embed ticket.svg
var svgTemplate string

var port *int

type group struct{
	name string
	color string
}

var groups = map[string]group{
	"1": group{
		name: `
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(477.289512, 498.047779)">
		<g>
			<path class="st7" d="M161.9,152.1h-2.1V150h4.5V166h-2.4V152.1z"/>
		</g>
	</g>
</g>
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(486.805114, 498.047779)">
		<g>
			<path class="st7" d="M163.3,158.6h6.7v1.5h-6.7V158.6z"/>
		</g>
	</g>
</g>
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(497.782046, 498.047779)">
		<g>
			<path class="st7" d="M172.9,166l-6.7-16.1h2.6l5.4,13l5.4-13h2.6l-6.7,16.1H172.9z"/>
		</g>
	</g>
</g>
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(510.220297, 498.047779)">
		<g>
			<path class="st7" d="M181.2,150v2.1h-7.5v4.4h6.6v2.2h-6.6v5.2h7.5v2.2h-9.8V150H181.2z"/>
		</g>
	</g>
</g>
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(519.548982, 498.047779)">
		<g>
			<path class="st7" d="M187.7,166h-3l-5.1-6.9h-2.8v6.9h-2.4V150h6.5c0.6,0,1.2,0.1,1.8,0.4c0.6,0.2,1.1,0.6,1.5,1     c0.4,0.4,0.7,0.9,1,1.4c0.2,0.6,0.4,1.1,0.4,1.8c0,0.5-0.1,1-0.2,1.5c-0.2,0.5-0.4,0.9-0.7,1.3s-0.6,0.7-1.1,1     c-0.4,0.3-0.8,0.5-1.3,0.6L187.7,166z M181,156.8c0.3,0,0.6-0.1,0.9-0.2c0.3-0.1,0.5-0.3,0.7-0.5c0.2-0.2,0.4-0.5,0.5-0.8     c0.1-0.3,0.2-0.6,0.2-0.9c0-0.3-0.1-0.6-0.2-0.9c-0.1-0.3-0.3-0.5-0.5-0.8c-0.2-0.2-0.4-0.4-0.7-0.5c-0.3-0.1-0.6-0.2-0.9-0.2     h-4.1v4.8L181,156.8z"/>
		</g>
	</g>
</g>
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(530.695836, 498.047779)">
		<g>
			<path class="st7" d="M178.3,166V150h5.7c1.2,0,2.2,0.2,3.1,0.6c0.9,0.4,1.7,1,2.2,1.7c0.6,0.7,1,1.6,1.4,2.6c0.3,1,0.5,2,0.5,3.1     c0,1.1-0.2,2.2-0.5,3.1c-0.3,1-0.8,1.8-1.3,2.5c-0.6,0.7-1.2,1.3-2,1.7c-0.8,0.4-1.6,0.6-2.5,0.6H178.3z M183.9,163.8     c0.9,0,1.6-0.2,2.2-0.5c0.6-0.3,1.1-0.7,1.5-1.3c0.4-0.5,0.7-1.1,0.9-1.9c0.2-0.7,0.3-1.5,0.3-2.2c0-0.8-0.1-1.5-0.3-2.2     c-0.2-0.7-0.5-1.4-0.9-1.9c-0.4-0.5-0.9-1-1.5-1.3c-0.6-0.3-1.3-0.5-2.2-0.5h-3.2v11.7H183.9z"/>
		</g>
	</g>
</g>
<g xmlns="http://www.w3.org/2000/svg">
	<g transform="translate(542.216511, 498.047779)">
		<g>
			<path class="st7" d="M191.8,150v2.1h-7.5v4.4h6.6v2.2h-6.6v5.2h7.5v2.2H182V150H191.8z"/>
		</g>
	</g>
</g>
		`,
		color: "36603E",
	},
	"2": group{
		name: `
		<g xmlns="http://www.w3.org/2000/svg">
    <g transform="translate(484.566856, 498.047779)">
      <g>
        <path class="st7" d="M162.5,166v-1.6l4.6-5.3c0.7-0.8,1.2-1.5,1.6-1.9c0.4-0.5,0.7-0.8,0.9-1.1c0.2-0.3,0.3-0.5,0.4-0.8 c0.1-0.2,0.1-0.5,0.1-0.8c0-0.8-0.2-1.5-0.8-1.9c-0.5-0.4-1.1-0.6-1.8-0.6c-0.4,0-0.7,0.1-1,0.2c-0.3,0.1-0.6,0.3-0.8,0.6 c-0.2,0.2-0.4,0.5-0.6,0.8c-0.1,0.3-0.2,0.6-0.2,0.9h-2.3c0-0.6,0.1-1.2,0.4-1.8c0.3-0.6,0.6-1.1,1.1-1.5c0.4-0.4,1-0.8,1.5-1 c0.6-0.3,1.2-0.4,1.9-0.4c0.7,0,1.3,0.1,1.9,0.4c0.6,0.2,1.1,0.6,1.5,1c0.4,0.4,0.8,0.9,1,1.5c0.3,0.6,0.4,1.2,0.4,1.8 c0,0.5-0.1,1-0.2,1.4c-0.1,0.5-0.3,0.9-0.6,1.2l-5.7,6.6h6.8v2.2H162.5z"/>
      </g>
    </g>
  </g>
  <g xmlns="http://www.w3.org/2000/svg">
    <g transform="translate(498.364482, 498.047779)">
      <g>
        <path class="st7" d="M167.2,158.6h6.7v1.5h-6.7V158.6z"/>
      </g>
    </g>
  </g>
  <g xmlns="http://www.w3.org/2000/svg">
    <g transform="translate(509.341414, 498.047779)">
      <g>
        <path class="st7" d="M172.6,166H170l6.7-16.1h2.6L186,166h-2.6l-2.1-5h-6.6L172.6,166z M175.6,158.8h4.8L178,153L175.6,158.8z"/>
      </g>
    </g>
  </g>
  <g xmlns="http://www.w3.org/2000/svg">
    <g transform="translate(521.779664, 498.047779)">
      <g>
        <path class="st7" d="M174.7,166v-1.9l8.5-12h-7.7V150H186v1.9l-8.3,11.9h8.8v2.2H174.7z"/>
      </g>
    </g>
  </g>
  <g xmlns="http://www.w3.org/2000/svg">
    <g transform="translate(532.059914, 498.047779)">
      <g>
        <path class="st7" d="M190.4,160.3c0,0.8-0.2,1.6-0.5,2.3c-0.3,0.7-0.7,1.3-1.2,1.9c-0.5,0.5-1.1,1-1.9,1.3 c-0.7,0.3-1.5,0.5-2.3,0.5c-0.8,0-1.6-0.2-2.3-0.5c-0.7-0.3-1.3-0.7-1.9-1.3c-0.5-0.5-1-1.2-1.3-1.9c-0.3-0.7-0.5-1.5-0.5-2.3 V150h2.4v10.4c0,0.5,0.1,1,0.3,1.4c0.2,0.4,0.4,0.8,0.8,1.2c0.3,0.3,0.7,0.6,1.1,0.8c0.4,0.2,0.9,0.3,1.4,0.3 c0.5,0,0.9-0.1,1.4-0.3c0.4-0.2,0.8-0.5,1.1-0.8c0.3-0.3,0.6-0.7,0.8-1.2c0.2-0.4,0.3-0.9,0.3-1.4V150h2.4V160.3z"/>
      </g>
    </g>
  </g>
  <g xmlns="http://www.w3.org/2000/svg">
    <g transform="translate(542.832936, 498.047779)">
      <g>
        <path class="st7" d="M182.3,166V150h2.4v13.9h7.4v2.2H182.3z"/>
      </g>
    </g>
  </g>
		`,
		color: "465B7B",
	},
}

func init() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	port = parser.Int("p", "port", &argparse.Options{Default: 9000})
	parser.Parse(os.Args)
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) (err error) {
		name := ctx.QueryParam("name")
		groupQuery := ctx.QueryParam("group")
		ddd := ctx.QueryParam("ddd")
		phone := ctx.QueryParam("phone")
		raffle := ctx.QueryParam("raffle-number")

		var ok bool
		var group group
		if group, ok = groups[groupQuery]; !ok {
			group = groups["1"]
		}

		filename := writeSVG(fmt.Sprintf(svgTemplate, group.color, textToSVG(name, 55, "st7"), textToSVG(ddd, 50, "st7"), textToSVG(phone, 50, "st7"), group.name, textToSVG(raffle, 300, "st3")))
		defer os.Remove(filename)

		return ctx.Inline(filename, filename)
	})
	e.GET("/Back.png", func(ctx echo.Context) (err error) {
		return ctx.Inline("Back.png", "Back.png")
	})

	e.HideBanner = true
	e.Start(fmt.Sprintf(":%d", *port))
}

func textToSVG(text string, size float64, color string) string {
	fontFamily := canvas.NewFontFamily("Roboto")
	if err := fontFamily.LoadFontFile("Roboto-Regular.ttf", canvas.FontRegular); err != nil {
		panic(err)
	}

	face := fontFamily.Face(size, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	path, _, err := face.ToPath(text)
	if err != nil {
		panic(err)
	}

	tpl := `<path class="%s" transform="scale(1, -1)" d="%s"/>`
	return fmt.Sprintf(tpl, color, path.ToSVG())
}

func writeSVG(SVGData string) string {
	file, err := ioutil.TempFile("./", "svg")
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(SVGData)
	return file.Name()
}
