package main

import (
	"fmt"
	"os"

	"github.com/mikejwhitehead/sandbox/config"
)

func main()  {

	// if len(os.Args) != 2 {
	// 	log.Fatalf("Usage: %s <config>", os.Args[0])
	// }

	cfg := config.Load(os.Args[1])

	fmt.Printf("%d inputs", len cfg.In)
}

