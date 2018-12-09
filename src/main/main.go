package main;

import (
	"fmt"
	"Util"
)



func main(){
	fmt.Println("Hello Go!")

	k := RangeUtil.XRange(10)

	for _, var1 := range k {
		fmt.Println(var1)
	}

	a1 := apple {address:"yantai",color:COLOR_RED,yy:28.6}

	a1.printInformation()

	a1.setColor(COLOR_GREEN)


	a1.printInformation()

	j,err := RangeUtil.XRangeWithParam(3,13,2)

	if err!=nil {
		fmt.Println(err)
		return 
	}

	for _, var2 := range j {
		fmt.Println(var2)
	}
}
