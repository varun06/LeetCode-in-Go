package main

import (
	"log"
)

// 程序辅助设置
const (
<<<<<<< HEAD
	VERSION = "7.0.5"
=======
	VERSION = "7.0.8"
>>>>>>> f33c3a477711033e1c5c5c04e72ce2c3c83f449e
)

func main() {
	log.Printf("Hi, %s. I'm %s\n", getConfig().Username, VERSION)

	cli := new(CLI)
	cli.Run()
}
