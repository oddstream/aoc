--[[
	local function bfs(start, goal)
		for row = 1, #galaxy do
			for col = 1, #galaxy[1] do
				galaxy[row][col].parent = nil
			end
		end
		if start.x == goal.x and start.y == goal.y then
			print('you are kidding, right?')
			return
		end
		local q = {start}
		start.parent = start
		while #q > 0 do
			local t = table.remove(q, 1)
			if t.x == goal.x and t.y == goal.y then
				return
			end
			for _, dir in ipairs({{x=1,y=0},{x=-1,y=0},{x=0,y=1},{x=0,y=-1}}) do
				local nx = t.x + dir.x
				local ny = t.y + dir.y
				if nx > 0 and ny > 0 and nx < #galaxy[1] and ny < #galaxy then
					local tn = galaxy[ny][nx]
					if tn.parent == nil then
						tn.parent = t
						q[#q+1] = tn
					end
				end
			end
		end
		assert(false, 'BFS not found')
	end

	local function bfscount(start, goal)
		local count = 0
		while goal ~= nil do
			if goal.x == start.x and goal.y == start.y then
				break
			end
			print(goal.y, goal.x)
			goal = goal.parent
			count = count + 1
		end
		return count
	end
]]
