package bridge

import "testing"

func Test(t *testing.T) {
	redCircle := &Circle{x: 100, y: 100, radius: 10, drawAPI: &RedCircle{}}
	blueCircle := &Circle{x: 100, y: 100, radius: 10, drawAPI: &BlueCircle{}}

	redCircle.Draw()
	blueCircle.Draw()
}
