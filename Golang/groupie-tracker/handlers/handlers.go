package handlers

import (
	"fmt"
	"groupie/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var Artists = []json.Artists{}

func Data() {
	channel := make(chan []json.Artists)
	go json.GetData(channel)
	Artists = <-channel
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound )
		return
	}
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	// searched := r.FormValue("search")
	// tab := Search(searched)
	// sort := r.FormValue("sort")
	// on := ""
	// dir := ""
	// if sort != "" {
	// 	on = strings.Split(sort, ":")[0]
	// 	dir = strings.Split(sort, ":")[1]
	// }
	// json.Sort(tab, on, dir)
	data := map[string]interface{}{
		"Artist": Artists,
	}
	tmpl.Execute(w, data)
}

func Error(w http.ResponseWriter, err int) {
	tmpl := template.Must(template.ParseFiles("./templates/error.html"))
	data := map[string]interface{}{}
	var content string
	if err == 404 {
		w.WriteHeader(http.StatusNotFound)
		content = ": Page Not Found"
		data["Error"] = "404"
		data["Data"] = "Page Not Found"
	} 
	if err == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		content = ": Something went wrong"
		data["Error"] = "500"
		data["Data"] = "Something went wrong"
	}
	if err == 400 {
		w.WriteHeader(http.StatusBadRequest)
		content = ": Bad Request"
		data["Error"] = "400"
		data["Data"] = "Bad Request"
	}
	fmt.Println("Error",err, content)
	tmpl.Execute(w, data)
}

func Artist(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/artist.html"))
	id := r.URL.Path[8:]
	id_int, _ := strconv.Atoi(id)
	if id_int > 52 || id_int < 1 {
		Error(w, http.StatusInternalServerError)
		return
	}
	id_int = id_int - 1
	prev := id_int
	next := id_int + 2
	if id_int == 0 {
		prev = 52
	}
	if id_int == 53 || next == 53 {
		next = 1
	}
	data := map[string]interface{}{
		"Next":         strconv.Itoa(next),
		"Prev":         strconv.Itoa(prev),
		"Image":        Artists[id_int].Image,
		"Name":         Artists[id_int].Name,
		"CreationDate": Artists[id_int].CreationDate,
		"FirstAlbum":   Artists[id_int].FirstAlbum	,
		"Members":      Artists[id_int].Members,
		"Concerts":     json.Relation(Artists[id_int]),
	}
	tmpl.Execute(w, data)
}

func Search(data string) []json.Artists {
	temp := []json.Artists{}
	if data == "" {
		return Artists
	}
	if !isNumeric(data) {
		for _, i := range Artists {
			for _, j := range strings.Split(i.Name, " ") {
				if len(data) > len(j) {
					continue
				}
				if strings.EqualFold(data, j[:len(data)]) {
					temp = append(temp, i)
					break
				}
			}

		}
	} else {
		for _, i := range Artists {
			date := strconv.Itoa(i.CreationDate)
			album := i.FirstAlbum[6:]
			if len(data) > len(album) && len(data) > len(date) {
				continue
			}
			if strings.EqualFold(data, album[:len(data)]) || strings.EqualFold(data, date[:len(data)]) {
				temp = append(temp, i)
			}
		}
	}

	return temp
}

func isNumeric(s string) bool {
	for _, i := range s {
		if (i < '0' || i > '9') && i != '-' {
			return false
		}
	}
	return true
}
