package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/marcusolsson/tui-go"
)

// User ... A user
type User struct {
	Username string `json:"username"`
}

// Users ... Users
type Users struct {
	Users []User `json:"users"`
}

// Post ... A post
type Post struct {
	ID       string `json:"id"`
	Username string `json:"username'`
	Message  string `json:"message"`
}

// MessagesBody ... Message fetch response
type MessagesBody struct {
	Messages []Post `json:"messages"`
}

var messages MessagesBody

// Room ... A room
func Room(domain, token string) {
	var data Users
	err := json.Unmarshal([]byte(Fetch(domain, token)), &data)

	if err != nil {
		log.Fatal(err)
	}

	sidebar := tui.NewVBox(
		tui.NewLabel("USERS"),
		tui.NewLabel(""),
	)

	for i := 0; i < len(data.Users); i++ {
		sidebar.Append(tui.NewLabel(data.Users[i].Username))
	}

	sidebar.SetBorder(true)

	history := tui.NewVBox()

	err = json.Unmarshal(Messages(domain, token), &messages)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, m := range messages.Messages {
		timestamp, err := strconv.ParseInt(m.ID, 10, 64)

		if err != nil {
			log.Fatal(err)
			continue
		}

		timestampStr := strconv.FormatInt(timestamp>>23, 10)

		history.Append(tui.NewHBox(
			tui.NewLabel(timestampStr),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.Username))),
			tui.NewLabel(m.Message),
			tui.NewSpacer(),
		))
	}

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		history.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("00:00")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", "visulelen"))),
			tui.NewLabel(e.Text()),
			tui.NewSpacer(),
		))
		input.SetText("")
	})

	root := tui.NewHBox(sidebar, chat)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
