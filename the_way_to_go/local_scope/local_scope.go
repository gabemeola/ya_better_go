package main

/**
 * Should output "GMG", because function m() is setting a new var locally
 *
 */

var a = "G"

func main() {
	n()
	m()
	n()
	println()
}

func n() {
	print(a)
}

func m() {
	a := "M"
	print(a)
}
