package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

var ctx js.Value
var canvas js.Value
var myEntity int

func main() {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 30; i++ {
		currEntity := emptyEID
		createEntity(controller, position, draw, bot)
		controllers[currEntity].Speed = 1 + 3*rand.Float64()
		controllers[currEntity].DirX = 1
		draws[currEntity].Color = "blue"
		draws[currEntity].Width = 5 + 10*rand.Float32()
		draws[currEntity].Height = 5 + 10*rand.Float32()
		positions[currEntity].Y = 500 * rand.Float64()
	}

	ctx = js.Global().Get("document").
		Call("getElementById", "myCanvas").
		Call("getContext", "2d")

	myEntity = emptyEID
	createEntity(controller, position, draw)
	controllers[myEntity].Speed = 0.5
	draws[myEntity].Color = "red"
	draws[myEntity].Width = 5
	draws[myEntity].Height = 5

	js.Global().Get("document").
		Call("getElementById", "left").
		Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			controllers[myEntity].DirX--
			return nil
		}))

	js.Global().Get("document").
		Call("getElementById", "right").
		Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			controllers[myEntity].DirX++
			return nil
		}))

	mainLoop(js.Value{}, []js.Value{})

	c := make(chan bool)
	<-c
}

func mainLoop(this js.Value, args []js.Value) interface{} {

	canvas = js.Global().Get("document").
		Call("getElementById", "myCanvas")

	ctx.Call("clearRect", 0, 0, canvas.Get("width"), canvas.Get("height"))

	System()

	js.Global().Get("document").
		Call("getElementById", "myTextBox").
		Set("value", positions[myEntity].X)

	js.Global().Get("window").Call("requestAnimationFrame", js.FuncOf(mainLoop))
	return nil
}
