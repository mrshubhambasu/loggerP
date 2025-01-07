package loggerP

import (
	"fmt"
	"time"
)

func LogWT(i interface{}) {
	fmt.Println(time.Now(), " - ", i)
}
defer func()