package _041_robot_bounded_in_circle

func isRobotBounded(instructions string) bool {
	//we consider the 4 directions as: 0=Up, 1=Right, 2=Down, 3=Left
	dir := 0 // initial facing up

	// this array will maintain the distance travelled on X and Y axes (from origin)
	// walked[0] -> Y distance, walked[1] -> X distance
	walked := [2]int{}

	for _, c := range instructions {
		if c == 'G' {
			if dir == 0 {
				walked[0] += 1
			} else if dir == 1 {
				walked[1] += 1
			} else if dir == 2 {
				walked[0] -= 1
			} else {
				walked[1] -= 1
			}
		} else if c == 'L' {
			dir = turnLeft(dir)
		} else {
			dir = turnRight(dir)
		}
	}

	if walked[0] == 0 && walked[1] == 0 {
		return true
	}

	if dir != 0 {
		return true
	}

	return false
}

func turnLeft(dir int) int {
	if dir == 0 {
		return 3
	} else {
		return dir - 1
	}
}

func turnRight(dir int) int {
	if dir == 3 {
		return 0
	} else {
		return dir + 1
	}
}
