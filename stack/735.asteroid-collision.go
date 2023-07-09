package main

/*
 * @lc app=leetcode id=735 lang=golang
 *
 * [735] Asteroid Collision
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	output := []int{}
	rightStack := []int{}

	for i := 0; i < len(asteroids); i++ {
		asteroid := asteroids[i]
		if asteroid < 0 {
			sameSize := false
			for j := len(rightStack) - 1; j >= 0; j-- {
				if rightStack[j]+asteroid > 0 {
					break
				} else if rightStack[j]+asteroid == 0 {
					rightStack = rightStack[:len(rightStack)-1]
					sameSize = true
					break
				} else {
					rightStack = rightStack[:len(rightStack)-1]
				}
			}
			if len(rightStack) == 0 && !sameSize {
				output = append(output, asteroid)
			}
		} else {
			rightStack = append(rightStack, asteroid)
		}
	}
	if len(rightStack) != 0 {
		output = append(output, rightStack...)
	}
	return output
}

// @lc code=end
