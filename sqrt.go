package sqrt

import (
	"math/big"
)

func Sqrt(c *big.Int, rank *big.Int) (result *big.Int, steps int) {

	var prevResult *big.Int

	result = big.NewInt(1)
	steps = 0

	loopCount := 0

	for {
		steps++

		prevResult = result
		result = calculateRoundup(c, rank, result)

		compareResultWithPrev := result.Cmp(prevResult)
		if compareResultWithPrev == 0 {
			loopCount++
		} else {
			loopCount = 0
		}

		if steps != 1 && compareResultWithPrev == 1 {
			return prevResult, steps
		}

		if loopCount == 3 {
			break
		}
	}

	return result, steps
}

func calculateRoundup(c *big.Int, rank *big.Int, prevValue *big.Int) (*big.Int) {

	var (
		part1 *big.Int
		part2 *big.Int
	)

	part1 = new(big.Int).Mul( new(big.Int).Sub(rank, big.NewInt(1)), prevValue )
	part2 = c

	for i := int64(0); new(big.Int).Sub(rank, big.NewInt(1)).Cmp(big.NewInt(i)) > 0; i++ {
		part2 = new(big.Int).Div(part2, prevValue)
	}

	return new(big.Int).Div( new(big.Int).Add(part1, part2), rank )
}

