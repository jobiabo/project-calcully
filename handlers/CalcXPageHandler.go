package handler

import (
	"calcully/calcully"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("template/calculator.html"))

func CalcXPageHandler(w http.ResponseWriter, r *http.Request) {
	data := PageOutput{}
	if r.Method == http.MethodPost {
		r.ParseForm()
		exp := r.FormValue("expression")

		log.Println(exp)

		switch {
		case strings.HasPrefix(exp, "sqrt"):
			result, err := calcully.Sqrt(exp)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			resultcleared := fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp

		case strings.HasPrefix(exp, "sin"):
			result, err := calcully.Sin(exp)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			resultcleared := fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp

		case strings.HasPrefix(exp, "cos"):
			result, err := calcully.Cos(exp)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			resultcleared := fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp

		case strings.HasPrefix(exp, "tan"):
			result, err := calcully.Tan(exp)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			resultcleared := fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp

		case strings.HasPrefix(exp, "log"):
			result, err := calcully.Log(exp)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			resultcleared := fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp

		case strings.HasPrefix(exp, "x²("):
			fmt.Println(exp)
			result, err := (calcully.Power(exp))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
			resultcleared := fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp

		default:
			resultcleared := ""
			cleared, odd := calcully.Check(exp)

			result, err := calcully.Solve(odd, cleared)
			if err != nil {
				resultcleared = "Math error"
				resultcleared = fmt.Sprintf("%s", resultcleared)
				data.Result = (resultcleared)
				log.Println(resultcleared)
				tmpl.Execute(w, data)

				return
			}

			resultcleared = fmt.Sprintf("%f", result)
			data.Result = (resultcleared)
			data.Input = exp
		}

	}

	tmpl.Execute(w, data)

}
