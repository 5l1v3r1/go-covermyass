package main

import (
	"fmt"
	"os"

	"covermyass/src/config"
	"covermyass/src/utils"
)

func displayMenu() {

	menu := `Welcome to Cover My Ass !

Select an option :

1) Clear your own logs
2) Permenently disable auth & bash history
3) Restore settings to default
99) Exit tool
`

	fmt.Println(menu)
}
func main() {
	displayMenu()
	run()

	os.Exit(0)
}

func run() {
	choice := utils.ReadInput("> ")

	switch choice {
	case "1":
		utils.Clear(config.GeneralLogs)
		utils.Clear(config.HistoryLogs)
	case "2":
		utils.Mock(config.AuthLogs)
		utils.Mock(config.HistoryLogs)
	case "3":
		utils.Unmock(config.AuthLogs)
		utils.Unmock(config.HistoryLogs)
	case "99":
		os.Exit(0)
	default:
		fmt.Println("[!] Unreconized option.")
		run()
	}
}
