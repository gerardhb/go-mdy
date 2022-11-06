package mdy

import "testing"

func TestNew(t *testing.T) {

}

func NewMdy() *mdy {
	return New("384ed23690c2bee1", "NTViMmFjNGRhMDQ4NTQzNjhjNGFmZDQwMzRmM2VmZTM2YzYwYTVlMjcxMGE3YmVmMTZiZmU5ODk4NWRiMTQ2Nw==", "").WithDebug()
	//return New("2ba7c7dcc8c433ad", "YmZjOTNjMjhkNTJhOTI1ZmRmMzdiOWZkMDBkYzIyYmZlMWE4ODk4MWM5MDU1NWYwNWFiNDdjNjMyMzkyZTBjZg==", "").WithDebug()
}
