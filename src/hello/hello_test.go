package hello

import "testing"
import "strings"

func TestStrReverse(t *testing.T) {  
    reversedString := strReverse("Hello World")
    expectedString := "dlroW olleH"
    if strings.Compare(reversedString,expectedString)!=0  {
       t.Errorf("String Value Incorrect, Actual: %s, Expected: %s.", reversedString, expectedString)
    }
}
