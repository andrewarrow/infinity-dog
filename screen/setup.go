package screen

import (
	"fmt"
	"infinity-dog/dog"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	services := widgets.NewList()
	services.SelectedRowStyle.Fg = ui.ColorWhite
	services.SelectedRowStyle.Bg = ui.ColorMagenta
	services.TextStyle.Fg = ui.ColorWhite
	services.TextStyle.Bg = ui.ColorBlack
	items := dog.ServicesHitsFromSql()
	for _, item := range items {
		services.Rows = append(services.Rows, fmt.Sprintf("% 9d %s", item.Hits, item.Name))
	}

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, services),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				services.ScrollDown()
			case "k", "<Up>":
				services.ScrollUp()
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			}
		}
		ui.Render(grid)
	}
}
