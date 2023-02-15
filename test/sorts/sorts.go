package sorts

func BubbleSort(listPtr *[]int) bool {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(*listPtr)-1; i++ {
			if (*listPtr)[i] > (*listPtr)[i+1] {
				swap(&(*listPtr)[i], &(*listPtr)[i+1])
				swapped = true
			}
		}
	}
	return true
}

func SelectionSort(listPtr *[]int) bool {
	var minPtr *int
	for i := 0; i < len(*listPtr); i++ {
		minPtr = &(*listPtr)[i]
		for j := i; j < len(*listPtr); j++ {
			if (*listPtr)[j] < *minPtr {
				minPtr = &(*listPtr)[j]
			}
		}
		swap(minPtr, &(*listPtr)[i])
	}

	return true
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
