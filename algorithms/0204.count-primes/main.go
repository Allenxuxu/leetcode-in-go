package algorithm

func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrimes[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrimes[i] {
			for j := i * 2; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	var ret int
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			ret++
		}
	}

	return ret
}

func countPrimes1(n int) int {
	isPrimes := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrimes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrimes[i] {
			for j := i * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	var ret int
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			ret++
		}
	}

	return ret
}
