package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 PM3:04:05 -0700", "2025-03-30 AM4:34:56 +0900")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.Format("1/2/2006 15:04:05 MST"))
}
