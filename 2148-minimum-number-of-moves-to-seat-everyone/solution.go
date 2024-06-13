func minMovesToSeat(seats []int, students []int) int {
    count := 0
    sort.Ints(seats)
    sort.Ints(students)

    for i := range(seats) {
        if seats[i] > students[i] {
            count += seats[i] - students[i]
            continue
        }

        count += students[i] - seats[i]
    }

    return count
}
