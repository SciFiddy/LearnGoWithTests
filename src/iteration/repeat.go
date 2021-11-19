//
//  11/18/2021
//
package iteration

func repeat(input string, count int) string {

	var output string
	for i := 0; i < count; i++ {
		output += input
	}
	return output
}
