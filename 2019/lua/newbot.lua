---@comments https://github.com/mnml/aoc/blob/main/2019/15/2.go
---@comments TODO needs mem[+1] adjustment; ip is 0-based, mem is 1-based
---@param mem integer[]
---@param infn function
---@param outfn function
local function run(mem, infn, outfn)
	local arity = {4,4,2,2,3,3,4,4,2}
	local ip, rb = 1,0
	while true do
		local ins = string.format('%05d', mem[ip])
		local op = tonumber(ins:sub(4,5))
		local par = function(i)
			local mode
			if i == 1 then
				mode = ins:sub(3,3)
			elseif i == 2 then
				mode = ins:sub(2,2)
			elseif i == 3 then
				mode = ins:sub(1,1)
			else
				log.error('unknown parameter %d\n', i)
			end
			if mode == '0' then		-- POSITION
				assert(mem[ip+i]~=nil)
				assert(mem[mem[ip+i]]~=nil)
				return mem[ip+i]
			elseif mode == '1' then	-- IMMEDIATE
				return ip + i
			elseif mode == '2' then	-- RELATIVE
				assert(mem[ip+i]~=nil)
				assert(mem[mem[ip+i]]~=nil)
				return rb + mem[ip+i]
			else
				log.error('unknown mode %s\n', mode)
			end
		end
		if op == 1 then
			mem[par(3)+1] = mem[par(1)+1] + mem[par(2)+1]
		elseif op == 2 then
			mem[par(3)+1] = mem[par(1)+1] * mem[par(2)+1]
		elseif op == 3 then
			mem[par(1)+1] = infn()
		elseif op == 4 then
			outfn(mem[par(1)+1])
		elseif op == 5 then
			if mem[par(1)+1] ~= 0 then
				ip = mem[par(2)+1]
				goto continue
			end
		elseif op == 6 then
			if mem[par(1)+1] == 0 then
				ip = mem[par(2)+1]
				goto continue
			end
		elseif op == 7 then
			if mem[par(1)+1] < mem[par(2)+1] then
				mem[par(3)+1] = 1
			else
				mem[par(3)+1] = 0
			end
		elseif op == 8 then
			if mem[par(1)+1] == mem[par(2)+1] then
				mem[par(3)+1] = 1
			else
				mem[par(3)+1] = 0
			end
		elseif op == 9 then
			rb = rb + mem[par(1)+1]
		elseif op == 99 then
			return
		else
			log.error('unknown op %d\n', op)
			break
		end
		ip = ip + arity[op]
::continue::
	end
end

