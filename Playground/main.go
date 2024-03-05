package main

// Recursive function for factorial
func factorial(n int) int {
	// limit condition
	if n == 0 {
		return 1
	}

	// recursive call
	return n * factorial(n-1)
}

/*
	n = 5
	
	1° Call: 5 * factorial(4)																5 * 24 = 120  ↑
				 |
				 2° Call: 4 * factorial(3)													4 * 6 = 24	  ↑
				 			 |
							 3° Call: 3 * factorial(2)										3 * 2 = 6	  ↑
							 			 |
										 4° Call: 2 * factorial(1)							2 * 1 = 2	  ↑
										 			  |
													  5° Call: 1 * factorial(0)				1 * 1 = 1	  ↑
													  			  |
																  6° Call: n == 0 -> return 1
*/

// ascii upware arrow: ↑