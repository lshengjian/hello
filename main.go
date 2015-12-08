package main
import (
	"fmt"
	"math"
//	"os"
//	"strconv"
//	"github.com/lshengjian/mymath"
)
func IsPrime(x int) bool {
        for i := 2; i<=int(math.Sqrt(float64(x))); i++ {
                if x%i==0 {
                        return false
                }
        }
        return true
}
func RemoveItem(data []int,idx int) []int {
	return append(data[:idx], data[idx+1:]...)
}
/*func Filter(s []int, fn func(int) bool) []int {
    var p []int 
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}*/
func Filter(s []int, idx int) []int {
    p := s[0:idx+1]
    d := s[idx]  
    for i:=(idx+1);i<len(s);i++ {
        if s[i]%d != 0 {
            p = append(p, s[i])
        }
    }
    return p
}
func main() {
	//argsWithoutProg := os.Args[1:]
	TOTAL:= 10//strconv.Atoi(argsWithoutProg[0])
	data:= make([]int,TOTAL)
	for i:=0 ;i<TOTAL;i++ {
		data[i]=(i+2)
	}
	max := int( math.Sqrt( float64 ( data[len(data)-1] ) ) )+1
	fmt.Println(data,max)
	fmt.Println("=============")
	idx := 0
	for {
		if data[idx] > max{
            break
		}
		data = Filter(data , idx)
		fmt.Println(data)
		fmt.Println("------------------------")
		idx++
	}
}
