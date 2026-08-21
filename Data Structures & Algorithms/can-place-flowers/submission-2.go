func canPlaceFlowers(flowerbed []int, n int) bool {
    isPrevPlotPlanted := false
    isCurrPlotPlanted := false
    isNextPlotPlanted := false

    for i := 0; i < len(flowerbed) && n > 0; i++ {
        isCurrPlotPlanted = flowerbed[i] == 1

        if !isCurrPlotPlanted && !isPrevPlotPlanted {
            if i < len(flowerbed) - 1 {
                isNextPlotPlanted = flowerbed[i + 1] == 1
            } else {
                isNextPlotPlanted = false
            }

            if !isNextPlotPlanted {
                n--
                isCurrPlotPlanted = true
            }
        }

        isPrevPlotPlanted = isCurrPlotPlanted
    }

    return n == 0
}