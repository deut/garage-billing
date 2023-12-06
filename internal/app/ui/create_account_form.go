package ui

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/deut/garage-accounting/internal/models"
)

type CreateAccountForm struct {
	Window  fyne.Window
	Account *models.Account
}

func NewCreateAccountForm(w fyne.Window, a *models.Account) CreateAccountForm {

	return CreateAccountForm{Window: w, Account: a}
}

func (caf *CreateAccountForm) Build() fyne.CanvasObject {
	garageNumBind := binding.NewString()
	fullNameBind := binding.NewString()
	phoneBind := binding.NewString()
	addressBind := binding.NewString()

	garageNumText := widget.NewEntryWithData(garageNumBind)
	garageNumText.PlaceHolder = "garageNumText"
	garageNumText.Validator = func(s string) error {
		if s == "" {
			return errors.New("should not be blank")
		}

		return nil
	}
	fullNameText := widget.NewEntryWithData(fullNameBind)
	fullNameText.PlaceHolder = "firstNameText"
	fullNameText.Validator = func(s string) error {
		if s == "" {
			return errors.New("should not be blank")
		}

		return nil
	}
	phoneText := widget.NewEntryWithData(phoneBind)
	phoneText.PlaceHolder = "phoneText"
	phoneText.Validator = func(s string) error {
		if s == "" {
			return errors.New("should not be blank")
		}

		return nil
	}
	addressText := widget.NewEntryWithData(addressBind)
	addressText.PlaceHolder = "addressText"
	addressText.Validator = func(s string) error {
		if s == "" {
			return errors.New("should not be blank")
		}

		return nil
	}

	var err error
	submitBtn := widget.NewButton("createAccount", func() {
		if err := garageNumText.Validate(); err != nil {
			return
		}

		if err := fullNameText.Validate(); err != nil {
			return
		}

		if err := phoneText.Validate(); err != nil {
			return
		}
		addressText.Validate()
		if err := garageNumText.Validate(); err != nil {
			return
		}

		if err != nil {
			dialog.NewError(err, caf.Window).Show()
		}

		caf.Account.GarageNumber, err = garageNumBind.Get()
		if err != nil {
			err = fmt.Errorf("garageNumBind error: %w", err)
			return
		}
		caf.Account.FullName, err = fullNameBind.Get()
		if err != nil {
			err = fmt.Errorf("fullNameBind error: %w", err)
			return
		}
		caf.Account.PhoneNumber, err = phoneBind.Get()
		if err != nil {
			err = fmt.Errorf("phoneBind error: %w", err)
			return
		}
		caf.Account.Address, err = addressBind.Get()
		if err != nil {
			err = fmt.Errorf("addressBind error: %w", err)
			return
		}

		err = caf.Account.Insert()
		if err != nil {
			err = fmt.Errorf("account error: %w", err)
		}
	})

	layout := container.New(
		layout.NewGridLayoutWithColumns(5),
		fullNameText,
		phoneText,
		addressText,
		submitBtn,
	)

	return layout
}
