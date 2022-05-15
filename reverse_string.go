package reverse_string

//package main

func ReverseString(input string) (output string) {
	// solution goes here

	result := []string{}
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, string(input[i]))
	}

	for _, j := range result {
		output = output + " " + j
	}
	return output
}
