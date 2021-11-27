package main

import "testing"

func TestRun(t *testing.T){
	_ , err := InitProject()
	if err != nil {
		t.Error("Failed InitProject()")
	}
}