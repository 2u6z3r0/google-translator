This is a fork from zijiren233/google-translator
# google-translater

```go
package main

import (
	"context"
	"fmt"

	translater "github.com/zijiren233/google-translater"
	"golang.org/x/time/rate"
)

func translate(text string) string {
	translated, err := translater.Translate(
		text,
		translater.TranslationParams{
			From: "auto",
			To:   "en",
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return translated
}

func main() {
	l := rate.NewLimiter(100, 100)
	for {
		l.Wait(context.Background())
		go func() { fmt.Println(translate("测试")) }()
	}
}
```
