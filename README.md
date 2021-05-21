⚠️  THIS IS NOT STABLE YET. I AM MAKING MAJOR API CHANGES OFTEN. ⚠️ 

# Window
A Go package for Document Object Model (DOM) stuff when GOOS=js and GOARCH=wasm.

## Examples

### Counter

```go
package main

import (
	"strconv"

	"github.com/crhntr/window"
)

func main() {
	window.AddEventListenerFunc("click", func(event window.Event) {
		target := event.TargetElement()

		switch {
		case target.HasAncestor("#increment"):
			n := getNumber()
			n++
			setNumber(n)
		case target.HasAncestor("#decrement"):
			n := getNumber()
			n--
			setNumber(n)
		}  
	})
}

func getNumber() int {
	countEl := window.Document.QuerySelector("#count")
	n, _ := strconv.Atoi(countEl.InnerHTML())
	return n
}

func setNumber(n int) {
	countEl := window.Document.QuerySelector("#count")
	countEl.SetInnerHTMLf("%d", n)
}
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Counter</title>

    <script>go_init('/assets/wasm/counter.wasm')</script>
</head>
<body>
    <div id="count">420</div>
    <button id="increment">+</button>
    <button id="decrement">-</button>
</body>
```
