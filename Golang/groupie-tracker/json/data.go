package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetData(c chan []Artists) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal()
	}
	var r []Artists
	json.Unmarshal(body, &r)
	c <- r
}

func GetLocations(link string) Locations {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return Locations{}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Locations{}
	}
	var r Locations
	json.Unmarshal(body, &r)
	return r
}

func GetDates(link string) Dates {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return Dates{}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Dates{}
	}
	var r Dates
	json.Unmarshal(body, &r)
	return r
}

func GetRelations(link string) Relations {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return Relations{}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Relations{}
	}
	var r Relations
	json.Unmarshal(body, &r)
	return r
}

func Relation(A Artists) []map[string]interface{} {
	l := GetLocations(A.Locations)
	d := GetDates(A.ConcertDates)
	r := GetRelations(A.Relations)
	rel := []map[string]interface{}{}
	for _, i := range l.Locations {
		date := []string{}
		for _, j := range d.Dates {
			for _, l := range r.DatesLocations[i] {
				if l == j[len(j)-10:] {
					date = append(date, j[len(j)-10:])
				}
			}
		}
		rel = append(rel, map[string]interface{}{
			"Locations": FormatLocations(i),
			"Dates":     date,
		})
	}
	return rel
}

func Sort(tab []Artists, on string, dir string) {
	if on == "name" {
		for i := 1; i < len(tab); i++ {
			for j := 0; j < i; j++ {
				if tab[i].Name < tab[j].Name {
					tab[j], tab[i] = tab[i], tab[j]
				}
			}
		}
	} else if on == "date" {
		for i := 1; i < len(tab); i++ {
			for j := 0; j < i; j++ {
				if tab[i].CreationDate < tab[j].CreationDate {
					tab[j], tab[i] = tab[i], tab[j]
				}
			}
		}
	}
}

func FormatLocations(s string) string{
	city := Space(strings.Split(s, "-")[0])
	country := Space(strings.Split(s, "-")[1])
	return city + ", " + country
}

func Space(s string) string{
	str := strings.ToUpper(string(s[0])) + s[1:]
	for i := 0; i < len(str); i++ {
		if str[i] == '_' {
			str = str[:i] + " " + strings.ToUpper(string(str[i+1])) + str[i+2:]
		}
	}
	return str
}

