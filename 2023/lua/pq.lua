-- pq very basic priority queue

local pq = {}
pq.__index = pq

function pq.new()
	return setmetatable({q={}}, pq)
end

function pq:len()
	return #self.q
end

function pq:add(p, v)
	if type(p) == 'table' then
		table.insert(self.q, p)
	elseif type(p) == 'number' or type(p) == 'string' then
		table.insert(self.q, {p=p, v=v})
	else
		error('pq:add unknown type of argument')
	end
	table.sort(self.q, function(a, b) return a.p < b.p end)
end

function pq:pop()
	if #self.q == 0 then return end
	local p, v = self.q[1].p, self.q[1].v
	table.remove(self.q, 1)
	return p, v
end

return pq
