package examples

import (
	"testing"
)

func TestGet(t *testing.T) {
	//initialization

	//execution
	endpoints, err := GetEndPoints()
	if err !=nil {
		t.Error(err)
	}
	if endpoints == nil {
		t.Error("algum erro")
	}
	//validation
}
