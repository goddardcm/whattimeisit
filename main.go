package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8080", RequestHandler{})
}

type RequestHandler struct{}

func (RequestHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {

	centralTime, _ := time.LoadLocation("America/Chicago")

	now := time.Now().In(centralTime)

	meridiem := "a m"
	hour := now.Hour()
	if hour > 12 {
		meridiem = "p m"
		hour -= 12
	}

	twimlResponse := TwiML{
		Say: fmt.Sprintf("The current time is %d, %02d, %s, on %s, %s, %d, %d",
			hour,
			now.Minute(),
			meridiem,
			now.Weekday().String(),
			now.Month().String(),
			now.Day(),
			now.Year(),
		),
	}
	bytes, _ := xml.Marshal(twimlResponse)

	w.Header().Set("Content-Type", "application/xml")
	w.Write(bytes)

}

type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	Say  string `xml:",omitempty"`
	Play string `xml:",omitempty"`
}
