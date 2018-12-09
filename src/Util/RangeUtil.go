
package RangeUtil

import "errors"

func XRange(i int) []int{
	var arr []int

	arr = make([]int,i)
	var k int
	for k=0;k<i;k++{
		arr[k] = k
	}

	return arr
}

func XRangeWithParam(start int, end int,step int)([]int,error){
	var arr []int

	count := end-start
	if end-start<0 {
		return nil,errors.New("End must great then start!")
	}

	arr = make([]int,count)
	var k int
	i := 0
	for k=start;k<end;k=k+step{
		arr[i] = k
		i++
	}

	if i+1<count {
		return arr[0:i],nil
	}

	return arr,nil	
}