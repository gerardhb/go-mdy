package mdy

import "testing"

func TestNew(t *testing.T) {

}

func NewMdy() *mdy {
	return New("", "", "").WithDebug()
	//return New("", "", "").WithDebug()
}
