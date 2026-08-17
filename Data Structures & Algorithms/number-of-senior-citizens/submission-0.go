func countSeniors(details []string) int {
    count := 0

    for i := 0; i < len(details); i++ {
        detail := details[i]
        age, _ := strconv.Atoi(detail[len(detail)-4:len(detail)-2])
        if age > 60 {
            count++
        }
    }
    return count
}