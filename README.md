# piusb-switcher

Simple interface for controlling a USB switch with a raspberry pi.

## Usage

Running `piusb-switcher` on a raspberry pi by default will listen on `:8080` and render a simple html page for controlling a switch with 4 states.

A single pin is used to control the select button on the switch.

For example, if the current state is `0` and `2` is selected the pin will be toggled twice.

