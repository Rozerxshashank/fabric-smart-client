package tv

// TV is a dummy service for the contribution process testbed (Task 11).
type TV struct {
	isOn      bool
	isNetflix bool
}

// New returns a new TV.
func New() *TV {
	return &TV{}
}

// On turns the TV on.
func (t *TV) On() string {
	t.isOn = true
	return "TV is now ON"
}

// Off turns the TV off.
func (t *TV) Off() string {
	t.isOn = false
	t.isNetflix = false
	return "TV is now OFF"
}

// NetflixOn turns Netflix on.
func (t *TV) NetflixOn() string {
	if !t.isOn {
		return "Cannot turn on Netflix: TV is OFF"
	}
	t.isNetflix = true
	return "Netflix is now ON"
}
