//go:build js && wasm
// +build js,wasm

package window

import (
	"fmt"
	"net/url"
)

type location int

const Location = location(0)

func (document document) Location() location { return location(0) }

func (location) HREF() string {
	return win.Get("location").Get("href").String()
}

func (location) Protocol() string {
	return win.Get("location").Get("protocol").String()
}

func (location) Host() string {
	return win.Get("location").Get("host").String()
}

func (location) Port() string {
	return win.Get("location").Get("port").String()
}

func (location) PathName() string {
	return win.Get("location").Get("pathname").String()
}

func (location) Search() string {
	return win.Get("location").Get("search").String()
}

func (location) Hash() string {
	return win.Get("location").Get("hash").String()
}

func (location) SetHash(fragment string) {
	win.Get("location").Set("hash", fragment)
}

func (location) Origin() string {
	return win.Get("location").Get("origin").String()
}

func (location) Assign(url string) {
	win.Get("location").Call("assign", url)
}

func (location) Reload() {
	win.Get("location").Call("reload")
}

func (location) Replace(url string) {
	win.Get("location").Call("replace", url)
}

func (location) String() string {
	return win.Get("location").Call("toString").String()
}

// URL parses the URL string. It may panic if parse fails;
// however, this should not occur in practice. It does not
// cache the result so it is not super efficient to call
// URL successively.
func (location) URL() *url.URL {
	u, err := url.Parse(Location.String())
	if err != nil {
		panic(fmt.Errorf("url was not parsable: %w", err))
	}
	return u
}
