# piusb-switcher

Simple interface for controlling a USB switch with a raspberry pi.

## Usage

Running `piusb-switcher` on a raspberry pi by default will listen on `:8080` and render a simple html page for controlling a switch with 4 states.

A single pin is used to control the select button on the switch.

For example, if the current state is `0` and `2` is selected the pin will be toggled twice.

## FLIRC Support

To control via FLIRC, specify the `--flirc` flag with a list of comma-separated HID codes. See below for example usage.

## Example

```bash
piusb-switcher -p 8 -d 10ms --flirc 1,2,3,4
```

|Argument |Description|
|---  |---- |
|-p 8  |Use **physical** pin `#8` to toggle switch state|
|-d 10ms|Delay 10ms between button presses (up and down)|
|--flirc 1,2,3,4|HID key `1` will set state `0`, key `2` will set state `1`, etc...|

