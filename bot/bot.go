package main

import (
	"github.com/gorilla/mux"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"

	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var r *raspi.Adaptor
var led *gpio.LedDriver

func main() {
	r = raspi.NewAdaptor()
	led = gpio.NewLedDriver(r, "31")

	router := mux.NewRouter()
	router.HandleFunc("/start", handleStart).Methods("GET")
	router.HandleFunc("/snapshot", handleSnapshot).Methods("GET")
	router.HandleFunc("/motion/on", handleMotionOn).Methods("GET")
	router.HandleFunc("/motion/off", handleMotionOff).Methods("GET")

	work := func() {
		led.Off()
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

func handleStart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit /start")
	fmt.Fprintln(w, "ok")

	go func() {
		led.On()
		startRecordingCmd()
		led.Off()
	}()
}

func handleSnapshot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit /snapshot")
	fmt.Fprintln(w, "ok")

	go func() {
		led.On()
		doSnapshot()
		led.Off()
	}()

}

func handleMotionOn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit /motion/on")
	fmt.Fprintln(w, "ok")

	go func() {
		led.On()
		doMotionOn()
	}()

}

func handleMotionOff(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit /motion/off")
	fmt.Fprintln(w, "ok")

	go func() {
		led.Off()
		doMotionOff()
	}()

}

func startRecordingCmd() {
	cmd := exec.Command("sh", "/home/pi/scripts/record.sh")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("RECORDING RESULTS: %q\n", out.String())

	if err != nil {
		log.Fatal(err)
	}
}

func doSnapshot() {
	cmd := exec.Command("sh", "/home/pi/scripts/snapshot.sh")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("SNAPSHOT RESULTS: %q\n", out.String())

	if err != nil {
		log.Fatal(err)
	}
}

func doMotionOn() {
	cmd := exec.Command("sudo", "systemctl", "start", "motion")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("START MOTION %q\n", out.String())

	if err != nil {
		log.Fatal(err)
	}
}

func doMotionOff() {
	cmd := exec.Command("sudo", "systemctl", "stop", "motion")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("STOPPING MOTION %q\n", out.String())

	if err != nil {
		log.Fatal(err)
	}
}
