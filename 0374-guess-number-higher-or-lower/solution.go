/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    low, high := 0, n

    for low <= high {
        g := (low + high) / 2
        res := guess(g)
        
        if res == 0 {
            return g
        }

        if res == 1 {
            low = g + 1
        }

        if res == -1 {
            high = g - 1
        }
    }

    return -1
}
