package main
const compNum = 3
var draws [entitySize]Draw
var positions [entitySize]Position
var controllers [entitySize]Controller
const (
	draw = iota
	position
	controller
)