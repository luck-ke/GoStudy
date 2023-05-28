package algorithnm

func Quick(sli []int) []int {
	if len(sli) < 1 {
		return sli
	}
	r := 0
	l := len(sli) - 1
	for i := 0; i < l; i++ {
		if r > l {
			break
		}
		if sli[r] > sli[l] {
			sli[r], sli[l] = sli[l], sli[r]
		}
		r++
		l--
	}
	Quick(sli[:r])
	Quick(sli[l:])
	return sli
}
