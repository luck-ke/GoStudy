package algorithnm

//	func Quick(sli []int) {
//		r := 0
//		l := len(sli) - 1
//		sort(r, l, sli)
//	}
//
//	func sort(r, l int, sli []int) {
//		if len(sli) < 2 {
//			return
//		}
//		right := r
//		left := l
//		for i := 0; i < l; i++ {
//			if r > l {
//				break
//			}
//			if sli[r] > sli[l] {
//				sli[r], sli[l] = sli[l], sli[r]
//			}
//			r++
//			l--
//		}
//		sort(right, r, sli)
//		sort(l, left, sli)
//
// }
func Quick(sli []int) {
	sort(0, len(sli)-1, sli)
}

func sort(left, right int, sli []int) {
	if left >= right {
		return
	}
	pivot := partition(left, right, sli)
	sort(left, pivot-1, sli)
	sort(pivot+1, right, sli)
}

func partition(left, right int, sli []int) int {
	pivot := sli[right]
	i := left - 1
	for j := left; j < right; j++ {
		if sli[j] <= pivot {
			i++
			sli[i], sli[j] = sli[j], sli[i]
		}
	}
	sli[i+1], sli[right] = sli[right], sli[i+1]
	return i + 1
}
