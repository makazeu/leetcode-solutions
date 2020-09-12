func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	squareCenterX := float64(x2+x1) / 2
	squareCenterY := float64(y2+y1) / 2
	squareX2 := float64(x2) - squareCenterX
	squareY2 := float64(y2) - squareCenterY
	circleX := float64(x_center) - squareCenterX
	circleY := float64(y_center) - squareCenterY
	if circleX < 0 {
		circleX = -circleX
	}
	if circleY < 0 {
		circleY = -circleY
	}
	if circleX <= squareX2 && circleY <= squareY2 {
		return true
	}

	mx := circleX - squareX2
	my := circleY - squareY2
	if mx < 0 {
		mx = 0
	}
	if my < 0 {
		my = 0
	}
	return mx*mx+my*my <= float64(radius*radius)
}

