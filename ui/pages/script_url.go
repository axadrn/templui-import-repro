package pages

import (
	"fmt"
	"time"
)

var scriptVersion = fmt.Sprintf("%d", time.Now().Unix())

func cacheBustedScriptURL(path string) string {
	return path + "?v=" + scriptVersion
}
