package main

var (
	message string = "Hello Go!"
)

func main() {
	println(message)
}

func init() {
	message = "Hello there"
}
