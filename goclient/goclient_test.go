package main

import (
	"fmt"
	"testing"
)

func TestPrintAnimal(t *testing.T) {
	tc := []struct{
		name string
		animal string
		wantErr bool
	}{
		{
			name: "fail_noanimal",
			animal: "",
			wantErr: true,
		}, {
			name: "pass_cat",
			animal: "Cat",
		},	
	}

	for _, test := range tc {
		res, err := printAnimal(test.animal)
		if err != nil && !test.wantErr {
			t.Errorf("error not expected: %v", err)
		} else {
			fmt.Printf("Result for %s is: %s\n", test.name, res)
		}
	}
}
