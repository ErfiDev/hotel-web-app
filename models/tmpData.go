package models

type TmpData struct{
	StringMap map[string]string
	IntMap map[string]int
	Float64Map map[string]float64
	CSRF string
	Data map[string]string
	Error string
	Warning string
	Flash string
}