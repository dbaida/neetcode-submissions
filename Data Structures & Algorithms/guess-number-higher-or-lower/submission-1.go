/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    i, j := 1, n

	for {
		guessOption := i + (j - i) / 2
		res := guess(guessOption)
		if res == 0 {
			return guessOption
		}
		if res == -1 {
			j = guessOption - 1
		} else {
			i = guessOption + 1
		}
	}
}
