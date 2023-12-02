BEGIN {
	FS = ":"

	bagContents["red"] = 12
	bagContents["green"] = 13
	bagContents["blue"] = 14
}

{
	rhs = $2
	FS = " "
	$0 = $1
	game = $2
	print game, rhs

	split(rhs, arr, ";")
	for (k in arr) {
		print "arr", arr[k]
		split(arr[k], subarr, ",")
		for (j in subarr) {
			print "subarr", subarr[j]
		}
	}
	FS = ":"
}
