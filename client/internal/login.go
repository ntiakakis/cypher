package internal

import (
	"log"

	"edpasenidis.tech/cypher/internal/auth"
	"github.com/marcusolsson/tui-go"
)

func Login() {
	logo := `
                   _                 
                  | |                
  ____ _   _ ____ | | _   ____  ____ 
 / ___) | | |  _ \| || \ / _  )/ ___)
( (___| |_| | | | | | | ( (/ /| |    
 \____)\__  | ||_/|_| |_|\____)_|    
      (____/|_|                      
`
	user := tui.NewEntry()
	user.SetFocused(true)

	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)

	room := tui.NewEntry()

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("User"), tui.NewLabel("Password"))
	form.AppendRow(user, password)

	form.AppendRow(tui.NewLabel("Server"))
	form.AppendRow(room)

	status := tui.NewStatusBar("Ready.")

	login := tui.NewButton("[Login]")
	login.OnActivated(func(b *tui.Button) {
		status.SetText("Logging in...")
		status.SetText(auth.Login(room.Text(), user.Text(), password.Text()))
	})

	register := tui.NewButton("[Register]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, login),
		tui.NewPadder(1, 0, register),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 1, tui.NewLabel(logo)),
		tui.NewPadder(12, 0, tui.NewLabel("Welcome to Cypher! Login or register.")),
		tui.NewPadder(1, 1, form),
		buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(
		content,
		status,
	)

	tui.DefaultFocusChain.Set(user, password, room, login, register)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
