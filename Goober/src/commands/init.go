package commands

import (
	"fmt"
	"log"
	"os"
)

func Init() {
	err := os.Mkdir(".goobers", 0755)
	if err != nil {
		log.Fatal("Repository created dumbass")
	}

	os.Mkdir(".goobers/objects", 0755)
	os.Mkdir(".goobers/refs", 0755)
	fmt.Println("initialized repo")
}
