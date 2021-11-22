package Data

func Fibbonaci(size int) []int{
	if size < 0{
		return nil
	}else{
		array := []int{0}
		for num := 0; num < size - 1; num++ {
			if num == 0 {
				array = append(array, 1)
			}else {
				newNum := array[num] + array[num-1]
				array = append(array, newNum)
			}
		}
		return array
	}
}


