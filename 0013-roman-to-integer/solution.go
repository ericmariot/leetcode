func romanToInt(s string) int {
    var result, lastValue, currentValue int
    romanToInt := map[uint8]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i := len(s) -1; i >= 0; i-- {
        currentValue = romanToInt[s[i]]
        if currentValue < lastValue {
            result -= currentValue
        }
        if currentValue >= lastValue {
            result += currentValue
        }

        lastValue = currentValue
    }

    return result
}
