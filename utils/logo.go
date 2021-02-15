package utils

import (
	"fmt"
	"runtime"
)

func PrintLogo() {
	switch runtime.GOOS {
	case "darwin":
		colorPrint()
	case "windows":
		textPrint()
	case "linux":
		colorPrint()
	}
}

func colorPrint() {
	fmt.Println(Blue("________                __             _____ __________.___ " +
		"\n\\_____  \\  __ _______  |  | __ ____   /  _  \\\\______   \\   |\n /  " +
		"/ \\  \\|  |  \\__  \\ |  |/ // __ \\ /  /_\\  \\|     ___/   |\n/   \\_/. " +
		" \\  |  // __ \\|    <\\  ___//    |    \\    |   |   |\n\\_____\\ \\_/____/" +
		"(____  /__|_ \\\\___  >____|__  /____|   |___|\n       \\__>          \\/    " +
		" \\/    \\/        \\/              "))
	fmt.Println(Red("360 Quake API") + "   " + Green("Author:4ra1n"))
}

func textPrint() {
	fmt.Println("________                __             _____ __________.___ " +
		"\n\\_____  \\  __ _______  |  | __ ____   /  _  \\\\______   \\   |\n /  " +
		"/ \\  \\|  |  \\__  \\ |  |/ // __ \\ /  /_\\  \\|     ___/   |\n/   \\_/. " +
		" \\  |  // __ \\|    <\\  ___//    |    \\    |   |   |\n\\_____\\ \\_/____/" +
		"(____  /__|_ \\\\___  >____|__  /____|   |___|\n       \\__>          \\/    " +
		" \\/    \\/        \\/              ")
	fmt.Println("360 Quake API   Author:4ra1n")
}
