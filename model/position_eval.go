package model

func eval(pos Position) int {
	eval := 0
	if []rune(pos.White())[1] == '9' {
		eval += 1000000
	}
	if []rune(pos.Black())[1] == '1' {
		eval -= 1000000
	}
	eval += distanceToFinish(pos, pos.Black(), '1')
	eval -= distanceToFinish(pos, pos.White(), '9')
	eval -= 2 * pos.BlackWallsLeft()
	eval += 2 * pos.WhiteWallsLeft()
	return eval
}
