package main

import (
	"go/Crypto-Hacks/BIN"
	"go/Crypto-Hacks/Base"
	"go/Crypto-Hacks/Caeser"
	"go/Crypto-Hacks/Hex"

	"strconv"
	"fmt"
	"go/Crypto-Hacks/msg"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)


func main() {
app := tview.NewApplication()
	wm := winman.NewWindowManager()

	quitMsgBox := Msg.MsgBox("Confirmation", "Really quit?", []string{"Yes", "No"}, func(clicked string) {
		if clicked == "Yes" {
			app.Stop()
		}
	})
	wm.AddWindow(quitMsgBox)


	setFocus := func(p tview.Primitive) {
		go app.QueueUpdateDraw(func() {
			app.SetFocus(p)
		})
	}

	var createForm func(modal bool) *winman.WindowBase
	var counter = 0

	createForm = func(modal bool) *winman.WindowBase {
		counter++
		form := tview.NewForm()
		window := winman.NewWindow().
			SetRoot(form).
			SetResizable(true).
			SetDraggable(true).
			SetModal(modal)

		quit := func() {
			if wm.WindowCount() == 3 {
				quitMsgBox.Show()
				wm.Center(quitMsgBox)
				setFocus(quitMsgBox)
			} else {
				wm.RemoveWindow(window)
				setFocus(wm)
			}
		}

		display := 	tview.NewTextView().
		SetText("Output Here:").
		SetTextAlign(tview.AlignLeft)

		form.AddInputField("Enter Here", "", 40, nil, nil).
			AddInputField("Enter ROT (Caesar)", "", 40, nil, nil).
			AddFormItem(display).
			AddButton("Normal 2 Hex", func() {
				out := Hex.Normal2Hex(form.GetFormItem(0).(*tview.InputField).GetText())
				display.SetText(fmt.Sprintf( "Output Here:" + out))
			}).
			AddButton("Normal 2 BIN", func() {
				out := BIN.Normal2BIN(form.GetFormItem(0).(*tview.InputField).GetText())
						display.SetText(fmt.Sprintf( "Output Here:" + out))
			}).
			AddButton("Normal 2 Base64", func() {
				out := Base.Normal2base64(form.GetFormItem(0).(*tview.InputField).GetText())
						display.SetText(fmt.Sprintf( "Output Here:" + out))
			}).
			AddButton("Normal 2 Base32", func() {
				out := Base.Normal2base32(form.GetFormItem(0).(*tview.InputField).GetText())
						display.SetText(fmt.Sprintf( "Output Here:" + out))
			}).
			AddButton("Normal 2 Caesar", func() {
				
				cas, _ := strconv.Atoi(form.GetFormItem(1).(*tview.InputField).GetText())
				fmt.Println(cas)
				out := Caeser.Normal2Cipher(form.GetFormItem(0).(*tview.InputField).GetText(),cas)
						display.SetText(fmt.Sprintf( "Output Here:" + out))
			}).

			AddButton("New", func() {
				newWnd := createForm(false).Show()
				wm.AddWindow(newWnd)
				setFocus(newWnd)
			}).
			AddButton("Close", quit)

		title := fmt.Sprintf("Crypto-Hacks%d", counter)
		window.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)
		window.SetRect(2+counter*2, 2+counter, 50, 30)
		window.AddButton(&winman.Button{
			Symbol:    'X',
			Alignment: winman.ButtonLeft,
			OnClick:   quit,
		})

		var maxMinButton *winman.Button
		maxMinButton = &winman.Button{
			Symbol:    '▴',
			Alignment: winman.ButtonRight,
			OnClick: func() {
				if window.IsMaximized() {
					window.Restore()
					maxMinButton.Symbol = '▴'
				} else {
					window.Maximize()
					maxMinButton.Symbol = '▾'
				}
			},
		}
		window.AddButton(maxMinButton)
		wm.AddWindow(window)
		return window
	}

	for i := 0; i < 1; i++ {
		createForm(false).Show()
	}

	if err := app.SetRoot(wm, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
