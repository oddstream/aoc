function string:startswith(start)
	return self:sub(1, #start) == start
end

function string:endswith(ending)
	return ending == "" or self:sub(-#ending) == ending
end
