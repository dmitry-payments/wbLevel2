package main

type Command interface {
	run()
}

type StartCommand struct {
	device Device
}

func (sc StartCommand) run() {
	sc.device.start()
}

type ShutdownCommand struct {
	device Device
}

func (sc ShutdownCommand) run() {
	sc.device.shutdown()
}
