package main

import (
	"github.com/hybridgroup/gobot/platforms/raspi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
	"time"
)

var (
	pin        int
	stateCount int
	delay      time.Duration
	bindAddr   string
	flircMap   []int

	currentState int
	a            *raspi.RaspiAdaptor
)

var (
	mainCmd = &cobra.Command{
		Run: run,
	}
)

func run(cmd *cobra.Command, args []string) {
	a = raspi.NewRaspiAdaptor("raspi")
	err := a.DigitalWrite(strconv.Itoa(pin), 0)
	if err != nil {
		panic(err)
	}
	if len(flircMap) > 0 {
		go OpenFLIRC(true)
	}
	log.Fatalln(http.ListenAndServe(bindAddr, http.HandlerFunc(ServeHTTP)))
}

func incr() {
	err := a.DigitalWrite(strconv.Itoa(pin), 1)
	if err != nil {
		panic(err)
	}
	time.Sleep(delay)
	err = a.DigitalWrite(strconv.Itoa(pin), 0)
	if err != nil {
		panic(err)
	}
	time.Sleep(delay)
	currentState = (currentState + 1) % stateCount
}

func setState(newState int) {
	log.WithField("state", newState).Infoln("changing state")
	if newState >= stateCount {
		panic("invalid state")
	}
	for currentState != newState {
		incr()
	}
}

func main() {
	mainCmd.Flags().IntVarP(&pin, "pin", "p", 14, "Pin number. Pin used to toggle the usb switcher")
	mainCmd.Flags().IntVarP(&stateCount, "state-count", "c", 4, "Number of states. This should equal the number of states on the USB switch.")
	mainCmd.Flags().DurationVarP(&delay, "delay", "d", time.Millisecond*50, "Toggle delay. Time between switch toggles.")
	mainCmd.Flags().StringVarP(&bindAddr, "bind-address", "b", ":8080", "Bind address. The address:port to bind to for HTTP requests.")
	mainCmd.Flags().IntSliceVar(&flircMap, "flirc", []int{}, "Map key codes from FLIRC to states. The first code will trigger state 0 and so on.")
	mainCmd.Execute()
}
