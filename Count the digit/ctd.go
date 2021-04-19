package CountTheDigit

import "fmt"

func NbDig(n int, d int) int{
	fmt.Print("the k*k that contain the digit 1 are: ")
	var count int
	for k := 0; k < n; k++ {
		s  := k*k

		for s/10!=0{
			if(s%10==d){
				count++
				fmt.Printf( " %d", k*k)
			}
			s = s/10

		}

	}
	if d==0 { count++ }
	fmt.Printf("So there are %d digits 1 for the squares of numbers between 0 and %d.", count, n)
	return d
}

