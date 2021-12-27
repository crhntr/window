# `window` has packages to help writing WASM apps and SSE webpages

The `dom` package has standards inspired interfaces that can be implemented by wrapping `syscall/js` and the browser DOM.
`dom` does not import `syscall/js` so implementations need not use `GOOS=js` `GOARCH=wasm`. The `ast` package implements
most of the interfaces in `dom` and can be used for server side rendering (SSE).

The `dom` package can be used to create sse/wasm friendly libraries/frameworks.
The `attr` package is a (work in progress) library both for SSE and front end webapps.
