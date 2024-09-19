package main

import(
	"testing"
)


func TestMaxInt(t *testing.T){
	a, b := 2, 7

	res := maxInt(a, b)
	if res != b{
		t.Errorf("expected %d,got %d", b, res)
	}

}

func TestMain(m *testing.M){
	main()
}