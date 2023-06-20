package sort

func Quick(sli []int) {
	sort(0, len(sli)-1, sli)
}

func sort(l, r int, sli []int) {
	if l > r || len(sli[l:r]) < 2 {
		return
	}
	pivot := sli[l]
	left := l + 1
	//fmt.Println(pivot)
	//fmt.Println("start", sli)
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
		//fmt.Println("for : ", i, sli)
		//l++
		//r--
	}

	sli[l], sli[left-1] = sli[left-1], sli[l]
	sort(l, left-2, sli)
	sort(left, r, sli)

}
