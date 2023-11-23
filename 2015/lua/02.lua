local log = require 'log'

local totalPaper, totalRibbon = 0, 0
for line in io.lines('02-input.txt') do
	local w, h, l = line:match('(%d+)x(%d+)x(%d+)')
	local list = {tonumber(w), tonumber(h), tonumber(l)}
	table.sort(list, function(a,b) return a < b end)
	-- log.trace('%d\t%d\t%d\n', list[1], list[2], list[3])
	local area = 2*l*w + 2*w*h + 2*h*l
	local extra = list[1] * list[2]
	totalPaper = totalPaper + area + extra
	local ribbon = list[1] + list[1] + list[2] + list[2]
	local bow = w * h * l
	totalRibbon = totalRibbon + ribbon + bow
end

log.report('paper %d\n', totalPaper) -- 1606483
log.report('ribbon %d\n', totalRibbon) -- 3842356

