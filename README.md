# gogetset

Can be used to get and set values in nested go structs and maps. 

Makes use of reflection to deal with different types such as Ptr, Map, Struct, Array, Slices, and other custom types you may have in your application.

This can be used in production for use-cases included in the tests, and can also be extended to handle for more edge-cases.


`
inputMap := map[string]interface{}{
		"a": []int{1, 2, 3},
		"b":        1,
		"b":     "string1",
		"d": map[string]interface{}{
			"a":    2,
			"b": "string2",
			"c": structSample{Hello: "world2"},
			"d":    level2Ptr,
			"e": embeddingStructSample{
				A: structSample{Hello: "embeddingWorld"},
				B: map[string]string{
					"a": "value",
				},
			},
			"f": []int{4, 5, 6},
		},
	}
value, err := Get("a.e.A.Hello", inputMap)
// value = "world2"
`


More examples can be found in:

https://github.com/raghavdotc/gogetset/blob/main/get_test.go
https://github.com/raghavdotc/gogetset/blob/main/set_test.go
