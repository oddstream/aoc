-- u/waffle3z
local examples, program = {}, {}
local before, data;
for v in getinput():gmatch("[^\n]+") do
	local numbers, index = {}, 0
	for n in v:gmatch("%d+") do
		numbers[index] = tonumber(n)
		index = index + 1
	end
	if v:match("Before") then
		before = numbers
	elseif v:match("After") then
		examples[#examples+1] = {before = before, after = numbers, data = data}
		before, data = nil
	elseif before then
		data = numbers
	else
		program[#program+1] = numbers
	end
end

function binary(n)
	local s = ""
	repeat
		s = (n%2)..s
		n = (n - n%2)/2
	until n == 0
	return ("0"):rep(16-#s)..s
end

function AND(a, b)
	a, b = binary(a), binary(b)
	local s = ""
	for i = 1, #a do
		s = s..((a:sub(i, i) == "1" and b:sub(i, i) == "1") and "1" or "0")
	end
	return tonumber(s, 2)	
end

function OR(a, b)
	a, b = binary(a), binary(b)
	local s = ""
	for i = 1, #a do
		s = s..((a:sub(i, i) == "1" or b:sub(i, i) == "1") and "1" or "0")
	end
	return tonumber(s, 2)	
end

local r;
local operations = {
	addr = function(a, b) return r[a] + r[b] end;
	addi = function(a, b) return r[a] + b end;
	mulr = function(a, b) return r[a] * r[b] end;
	muli = function(a, b) return r[a] * b end;
	banr = function(a, b) return AND(r[a], r[b]) end;
	bani = function(a, b) return AND(r[a], b) end;
	borr = function(a, b) return OR(r[a], r[b]) end;
	bori = function(a, b) return OR(r[a], b) end;
	setr = function(a, b) return r[a] end;
	seti = function(a, b) return a end;
	gtir = function(a, b) return a > r[b] and 1 or 0 end;
	gtri = function(a, b) return r[a] > b and 1 or 0 end;
	gtrr = function(a, b) return r[a] > r[b] and 1 or 0 end;
	eqir = function(a, b) return a == r[b] and 1 or 0 end;
	eqri = function(a, b) return r[a] == b and 1 or 0 end;
	eqrr = function(a, b) return r[a] == r[b] and 1 or 0 end;
}
local possible = {}
for _, f in pairs(operations) do
	local t = {}
	for i = 0, 15 do t[i] = true end
	possible[f] = t
end

local count = 0
for _, e in pairs(examples) do
	local valid = 0
	local n, a, b, c = unpack(e.data, 0, 3)
	r = e.before
	for _, f in pairs(operations) do
		if f(a, b) == e.after[c] then
			valid = valid + 1
		else
			possible[f][n] = false
		end
	end
	if valid >= 3 then count = count + 1 end
end
print(count)

local opcode, list = {}, {}
for f, t in pairs(possible) do list[#list+1] = f end
for i = 1, #list do
	table.sort(list, function(a, b)
		local c1, c2 = 0, 0
		for k, v in pairs(possible[a]) do if v then c1 = c1 + 1 end end
		for k, v in pairs(possible[b]) do if v then c2 = c2 + 1 end end
		return c1 < c2
	end)
	local f = table.remove(list, 1)
	for k, v in pairs(possible[f]) do
		if v then
			opcode[k] = f
			for _, y in pairs(possible) do
				y[k] = false
			end
			break
		end
	end
end

r = {[0] = 0, 0, 0, 0}
for _, line in pairs(program) do
	local n, a, b, c = unpack(line, 0, 3)
	r[c] = opcode[n](a, b)
end
print(r[0])
