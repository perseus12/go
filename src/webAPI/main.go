package main

import (
	"html/template"
	"log"
	"net/http"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
	CompanyNumber    string
	Selection        string
}

func main() {
	http.HandleFunc("/", DisplayWebForm)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DisplayWebForm(w http.ResponseWriter, r *http.Request) {
	// Display some radio buttons to the user

	Title := "Your data stream"
	MyRadioButtons := []RadioButton{
		RadioButton{"radioSelect", "psc", false, false, "PSC"},
		RadioButton{"radioSelect", "officers", false, false, "Officers"},
		RadioButton{"radioSelect", "roa", false, false, "ROA"},
		RadioButton{"radioSelect", "insolvency", false, false, "Insolvency"},
		RadioButton{"radioSelect", "charges", false, false, "Charges"},
		RadioButton{"radioSelect", "exemptions", false, false, "Exemptions"},
		RadioButton{"radioSelect", "registers", false, false, "Registers"},
	}

	MyPageVariables := PageVariables{
		PageTitle:        Title,
		PageRadioButtons: MyRadioButtons,
	}

	t, err := template.ParseFiles("HTML/select.html") //parse the html file HTML/select.html
	if err != nil {                                   // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func UserSelected(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	companyNo := r.Form.Get("inputSelection")
	yourSelection := r.Form.Get("radioSelect")
	jsonData := apiCall(yourSelection, companyNo)

	Title := "lel"
	MyPageVariables := PageVariables{
		PageTitle:     Title,
		Answer:        jsonData,
		CompanyNumber: companyNo,
		Selection:     yourSelection,
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("HTML/select.html") //parse the html file HTML/select.html
	if err != nil {                                   // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
