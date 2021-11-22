package fibbonaci

func Fibbonaci(size int) []int {

	if size <= 0 {
		return nil
	} else {
		array := []int{0}

		if size > 1 {
			array = append(array, 1)
		}
		for num := 1; num < size-1; num++ {
			newNum := array[num] + array[num-1]
			array = append(array, newNum)
		}
		return array
	}
}
