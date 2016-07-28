package main

import "github.com/snikch/goodman/hooks"

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(h)
	server.Serve()
	defer server.Listener.Close()
}
