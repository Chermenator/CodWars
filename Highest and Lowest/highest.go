package highest

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	max, min := Slise(in)

	return fmt.Sprintf("%d %d", max, min)
}

func Slise(inp string) (int32, int32) {
	inp = strings.TrimSpace(inp)

	var inpSliese []string = strings.Split(inp, " ")

	max_num := int32(math.MinInt32)
	min_num := int32(math.MaxInt32)

	//var numbers []int32
	for i := 0; i < len(inpSliese); i++ {
		num, err := strconv.Atoi(inpSliese[i])
		if err != nil {
			log.Fatal(err)
			continue
		}

		//numbers[i] = int32(num)
		if int32(num) >= max_num {
			max_num = int32(num)
		}
		if int32(num) < min_num {
			min_num = int32(num)
		}
	}
	return max_num, min_num
}
