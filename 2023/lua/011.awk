#!/bin/awk -f

# https://github.com/patsie75/aoc/blob/main/2023/day01a.awk

{
  gsub(/[^0-9]/, "")		# remove any non-numeric character
  n = split($0, arr, "")	# split digits to array
  val = arr[1] "" arr[n]	# get first and last digit
  sum += val			# add to sum
}

END {
  print sum
}