package main

//go:generate go run generator.go Draw Position Controller
const (
	entitySize = 1000
)

var on [compNum][entitySize]bool

var emptyEID int = 0

func System() {
	SystemControl()
	SystemDraw()
}

func SystemControl() {
	for i := 0; i < entitySize; i++ {
		if on[controller][i] && on[position][i] {
			positions[i].X += controllers[i].DirX * controllers[i].Speed
			positions[i].Y += controllers[i].DirY * controllers[i].Speed

		}

	}

}

func SystemDraw() {
	for i := 0; i < entitySize; i++ {
		if on[draw][i] && on[position][i] {
			ctx.Call("beginPath")
			ctx.Call("moveTo", 100+positions[i].X, 0+positions[i].Y)
			ctx.Call("lineTo", 100+positions[i].X, 100+positions[i].Y)
			ctx.Call("stroke")

			//ctx.Call("clearRect", 0, 0, width, height);
			ctx.Set("fillStyle", draws[i].Color)
			ctx.Call("fillRect", positions[i].X, positions[i].Y, draws[i].Width, draws[i].Height)

		}
	}
}

func createEntity(IDs ...int) {
	for _, ID := range IDs {
		on[ID][emptyEID] = true

	}
	for ; emptyEID < entitySize; emptyEID++ {
		if checkEmpty(emptyEID) {

			break
		}

	}
}

func removeEntity(id int) {

	for tID := range on {
		on[tID][id] = false

	}

	if emptyEID < id {

		emptyEID = id
	}
}

func removeComponent(eID int, comp int) {
	on[comp][eID] = false
	if checkEmpty(eID) {
		if emptyEID < eID {

			emptyEID = eID
		}
	}

}

func checkEmpty(eID int) bool {

	for tID := range on {
		if on[tID][eID] {

			return false
		}
	}

	return true
}
