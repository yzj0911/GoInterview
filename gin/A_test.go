package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	var getInferenceTaskRule = &GetInferenceTaskRule{1,1,1}
	a,_:=json.Marshal(getInferenceTaskRule)
	fmt.Println(string(a))

}
