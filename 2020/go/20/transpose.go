package main

// flipy: invert y
func flipy(in [][]string) [][]string {
	var out [][]string = make([][]string, len(in))
	for o, i := 0, len(in)-1; o < len(in); o, i = o+1, i-1 {
		out[o] = make([]string, len(in[i]))
		copy(out[o], in[i])
	}
	return out
}

// flipx: invert x
func flipx(in [][]string) [][]string {

	// adapted from canonical Go string reverse
	reverse := func(in []string) []string {
		var out []string = make([]string, len(in))
		copy(out, in)
		for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
		return out
	}

	n := len(in)
	var out [][]string = make([][]string, n)
	for i := 0; i < n; i++ {
		out[i] = reverse(in[i])
	}
	return out
}

// transpose: write columns of in as rows of out
// https://en.wikipedia.org/wiki/Transpose
func transpose(input [][]string) [][]string {
	if len(input) == 0 {
		return input
	}
	output := make([][]string, len(input[0]))
	for i := range output {
		output[i] = make([]string, len(input))
	}
	for i, row := range input {
		for j, val := range row {
			output[j][i] = val
		}
	}
	return output
}

func rotate(in [][]string) [][]string {
	out := transpose(in)
	return flipy(out)
}
