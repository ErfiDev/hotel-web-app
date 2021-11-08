package main

import "testing"

func TestRun(t *testing.T){
	err := InitProject()
	if err != nil {
		t.Error("Failed InitProject()")
	}
}