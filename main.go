package main

import("fmt")

func Fibbonaci(size int) []int{
	if size < 1{
		return nil
	}else{
		array := []int{1}
		for num := 0; num < size; num++ {
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

func main(){
	fmt.Println(Fibbonaci(6))
}
