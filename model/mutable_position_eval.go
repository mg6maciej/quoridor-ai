package model

func (pos *MutablePosition) Eval() int {
	eval := 0
	if []rune(pos.White())[1] == '9' {
		eval += 1000000
	}
	if []rune(pos.Black())[1] == '1' {
		eval -= 1000000
	}
	eval += pos.distanceToFinish(pos.Black(), '1')
	eval -= pos.distanceToFinish(pos.White(), '9')
	eval -= 2 * pos.BlackWallsLeft()
	eval += 2 * pos.WhiteWallsLeft()
	return eval
}
