package Rotate

import (
	"fmt"
	"math"
)

func Pow(x int64, y int64) int64{
	return int64(math.Pow(float64(x), float64(y)))
}

func MaxRot(n int64) int64 {
	n2 := n
	var k int64 = 0

	for n2>=10{
		n2/=10
		k++
	}

	n2 = n
	var n3 int64 = 0
	var f int64 = 0
	var maxNum int64 = 0
	for ;k>0; {
		var st int64 = Pow(10, k)
		a := n2/st
		b := n2%st
		n3 = b*10 + a
		f = n3/st+ f*10
		n2 = n3%st
		if maxNum<f*st+n2{
			maxNum = f*st+n2
		}
		fmt.Printf("%d, ",f*Pow(10,k)+n2)
		k--

	}
	return maxNum
}

func main(){
	fmt.Print("\nMax = ",MaxRot(992495570))
}
