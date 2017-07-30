package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"log"
	"net/http"
	"time"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "31")

	router := mux.NewRouter()
	router.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		led.On()

		fmt.Println("hit /start")
		fmt.Fprintln(w, "ok")

		time.Sleep(time.Second * 2)
		led.Off()

	}).Methods("GET")

	work := func() {

		fmt.Println("listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}

	robot := gobot.NewRobot("SafeWordBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	robot.Start()

}
