package main

func main() {

}

func addToSlice(a []string) {
	b := append(a, "1")
	*a = *b
}
