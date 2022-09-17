package screen

import (
	"fmt"
	"infinity-dog/database"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var services = widgets.NewList()
var messages = widgets.NewList()
var serviceItems = []database.Service{}

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	services.SelectedRowStyle.Fg = ui.ColorWhite
	services.SelectedRowStyle.Bg = ui.ColorMagenta
	services.TextStyle.Fg = ui.ColorWhite
	services.TextStyle.Bg = ui.ColorBlack
	serviceItems = database.ServicesByTotalBytes()
	for _, item := range serviceItems {
		services.Rows = append(services.Rows, fmt.Sprintf("% 9d %s", item.TotalBytes, item.Name))
	}
	messages.SelectedRowStyle.Fg = ui.ColorWhite
	messages.SelectedRowStyle.Bg = ui.ColorMagenta
	messages.TextStyle.Fg = ui.ColorWhite
	messages.TextStyle.Bg = ui.ColorBlack

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/3, services),
			ui.NewCol((1.0/3)*2, messages),
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
			case "<Enter>":
				handleEnter()
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			}
		}
		ui.Render(grid)
	}
}

func handleEnter() {
	serviceName := serviceItems[services.SelectedRow].Name
	items := database.MessagesFromService(serviceName)
	messages.Rows = []string{}
	for _, item := range items {
		messages.Rows = append(messages.Rows, item.Both)
	}
}
