package sort

import(
	"fmt"
)

func Sort(arr []int){
	fmt.Println("before sort:%a", arr)
	for i:=1; i < len(arr); i++ {
		var ta = arr[0:i+1];
		for j,_ := range ta{
			if ta[j] <= arr[i] {
				var tv = arr[i]
				copy(arr[j+1:i+1], arr[j:i])
				arr[j]=tv
				break
			}
		}
	}
	fmt.Println("end sort:%a", arr)
}

