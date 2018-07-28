package lintcode

// BinarySearch suppose all elem stored in arr is positive number
func BinarySearch(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	start := 0
	end := len(arr) - 1

	for start + 1 < end {
		mid := (start + end) / 2
		if arr[mid] == target {
			end = mid;
		} else if arr[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if arr[start] == target {
		return start;
	}

	if arr[end] == target {
		return end;
	}

	return -1
}
