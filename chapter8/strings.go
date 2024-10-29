package main

import (
	"fmt"
	"strings"
)

func main() {
	//To search for a smaller string in a bigger string, use the Contains function:
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "es"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "s"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToUpper("aaaaaa"))
	fmt.Println(strings.ToLower("AAAAA"))

	//Sometimes we need to work with strings as binary data. To convert a string to a slice
	//of bytes (and vice versa), do this:
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr)
	fmt.Println(str)

}
