package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//all informaion about a page
type PageVars struct {
	Name          string
	Title         string
	Scalearp      string
	Key           string
	Pitch         string
	DuetImgPath   string
	ScaleImgPath  string
	GifPath       string
	AudioPath     string
	AudioPath2    string
	DuetAudioBoth string
	DuetAudio1    string
	DuetAudio2    string
	LeftLabel     string
	RightLabel    string
	ScaleOptions  []ScaleOptions
	DuetOptions   []ScaleOptions
	PitchOptions  []ScaleOptions
	KeyOptions    []ScaleOptions
	OctaveOptions []ScaleOptions
}

//all information about a Duet
type ScaleOptions struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

func main() {

	// serve everything in the css folder, the img folder and mp3 folder as a file
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/mp3/", http.StripPrefix("/mp3/", http.FileServer(http.Dir("mp3"))))

	// when navigating to /home it should serve the home page
	http.HandleFunc("/", Home)
	http.HandleFunc("/form", form)
	http.HandleFunc("/duets", Duets)
	http.HandleFunc("/duetshow", DuetShow)
	http.ListenAndServe(getPort(), nil)

}

// Detect $PORT and if present uses it for listen and serve else defaults to :8080
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

//render a page
func render(w http.ResponseWriter, tmpl string, pageVars PageVars) {

	tmpl = fmt.Sprintf("templates/%s", tmpl) // prefix the name passed in with templates/
	t, err := template.ParseFiles(tmpl)      //parse the template file

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, pageVars) //execute the template and pass in the variables

	if err != nil { // if there is an error
		log.Print("template executing error: ", err)
	}
}
