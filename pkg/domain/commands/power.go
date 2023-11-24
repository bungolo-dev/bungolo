// The commands subdomin package of the bungolow domain provides all common commands.
// Commands include actions that are not only applicable to sources or displays but potentially sources, displays, rooms, etc.
package commands

import "errors"

const (
	Off     PowerState = 0
	On      PowerState = 1
	Warming PowerState = 2
	Cooling PowerState = 3
)

// The Power commands used to power on and power off a device.
type PowerControl interface {
	// Powers on the device
	PowerOn() PowerStatus
	// Powers off the device
	PowerOff() PowerStatus
}

// Provides required commands to query a device for the current power status.
type PowerFeedback interface {
	// Queries the current state of the devices power.
	// The device should report its current state opon request.
	PowerQuery() PowerStatus

	// Provides a registration method to request unsolisisited power status updates.
	// Once a device has registered for power updates each consumer callback should be invoked when changes to the power status occur
	ReguestPowerUpdates(*PowerConsumer)
}

// The power consumer interface is used to consumer power status updates.
type PowerConsumer interface {
	// Executed when the device has powered on.
	OnPowerOn()
	// Executed when the device has been powered off.
	OnPowerOff()
	// Executed when the device is cooling down.
	OnPowerCooling()
	// Executed when the device is warming up.
	OnPowerWarming()
}

// Stores the status of a devices current power status.
type PowerStatus struct {
	Off     bool
	On      bool
	Warning bool
	Cooling bool
}

// The power state of a device.
type PowerState int

// Creates a new power status stuct with the provided state
func New(ps PowerState) (PowerStatus, error) {
	switch ps {
	case 0:
		return PowerStatus{Off: true}, nil
	case 1:
		return PowerStatus{On: true}, nil
	case 2:
		return PowerStatus{Warning: true}, nil
	case 3:
		return PowerStatus{Cooling: true}, nil

	}
	return PowerStatus{Off: true}, errors.New("INVALID POWER STATE")
}
