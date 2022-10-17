package main
const compNum = 4
var draws [entitySize]Draw
var positions [entitySize]Position
var controllers [entitySize]Controller
var bots [entitySize]Bot
const (
	draw = iota
	position
	controller
	bot
)