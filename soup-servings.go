package main

import "fmt"

type Operation struct {
	code int
	a    int
	b    int
}

var ops = [4]Operation{
	{a: 100, b: 0},
	{a: 75, b: 25},
	{a: 50, b: 50},
	{a: 25, b: 75},
}

func soupServings(n int) float64 {
	if n >= 4800 {
		return 1
	}

	completedCombos := make(map[SoupCombo]float64)
	result := iterate(n, completedCombos, 0, 0)

	fmt.Println(result)
	fmt.Println(completedCombos)
	return result
}

func iterate(n int, completedCombos map[SoupCombo]float64, a int, b int) float64 {
	count := 0.0
	for _, op := range ops {
		currCombo := SoupCombo{
			a: a + op.a,
			b: b + op.b,
		}

		existingValue := completedCombos[currCombo]
		if existingValue != 0.0 {
			count += existingValue
		} else {
			if isFull(n, currCombo.a) && isFull(n, currCombo.b) {
				// The reason why we don't need to multiply this num by 0.25 is because it's a completed combo found within the step, but it isn't representetive of a completed step's value.
				// Remember, we only multiply by 0.25 at the end of the step, not necessarily if we find a solution within the step.
				// However, we multiply by 0.25 below because we completed an entire step & have found that, for this step, this is our holistic solution.
				count += 0.5
				completedCombos[currCombo] = 0.5
			} else if isFull(n, currCombo.a) {
				count += 1
				completedCombos[currCombo] = 1
			} else if !isFull(n, currCombo.b) {
				count += iterate(n, completedCombos, currCombo.a, currCombo.b)
			}
		}
	}

	if count > 0.0 {
		count *= 0.25
		currCombo := SoupCombo{
			a: a,
			b: b,
		}
		completedCombos[currCombo] = count
	}

	return count
}

func isFull(n int, amount int) bool {
	return amount >= n
}

type SoupCombo struct {
	a int
	b int
}

func main() {
	soupServings(50)
	soupServings(156)
	soupServings(100)
}
