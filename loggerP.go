package loggerP

import (
	"fmt"
	"time"
)

func LogWT(i interface{}) {
	fmt.Println("TIME: ", time.Now(), " MESSAGE: ", i)
}
