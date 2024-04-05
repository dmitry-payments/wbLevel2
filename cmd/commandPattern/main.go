package main

func main() {

	server := Server{}

	startC := StartCommand{
		device: server,
	}

	shutdownC := ShutdownCommand{
		device: server,
	}

	startC.run()
	shutdownC.run()
}
