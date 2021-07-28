package math

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'

type addTest struct {
     arg1, arg2, expected int
}

var addTests = []addTest{
               addTest{2, 3, 5},
               addTest{4, 8, 12},
               addTest{9, 9, 18},
               addTest{1, 3, 4},
             }


func TestAddo(t *testing.T) {
     for _, test := range addTests {
          if output := Add(test.arg1, test.arg2); output != test.expected {
             t.Errorf("Output %q not equal to expected %q", output, test.expected)
          }
     }
}

               
