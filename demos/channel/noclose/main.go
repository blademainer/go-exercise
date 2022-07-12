package main

func main() {
	for {
		createNoClosedChannel()
	}
}

func createNoClosedChannel() {
	_ = make(chan struct{})
	return
}
