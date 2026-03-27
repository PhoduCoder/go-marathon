package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

type DiffType string

const (
	Added    DiffType = "ADDED"
	Removed  DiffType = "REMOVED"
	Modified DiffType = "MODIFIED"
	TypeDiff DiffType = "TYPE_CHANGED"
)

type DiffEntry struct {
	Path     string
	Type     DiffType
	OldValue any
	NewValue any
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run main.go <old.json> <new.json>")
		os.Exit(1)
	}

	oldFile := os.Args[1]
	newFile := os.Args[2]

	oldData, err := os.ReadFile(oldFile)
	if err != nil {
		fmt.Printf("failed to read old file: %v\n", err)
		os.Exit(1)
	}

	newData, err := os.ReadFile(newFile)
	if err != nil {
		fmt.Printf("failed to read new file: %v\n", err)
		os.Exit(1)
	}

	var oldValue any
	if err := json.Unmarshal(oldData, &oldValue); err != nil {
		fmt.Printf("failed to parse old json: %v\n", err)
		os.Exit(1)
	}

	var newValue any
	if err := json.Unmarshal(newData, &newValue); err != nil {
		fmt.Printf("failed to parse new json: %v\n", err)
		os.Exit(1)
	}

	var diffs []DiffEntry
	diff("$", oldValue, newValue, &diffs)

	if len(diffs) == 0 {
		fmt.Println("No differences found.")
		return
	}

	for _, d := range diffs {
		switch d.Type {
		case Added:
			fmt.Printf("%-12s %s: %s\n", d.Type, d.Path, toJSON(d.NewValue))
		case Removed:
			fmt.Printf("%-12s %s: %s\n", d.Type, d.Path, toJSON(d.OldValue))
		case Modified, TypeDiff:
			fmt.Printf("%-12s %s: %s -> %s\n", d.Type, d.Path, toJSON(d.OldValue), toJSON(d.NewValue))
		}
	}
}

func diff(path string, oldVal, newVal any, diffs *[]DiffEntry) {
	// One side missing
	if oldVal == nil && newVal != nil {
		*diffs = append(*diffs, DiffEntry{
			Path:     path,
			Type:     Added,
			OldValue: nil,
			NewValue: newVal,
		})
		return
	}

	if oldVal != nil && newVal == nil {
		*diffs = append(*diffs, DiffEntry{
			Path:     path,
			Type:     Removed,
			OldValue: oldVal,
			NewValue: nil,
		})
		return
	}

	// Type check
	if reflect.TypeOf(oldVal) != reflect.TypeOf(newVal) {
		*diffs = append(*diffs, DiffEntry{
			Path:     path,
			Type:     TypeDiff,
			OldValue: oldVal,
			NewValue: newVal,
		})
		return
	}

	switch oldTyped := oldVal.(type) {
	case map[string]any:
		newTyped := newVal.(map[string]any)
		diffObjects(path, oldTyped, newTyped, diffs)
	case []any:
		newTyped := newVal.([]any)
		diffArrays(path, oldTyped, newTyped, diffs)
	default:
		if !reflect.DeepEqual(oldVal, newVal) {
			*diffs = append(*diffs, DiffEntry{
				Path:     path,
				Type:     Modified,
				OldValue: oldVal,
				NewValue: newVal,
			})
		}
	}
}

func diffObjects(path string, oldObj, newObj map[string]any, diffs *[]DiffEntry) {
	keysMap := make(map[string]struct{})

	for k := range oldObj {
		keysMap[k] = struct{}{}
	}
	for k := range newObj {
		keysMap[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		oldChild, oldExists := oldObj[key]
		newChild, newExists := newObj[key]

		childPath := path + "." + key

		switch {
		case !oldExists && newExists:
			*diffs = append(*diffs, DiffEntry{
				Path:     childPath,
				Type:     Added,
				OldValue: nil,
				NewValue: newChild,
			})
		case oldExists && !newExists:
			*diffs = append(*diffs, DiffEntry{
				Path:     childPath,
				Type:     Removed,
				OldValue: oldChild,
				NewValue: nil,
			})
		default:
			diff(childPath, oldChild, newChild, diffs)
		}
	}
}

func diffArrays(path string, oldArr, newArr []any, diffs *[]DiffEntry) {
	maxLen := len(oldArr)
	if len(newArr) > maxLen {
		maxLen = len(newArr)
	}

	for i := 0; i < maxLen; i++ {
		childPath := fmt.Sprintf("%s[%d]", path, i)

		switch {
		case i >= len(oldArr):
			*diffs = append(*diffs, DiffEntry{
				Path:     childPath,
				Type:     Added,
				OldValue: nil,
				NewValue: newArr[i],
			})
		case i >= len(newArr):
			*diffs = append(*diffs, DiffEntry{
				Path:     childPath,
				Type:     Removed,
				OldValue: oldArr[i],
				NewValue: nil,
			})
		default:
			diff(childPath, oldArr[i], newArr[i], diffs)
		}
	}
}

func toJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
