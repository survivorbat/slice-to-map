package slicetomap

import "fmt"

type TestObject struct {
	ID   int
	Name string
}

func ExampleToMap() {
	testObjects := []TestObject{
		{ID: 1, Name: "Test 1"},
		{ID: 2, Name: "Test 2"},
		{ID: 3, Name: "Test 3"},
	}

	testObjectsMap := ToMap(testObjects, func(obj TestObject) int {
		return obj.ID
	})

	fmt.Printf("%+v", testObjectsMap)
}
