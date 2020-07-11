package internal

import (
	"encoding/json"
	"log"

	"edpasenidis.tech/cypher/internal/auth"
	"github.com/marcusolsson/tui-go"
)

// LoginTokenBody ... Login response
type LoginTokenBody struct {
	Token string `json:"token"`
}

// Credentials ... Credential
type Credentials struct {
	Token  string
	Domain string
}

// Login ... User login
func Login() Credentials {
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

	email := tui.NewEntry()

	room := tui.NewEntry()

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("User"), tui.NewLabel("Password"))
	form.AppendRow(user, password)

	form.AppendRow(tui.NewLabel("Server"))
	form.AppendRow(room)

	status := tui.NewStatusBar("Ready.")

	var data LoginTokenBody

	login := tui.NewButton("[Login]")
	login.OnActivated(func(b *tui.Button) {
		status.SetText("Logging in...")
		err := json.Unmarshal([]byte(auth.Login(room.Text(), user.Text(), password.Text())), &data)
		if err != nil {
			status.SetText(err.Error())
			return
		}
		status.SetText("Press enter again to log-in.")
	})

	havesAccount := false
	register := tui.NewButton("[Register]")
	register.OnActivated(func(b *tui.Button) {
		if !havesAccount {
			havesAccount = true
			status.SetText("Write your e-mail to register.")
			form.AppendRow(tui.NewLabel("Email"))
			form.AppendRow(email)
			email.SetFocused(true)
		} else {
			auth.Register(room.Text(), user.Text(), password.Text(), email.Text())
			status.SetText("Registered")
		}
	})

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
	ui.SetKeybinding("Enter", func() {
		if len(data.Token) >= 10 {
			ui.Quit()
		}
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}

	return Credentials{Token: data.Token, Domain: room.Text()}
}
