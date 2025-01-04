package main

import (
	"fmt"
	"runtime"
)

func main() {
	//mengambil informasi sistem operasi
	sistemOperasi := runtime.GOOS
	switch sistemOperasi {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Mac OS")
	case "windows":
		fmt.Println("Windows")

	}

}
