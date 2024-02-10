package main

func topEdge(in [][]string) string {
	var out string
	for x := 0; x < len(in[0]); x++ {
		out += in[0][x]
	}
	return out
}

func bottomEdge(in [][]string) string {
	var out string
	for x := 0; x < len(in[0]); x++ {
		out += in[len(in)-1][x]
	}
	return out
}

func leftEdge(in [][]string) string {
	var out string
	for y := 0; y < len(in); y++ {
		out += in[y][0]
	}
	return out
}

func rightEdge(in [][]string) string {
	var out string
	for y := 0; y < len(in); y++ {
		out += in[y][len(in[y])-1]
	}
	return out
}
