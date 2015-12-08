package main

import (
  //"fmt"
)

type MyMathLib struct{
}

func (p *MyMathLib) Sum(maxNum int)(float64){
	sum:=0.0
	for i:=1;i<=maxNum;i++ {
		sum+=float64(i)
	}
	return sum
}
