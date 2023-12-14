package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func DumpData(data interface{}) {
	obj, _ := json.MarshalIndent(data, "", "\t")
	fmt.Print(string(obj))
}

func RemoveDuplicates[T string | int](sliceList []T) []T {

	allKeys := make(map[T]bool)
	list := []T{}

	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func StringInSlice(str string, slice []string) bool {

	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}

func IntInSlice(integer int, slice []int) bool {

	for _, v := range slice {
		if v == integer {
			return true
		}
	}

	return false
}

func EncodeMD5(input string) string {

	hash := md5.New()
	_, _ = hash.Write([]byte(input))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
