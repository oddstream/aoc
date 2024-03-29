---generate all permutations using Heap's algorithm
---@param arr any[]
---@return function iterator
function permutations(arr)
	local function swap(a, b)
		arr[a], arr[b] = arr[b], arr[a]
	end

	local function heap_permute(n)
		if n == 1 then
			coroutine.yield(arr)
		else
			for i = 1, n do
				heap_permute(n - 1)
				if n % 2 == 0 then
					swap(i, n)
				else
					swap(1, n)
				end
			end
		end
	end

	return coroutine.wrap(function()
		heap_permute(#arr)
	end)
end
