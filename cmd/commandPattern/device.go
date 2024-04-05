package main

type Device interface {
	start()
	shutdown()
}
