package test

import (
	"PicNext/utils"
	"fmt"
	"testing"
)

func TestRandName(t *testing.T) {
	s := utils.RandName()
	fmt.Println(s)
}
