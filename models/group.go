package models

type Group struct {
	items map[int]interface{} // Map of Item or Group. Its key is the ordering index in presentation.
}
