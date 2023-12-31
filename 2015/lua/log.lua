local log = {}

local colors = {
	black = 30,
	red = 31,
	green = 32,
	yellow = 33,
	blue = 34,
	magenta = 35,
	cyan = 36,
	white = 37,
}

local function l(col, fmt, ...)
	io.write(string.format("\27[%dm", colors[col]))
	io.write(string.format(fmt, ...))
	io.write("\27[0m")
end

function log.warn(fmt, ...)
	l('yellow', fmt, ...)
end

function log.report(fmt, ...)
	l('green', fmt, ...)
end

function log.trace(fmt, ...)
	l('blue', fmt, ...)
end

function log.error(fmt, ...)
	l('red', fmt, ...)
end

function log.info(fmt, ...)
	l('cyan', fmt, ...)
end

function log.map(t)
	for k, v in pairs(t) do
		print(k, v)
	end
end

function log.list(t)
	for k, v in ipairs(t) do
		print(k, v)
	end
end

return log