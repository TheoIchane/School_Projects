package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"web/ascii"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if banner == "" {
		tmpl = template.Must(template.ParseFiles("./templates/error400.html"))
		tmpl.Execute(w, http.StatusInternalServerError)
		return
	}
	if !VerifText(text) {
		tmpl = template.Must(template.ParseFiles("./templates/error500.html"))
		tmpl.Execute(w, http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Ascii": template.HTML(FormattoHtml(ascii.Ascii(text, banner))),
		"File":  true,
	}
	tmpl.Execute(w, data)
}

func Download(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./text.txt")
	if err != nil {
		tmpl := template.Must(template.ParseFiles("./templates/error500.html"))
		tmpl.Execute(w, http.StatusInternalServerError)
		return
	}
	l, err := file.Stat()
	if err != nil {
		tmpl := template.Must(template.ParseFiles("./templates/error500.html"))
		tmpl.Execute(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(int(l.Size())))
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base("text.txt"))

	// serve file out.
	http.ServeFile(w, r, "./text.txt")
}

func Error404(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/error404.html"))
	tmpl.Execute(w, nil)
}

func VerifText(s string) bool {
	for _, i := range s {
		if i == 10 || i == 13 {
			continue
		}
		if i < 32 || i > 126 {
			return false
		}
	}
	return true
}

func FormattoHtml(s string) string {
	format_str := "<p>"
	for _, i := range s {
		if i == 32 {
			format_str += "&nbsp;"
		} else if i == 10 {
			format_str += "<br>"
		} else {
			format_str += string(i)
		}
	}
	return format_str + "</p>"
}
