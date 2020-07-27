// +build js,wasm

package window

var win = js.Global()

func JSValue() js.Value                           { return win }
func Get(key string) js.Value                     { return win.Get(key) }
func Set(key string, value interface{})           { win.Set(key, value) }
func Call(m string, args ...interface{}) js.Value { return win.Call(m, args...) }
func Type() js.Type                               { return win.Type() }

func AddEventListener(eventName string, listener EventListener) {
	win.Call("addEventListener", eventName, js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener.HandleEvent(Event(args[0]))
		return nil
	}))
}

func AddEventListenerFunc(eventName string, listener EventListenerFunc) {
	win.AddEventListener(eventName, listener)
}

func AddEventListenerChannel(eventName string, c chan Event) {
	win.AddEventListenerFunc(eventName, func(event Event) {
		c <- event
	})
}
