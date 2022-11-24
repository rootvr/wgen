package utils

import (
	"fmt"
	"os"
)

func Kill(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
