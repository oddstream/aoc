-- https://rosettacode.org/wiki/Priority_queue#Lua

PriorityQueue = {
    __index = {
        put = function(self, p, v)
            local q = self[p]
            if not q then
                q = {first = 1, last = 0}
                self[p] = q
            end
            q.last = q.last + 1
            q[q.last] = v
        end,
        pop = function(self)
            for p, q in pairs(self) do
                if q.first <= q.last then
                    local v = q[q.first]
                    q[q.first] = nil
                    q.first = q.first + 1
                    return p, v
                else
                    self[p] = nil
                end
            end
        end
    },
    __call = function(cls)
        return setmetatable({}, cls)
    end
}

setmetatable(PriorityQueue, PriorityQueue)

-- Usage:
local pq = PriorityQueue()

local tasks = {
	{10, 'Wash dog'},
	{10, 'Wash dog again'},
    {3, 'Clear drains'},
    {4, 'Feed cat'},
    {5, 'Make tea'},
    {1, 'Solve RC tasks'},
    {2, 'Tax return'}
}

print(#pq)

for _, task in ipairs(tasks) do
    print(string.format("Putting: %d - %s", table.unpack(task)))
    pq:put(table.unpack(task))
end

print(#pq)

for prio, task in pq.pop, pq do
    print(string.format("Popped: %d - %s", prio, task))
end

print(#pq)

-- local pq2 = PriorityQueue()
-- pq2:put(2, {x=2, y=2})
-- pq2:put(1, {x=1, y=1})
-- pq2:put(0, {x=0, y=0})
-- pq2:put(2, {x=2, y=2})
-- print(#pq2)
