// Demonstrates read timeout (or "idle" timeout)
package main

func main() {
	port := "1234"
	responder(port)
	requestor(port)
}
