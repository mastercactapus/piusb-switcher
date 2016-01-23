package main

import (
	"fmt"
	"github.com/jteeuwen/evdev"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

const idVendorFLIRC uint16 = 0x20a0
const idProductFLIRC uint16 = 0x0001
const FLIRCName = "flirc.tv flirc"

func OpenFLIRC(grab bool) {
	i := 0
	var dev *evdev.Device
	var err error
	for {
		node := fmt.Sprintf("/dev/input/event%d", i)
		dev, err = evdev.Open(node)
		i++
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatalln("could not find FLIRC device")
			}
			log.Fatalln(err)
		}
		if strings.TrimRight(dev.Name(), "\x00") != FLIRCName {
			dev.Close()
			continue
		}
		break
	}

	if grab && !dev.Grab() {
		log.Warnln("failed to grab device")
	}
	log.Infoln("using device:", strings.TrimRight(dev.Path(), "\x00"))
	for event := range dev.Inbox {
		if event.Type != evdev.EvKeys {
			continue
		}
		if event.Value != 0 {
			continue
		}

		log.Infoln("FLIRC event", event.Code)
		for i, code := range flircMap {
			if int(event.Code) == code {
				setState(i)
				break
			}
		}
	}
}
