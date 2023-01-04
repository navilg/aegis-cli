package internal

import (
	"github.com/pquerna/otp"
	"github.com/rivo/tview"
)

var (
	form *tview.Form
	app  *tview.Application
	list *tview.List
)

func PasswordTUI() []byte {

	app = tview.NewApplication()
	form = tview.NewForm()
	form.AddPasswordField("Aegis Password", "", 20, '*', nil)
	form.AddButton("OK", func() {
		app.Stop()
	})
	form.SetBorder(true).SetTitle(" Authentication ").SetTitleAlign(tview.AlignLeft)
	form.SetCancelFunc(func() {
		app.Stop()
	})
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	password := form.GetFormItemByLabel("Aegis Password").(*tview.InputField).GetText()
	return []byte(password)
}

func ItemsTUI(db DB) {
	app = tview.NewApplication()
	list = tview.NewList()
	list.AddItem("Quit", "", rune(0), func() {
		app.Stop()
	})
	for _, entry := range db.Entries {
		list.AddItem(entry.Issuer, entry.Name, rune(0), func() {
			app.Stop()
			OTPTUI(entry.Info.Secret, entry.Info.Algo, entry.Info.Digits, entry.Info.Period)
		})
	}

	list.SetBorder(true).SetTitle(" Items ").SetTitleAlign(tview.AlignLeft)

	list.SetDoneFunc(func() {
		app.Stop()
	})

	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func OTPTUI(secret, algo string, digits, period int) {
	app = tview.NewApplication()
	form = tview.NewForm()

	otp := GenerateTOTP(secret, algo, otp.Digits(digits), uint(period))
	form.AddTextArea("OTP", otp, 8, 1, 10, nil)
	// form.AddButton(otp, func() {
	// 	app.Stop()
	// })
	form.AddButton("Done", func() {
		app.Stop()
	}).SetButtonsAlign(tview.AlignCenter)
	form.SetBorder(true).SetTitle(" OTP ").SetTitleAlign(tview.AlignCenter)
	form.SetCancelFunc(func() {
		app.Stop()
	})

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
