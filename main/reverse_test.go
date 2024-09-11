package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F)  {
	test_cases := []string{"Bharath","Mani!","Vasu123"}
	for _,c := range test_cases {
		f.Add(c)
	}

	f.Fuzz(func(t *testing.T, orig string){
		rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }

        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
	})
	
}