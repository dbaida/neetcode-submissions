func replaceElements(arr []int) []int {
    res := make([]int, 0, len(arr))

    for i := 0; i < len(arr) - 1; i++ {
        max := 0

        for j := i + 1; j < len(arr); j++ {
            if arr[j] > max {
                max = arr[j]
            }
        }
        res = append(res, max)
    }
    return append(res, -1)
}
