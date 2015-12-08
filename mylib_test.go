package main

import (
//	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_01HelloWord(t *testing.T) {
	assert.Equal(t, 2,1+1)
}

func Test_02Sum100OK(t *testing.T) {
	lib:=new(MyMathLib)
	assert.Equal(t, 5050.0,lib.Sum(100))
}

