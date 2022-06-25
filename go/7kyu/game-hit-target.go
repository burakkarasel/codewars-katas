package kata

func Solution(mtrx [][]rune) bool {
	arrIndex, arrNum := 0, 0
	goalIndex, goalNum := 0, 0
	for i := 0; i < len(mtrx); i++ {
		for index, v := range mtrx[i] {
			if v == '>' {
				arrNum = i
				arrIndex = index
			} else if v == 'x' {
				goalNum = i
				goalIndex = index
			}
		}
	}
	return arrNum == goalNum && arrIndex < goalIndex
}
