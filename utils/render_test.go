package utils

import "testing"

func TestTemplateCache(t *testing.T) {
	_ , err := CreateTemplateCache()
	if err != nil {
		t.Error("test failed!")
	}
}

func TestRenderTmp(t *testing.T){

}