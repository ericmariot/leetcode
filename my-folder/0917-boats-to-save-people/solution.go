import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
    // order the array
    // if people == limit it takes a boat, rescueBoats++
    // also remove the person from people
    // iterate through the array from two points beginning and end
    // create new array with the sum of those points, order
    // if the sum of the first and last > limit
    // rescueBoats + len(people)

    sort.Ints(people)

    left := 0
    right := len(people)-1
    boats := 0

    for left <= right {
        if people[right] + people[left] <= limit {
            left++
        }

        right--
        boats++
    }

    return boats
}
