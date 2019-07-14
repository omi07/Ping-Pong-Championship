package main

import "testing"

func TestGetrandomnumber(t *testing.T) {
	var num int
	num = getrandomnumber()
	if num > 0 {
		t.Logf("Success Got valid random number")
	} else {
		t.Errorf("Failed to get random integer")
	}
	num = getrandomnumber()
	if num <= 10 {
		t.Logf("Success random number is less than or equal to 10")
	} else {
		t.Errorf("Failed random integer is greater than 10")
	}
}

func TestChecknumberexist(t *testing.T) {
	num := getrandomnumber()
	defarr := []int{2, 8, 7, 5, 1, 9, 10, num}
	exist := checknumberexist(num, defarr)
	if exist == 1 {
		t.Logf("Success number exist in array ")
	} else {
		t.Errorf("Failed Number exist but got 0")
	}
	defarr = []int{11, 14, 15, 0}
	exist = checknumberexist(num, defarr)
	if exist == 0 {
		t.Logf("Success number does not exist in array ")
	} else {
		t.Errorf("Failed Number not exist but got 1")
	}

}
