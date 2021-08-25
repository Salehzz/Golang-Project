package main

import (
	"net/http"
)

//duets handler for renders the duets.html
func Duets(w http.ResponseWriter, r *http.Request) {

	// define default duet options
	duetoptions := []ScaleOptions{
		ScaleOptions{Name: "Duet", Value: "G Major", IsDisabled: false, IsChecked: true, Text: "G Major"},
		ScaleOptions{Name: "Duet", Value: "D Major", IsDisabled: false, IsChecked: false, Text: "D Major"},
		ScaleOptions{Name: "Duet", Value: "A Major", IsDisabled: false, IsChecked: false, Text: "A Major"},
	}

	pageVars := PageVars{
		Title:         "Duets",
		Key:           "G Major",
		DuetImgPath:   "img/duet/gmajor.png",
		DuetAudioBoth: "mp3/duet/gmajorduetboth.mp3",
		DuetAudio1:    "mp3/duet/gmajorduetpt1.mp3",
		DuetAudio2:    "mp3/duet/gmajorduetpt2.mp3",
		DuetOptions:   duetoptions,
	}
	render(w, "duets.html", pageVars)
}

//duetshow handler for renders the duets.html
func DuetShow(w http.ResponseWriter, r *http.Request) {

	// define default duet options
	duetoptions := []ScaleOptions{
		ScaleOptions{Name: "Duet", Value: "G Major", IsDisabled: false, IsChecked: true, Text: "G Major"},
		ScaleOptions{Name: "Duet", Value: "D Major", IsDisabled: false, IsChecked: false, Text: "D Major"},
		ScaleOptions{Name: "Duet", Value: "A Major", IsDisabled: false, IsChecked: false, Text: "A Major"},
	}

	// Set a image path, this will be changed later.
	DuetImgPath := "img/duet/gmajor.png"
	DuetAudioBoth := "mp3/duet/gmajorduetboth.mp3"
	DuetAudio1 := "mp3/duet/gmajorduetpt1"
	DuetAudio2 := "mp3/duet/gmajorduetpt2"

	//r is url.Values which is a map[string][]string
	r.ParseForm()
	var dvalues []string
	for _, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			dvalues = append(dvalues, value) // stick each value in a slice I know the name of
		}
	}

	switch dvalues[0] {
	case "D Major":
		duetoptions = []ScaleOptions{
			ScaleOptions{Name: "Duet", Value: "G Major", IsDisabled: false, IsChecked: false, Text: "G Major"},
			ScaleOptions{Name: "Duet", Value: "D Major", IsDisabled: false, IsChecked: true, Text: "D Major"},
			ScaleOptions{Name: "Duet", Value: "A Major", IsDisabled: false, IsChecked: false, Text: "A Major"},
		}
		DuetImgPath = "img/duet/dmajor.png"
		DuetAudioBoth = "mp3/duet/dmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/dmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/dmajorduetpt2.mp3"
	case "G Major":
		duetoptions = []ScaleOptions{
			ScaleOptions{Name: "Duet", Value: "G Major", IsDisabled: false, IsChecked: true, Text: "G Major"},
			ScaleOptions{Name: "Duet", Value: "D Major", IsDisabled: false, IsChecked: false, Text: "D Major"},
			ScaleOptions{Name: "Duet", Value: "A Major", IsDisabled: false, IsChecked: false, Text: "A Major"},
		}
		DuetImgPath = "img/duet/gmajor.png"
		DuetAudioBoth = "mp3/duet/gmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/gmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/gmajorduetpt2.mp3"

	case "A Major":
		duetoptions = []ScaleOptions{
			ScaleOptions{Name: "Duet", Value: "G Major", IsDisabled: false, IsChecked: false, Text: "G Major"},
			ScaleOptions{Name: "Duet", Value: "D Major", IsDisabled: false, IsChecked: false, Text: "D Major"},
			ScaleOptions{Name: "Duet", Value: "A Major", IsDisabled: false, IsChecked: true, Text: "A Major"},
		}
		DuetImgPath = "img/duet/amajor.png"
		DuetAudioBoth = "mp3/duet/amajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/amajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/amajorduetpt2.mp3"
	}

	// set default page variables
	pageVars := PageVars{
		Title:         "Duets",
		Key:           "G Major",
		DuetImgPath:   DuetImgPath,
		DuetAudioBoth: DuetAudioBoth,
		DuetAudio1:    DuetAudio1,
		DuetAudio2:    DuetAudio2,
		DuetOptions:   duetoptions,
	}
	render(w, "duets.html", pageVars)
}
