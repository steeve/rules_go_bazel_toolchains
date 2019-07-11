package main

import "fmt"

/*
void coucou() {
	printf("coucou2\n");
}
*/
import "C"

func main() {
	C.coucou()
	fmt.Println("coucou")
}
