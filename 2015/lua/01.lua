local log = require 'log'

local f = assert(io.open("01-input.txt", "r"))
local content = f:read("*all")
f:close()

local level = 0
for i = 1, #content do
	local c = content:sub(i,i)
	if c == '(' then
		level = level + 1
	elseif c == ')' then
		level = level - 1
	else
		log.error('unexpected character \'%s\' at position %d\n', c, i)
	end
	if level < 0 then
		log.report('basement at position %d\n', i)	-- 1771
		break
	end
end

log.report('level %d\n', level) -- 138
