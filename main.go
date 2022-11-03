package main

import "log"

func main() {

	defer func() {
		recovery := recover()
		if recovery != nil {
			log.Fatal("")
		}
	}()

	initLog()

	log.Println("\n\n\n[main]--------------START--------------")


	
	log.Println("\n\n\n[main]--------------END--------------")

}
