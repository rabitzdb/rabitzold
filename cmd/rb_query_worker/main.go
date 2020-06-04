package main

import (
	"fmt"
)

func main() {
	// -8060873578586766283   -7012432064478608139
	value := int64(-7012432064478608139)

	sign := int64(signum(value))
	abs := value*sign
	fmt.Println(abs)
	if abs == 0 {
		fmt.Println("0")
	} else {
		indexValue := int64(1)
		for i := int64(0);i<64;i++ {
			if indexValue > abs {
				break
			}
			valid := abs & indexValue
			if valid == indexValue {
				fmt.Println(int64(indexValue))
			}
			indexValue = indexValue << 1
		}
		//bits := strconv.FormatUint(abs,2)
		//length := len(bits)-1
		//fmt.Println("Bits: ",bits)
		//for pos,charInt := range bits {
		//	if(charInt == 49){
		//		index := float64(length-pos)
		//		indexValue := int(math.Pow(2,index)) * sign
		//		fmt.Println(indexValue)
		//	}
		//}
	}

}
func signum(x int64) int {
	return int((x >> 63) | int64(uint64(-x)>>63))
}
