package button

import "machine"

// LEDButton holds the pins
type LEDButton struct {
	button machine.Pin
	led    machine.Pin
}

// New creates a new led button
func New(button, led machine.Pin) LEDButton {
	return LEDButton{
		button: button,
		led:    led,
	}
}

// Configure will setup the button
func (lb *LEDButton) Configure() {
	lb.button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	lb.led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// ReadButton will return bool of button state
func (lb *LEDButton) ReadButton() bool {
	return !lb.button.Get()
}

// ButtonLEDOn will turn the buttons LED on
func (lb *LEDButton) ButtonLEDOn() {
	lb.led.Low()
}

// ButtonLEDOff will turn the buttons LED off
func (lb *LEDButton) ButtonLEDOff() {
	lb.led.High()
}