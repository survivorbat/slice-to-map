# 📍 Slice to Map

[![Go package](https://github.com/survivorbat/slicetomap/actions/workflows/test.yaml/badge.svg)](https://github.com/survivorbat/gorm-deep-filtering/actions/workflows/test.yaml)

I needed this functionality without having to download a big utility package.

## ⬇️ Installation

`go get github.com/survivorbat/slice-to-map`

## 📋 Usage

```go
package main

import (
	"fmt"
	"github.com/survivorbat/slicetomap"
)

type TestObject struct {
	ID   int
	Name string
}

func main() {
	testObjects := []TestObject{
		{ID: 1, Name: "Test 1"},
		{ID: 2, Name: "Test 2"},
		{ID: 3, Name: "Test 3"},
	}

	testObjectsMap := slicetomap.ToMap(testObjects, func(obj TestObject) int {
		return obj.ID
	})

	fmt.Printf("%+v", testObjectsMap) // Test 1
}

```

## 🔭 Plans

Nothing here yet.
