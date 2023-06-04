package algorithnm

import "fmt"

func Quick(sli []int) {
	sort(0, len(sli)-1, sli)
}

func sort(l, r int, sli []int) {
	if l > r || len(sli[l:r]) < 2 {
		return
	}
	pivot := sli[l]
	left := l + 1
	fmt.Println(pivot)
	fmt.Println("start", sli)
	//right := r
	//left := l
	for i := l + 1; i <= r; i++ {
		//fmt.Println("++")
		//if l > r {
		//	break
		//}
		if sli[i] < pivot {
			//fmt.Println("--")
			sli[i], sli[left] = sli[left], sli[i]
			left++
		}
		fmt.Println("for : ", i, sli)
		//l++
		//r--
	}

	sli[l], sli[left-1] = sli[left-1], sli[l]
	sort(l, left-2, sli)
	sort(left, r, sli)

}

//func Quick(sli []int) {
//	sort(0, len(sli)-1, sli)
//}
//
//func sort(left, right int, sli []int) {
//	if left >= right {
//		return
//	}
//	pivot := partition(left, right, sli)
//	sort(left, pivot-1, sli)
//	sort(pivot+1, right, sli)
//}
//
//func partition(left, right int, sli []int) int {
//	pivot := sli[right]
//	i := left - 1
//	for j := left; j < right; j++ {
//		if sli[j] <= pivot {
//			i++
//			sli[i], sli[j] = sli[j], sli[i]
//		}
//	}
//	sli[i+1], sli[right] = sli[right], sli[i+1]
//	return i + 1
//}
