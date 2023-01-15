package main

import (
	"strconv"
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/vjeantet/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	output := ""
	input := widget.NewLabel(output)
	var historyArray []string
	var historymsg string = ""
	var historyBool bool = false
	historyBtn := widget.NewButton("History", func() {

		if historyBool == true {
			historymsg = ""
		} else {
		for i := 0; i < len(historyArray); i++ {
			historymsg += historyArray[i] + " \n"
		}
		
	}
		output = historymsg
		input.SetText(output) 
		historyBool = !historyBool
	})
	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}

	})

	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	OpenBracketBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	closeBracketBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	DivideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	Numeric7 := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	Numeric8 := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	Numeric9 := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	Numeric4 := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	Numeric5 := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	Numeric6 := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	Numeric1 := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	Numeric2 := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	Numeric3 := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	Numeric0 := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)

	})
	addBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	subtractBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	
	equalsOperatorBtn := widget.NewButton("=", func() {
		var isIncorrect bool = false;
		var exp string = output
		expression, err := govaluate.NewEvaluableExpression(output)
		var ans string = ""
		if err != nil {
			isIncorrect = true
			output = "Invalid input"
			input.SetText(output)
			time.Sleep(2 * time.Second)
			output = ""
			input.SetText(output)
		} else {
			result, err := expression.Evaluate(nil)
			if err != nil {
				isIncorrect = true
				output = "Invalid input"
				input.SetText(output)
				time.Sleep(2 * time.Second)
				output = ""
				input.SetText(output)
			} else {
				ans = strconv.FormatFloat(result.(float64), 'f', -1, 64)
				output = ans
				input.SetText(output)
			}
		}
		if isIncorrect {
			return
		}
		var msg string = exp + " = " + ans
		historyArray = append(historyArray, msg)

	})
	equalsOperatorBtn.Importance = widget.HighImportance
	w.SetContent(container.NewVBox(
		input,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				closeBracketBtn,
				OpenBracketBtn,
				DivideBtn,
			),
			container.NewGridWithColumns(4,
				Numeric7,
				Numeric8,
				Numeric9,
				multiplyBtn,
			),
			container.NewGridWithColumns(4,
				Numeric4,
				Numeric5,
				Numeric6,
				subtractBtn,
			),
			container.NewGridWithColumns(4,
				Numeric1,
				Numeric2,
				Numeric3,
				addBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					Numeric0,
					dotBtn,
				),
				equalsOperatorBtn,
			),
		),
	))

	w.ShowAndRun()
}
