-- https://adventofcode.com/2023/day/19 Aplenty

-- https://sankeymatic.com/build/

local log = require 'log'

local function split(str, pat)
	local t = {}  -- NOTE: use {n = 0} in Lua-5.0
	local fpat = "(.-)" .. pat
	local last_end = 1
	local s, e, cap = str:find(fpat, 1)
	while s do
	   if s ~= 1 or cap ~= "" then
		  table.insert(t, cap)
	   end
	   last_end = e+1
	   s, e, cap = str:find(fpat, last_end)
	end
	if last_end <= #str then
	   cap = str:sub(last_end)
	   table.insert(t, cap)
	end
	return t
end

-- MAYBE prune the rules:
-- 76(ish) rules always resolve to R
-- 68(ish) rules always resolve to A
-- 572 rules could be reduced to 428,
-- then this process could be reapplied
-- until no more reductions made?
-- 19-test.txt reduced from 11 to 8
-- 19-input.txt reduced from 572 to 481
-- conclusion: gains are too small to help with part 2

-- MAYBE find min,max for xmas variables in workflows
-- and use that to prune 1 .. 4000
-- 19-input.txt
-- x	138	3936
-- m	98	3854
-- a	11	3910
-- s	100	3971
-- again, gains too small to help with part 2

--[[
local function canPruneWorkflows(workflows)
	for wfn, rules in pairs(workflows) do
		if #rules == 0 then
			print(wfn, 'has no rules')
			return
		end
		local dst = rules[1].dst
		if dst == 'A' or dst =='R' then
			for i = 2, #rules do
				if rules[i].dst ~= dst then
					dst = ''
					break
				end
			end
			if dst ~= '' then
				return wfn, dst
			end
		end
	end
end

local function pruneWorkflows(workflows, wfn, dst)
	assert(dst~=wfn)
	assert(workflows[wfn])
	print('removing', wfn, 'replacing with', dst)
	workflows[wfn] = nil -- remove from a map
	for _, rules in pairs(workflows) do
		for _, rule in ipairs(rules) do
			if rule.dst == wfn then
				rule.dst = dst
				rule.var = nil
				rule.cmp = nil
				rule.num = nil
			end
		end
	end
end
]]

-- MAYBE turn each rule set into a Lua function
-- This basically turns qqz{s>2770:qs,m<1801:hdj,R}
-- into the function qqz_ = lambda: s>2770 and qs_() or m<1801 and hdj_() or R_()

---turn rules in a string to a list of parsed rules
---@param rules string like s>2770:qs,m<1801:hdj,R
---@return table[]
local function parseRules(rules)
	local out = {}
	local rulesList = split(rules, ',')
	for _, rule in ipairs(rulesList) do
		local var, cmp, num, dst
		dst = rule:match'^(%a+)$'	-- could be another rule or A or R
		if dst then
			out[#out+1] = {dst=dst}
		else
			var, cmp, num, dst = rule:match'(%l+)([%<%>])(%d+):(%a+)'
			if var and cmp and num and dst then
				num = tonumber(num)
				out[#out+1] = {var=var, cmp=cmp, num=num, dst=dst}
			else
				print('cannot parse', rule)
			end
		end
	end
	return out
end

---@param wf table[]
---@param vars table
---@return string
local function machine(wf, vars)
	for _, rule in ipairs(wf) do
		if not rule.cmp then -- var cmp num will be nil
			return rule.dst
		end
		if rule.cmp == '<' then
			if vars[rule.var] < rule.num then
				return rule.dst
			end
		elseif rule.cmp == '>' then
			if vars[rule.var] > rule.num then
				return rule.dst
			end
		end
	end
	return ''	-- shouldn't ever come here
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local workflows = {}

	local buildingRules = true
	-- local prunedWorkflows = false
	for line in io.lines(filename) do
		if #line == 0 then
			buildingRules = false
			goto continue
		end
		if buildingRules then
			-- px{a<2006:qkq,m>2090:A,rfg}
			local workflowName, rules
			workflowName, rules = line:match'(%a+){(.+)}'
			if workflowName and rules then
				workflows[workflowName] = parseRules(rules)
			else
				print('input error', line)
			end
		else
--[[
			if not prunedWorkflows then
				-- pruneWorkflows(workflows, 'gd', 'R')
				-- pruneWorkflows(workflows, 'lnx', 'A')
				-- pruneWorkflows(workflows, 'qs', 'A')

				local prunes = 0
				local wfn, dst = canPruneWorkflows(workflows)
				while wfn and dst do
					pruneWorkflows(workflows, wfn, dst)
					prunes = prunes + 1
					wfn, dst = canPruneWorkflows(workflows)
				end
				print(prunes, 'pruned')

				-- if wfn and dst then
					-- pruneWorkflows(workflows, wfn, dst)
				-- end

				prunedWorkflows = true
			end
]]
			-- {x=787,m=2655,a=1222,s=2876}
			local f, err = load('return ' .. line)	-- don't do this at home
			if err then print(err) end
			local t = f()
			local res = 'in'
			repeat
				res = machine(workflows[res], t)
			until res == 'A' or res == 'R'
			if res == 'A' then
				result = result + t.x + t.m + t.a + t.s
			end
		end
::continue::
	end

--[[
	do
		local xmin, xmax, mmin, mmax, amin, amax, smin, smax
		xmin, mmin, amin, smin = 1/0, 1/0, 1/0, 1/0
		xmax, mmax, amax, smax = 0, 0, 0, 0
		for _, rules in pairs(workflows) do
			for _, rule in ipairs(rules) do
				if rule.var == 'x' then
					if rule.num > xmax then xmax = rule.num end
					if rule.num < xmin then xmin = rule.num end
				end
				if rule.var == 'm' then
					if rule.num > mmax then mmax = rule.num end
					if rule.num < mmin then mmin = rule.num end
				end
				if rule.var == 'a' then
					if rule.num > amax then amax = rule.num end
					if rule.num < amin then amin = rule.num end
				end
				if rule.var == 's' then
					if rule.num > smax then smax = rule.num end
					if rule.num < smin then smin = rule.num end
				end
			end
		end
		print('x', xmin, xmax)
		print('m', mmin, mmax)
		print('a', amin, amax)
		print('s', smin, smax)
	end
]]
	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	local workflows = {}

	for line in io.lines(filename) do
		if #line == 0 then
			break
		end
		local workflowName, rules
		-- px{a<2006:qkq,m>2090:A,rfg}
		workflowName, rules = line:match'(%a+){(.+)}'
		if workflowName and rules then
			workflows[workflowName] = parseRules(rules)
		else
			print('input error', line)
		end
	end

	-- https://old.reddit.com/r/adventofcode/comments/18lwcw2/2023_day_19_an_equivalent_part_2_example_spoilers/

	for x = 138-1, 3936+1 do
		for m = 98-1, 3854+1 do
			for a = 11-1, 3910+1 do
				for s = 100-1, 3971+1 do
					local res = 'in'
					repeat
						res = machine(workflows[res], {['x']=x, ['m']=m, ['a']=a, ['s']=s})
					until res == 'A' or res == 'R'
					if res == 'A' then
						result = result + 1
					end
				end
			end
		end
	end

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test %d\n', partOne('19-test.txt', 19114))
-- log.report('part one      %d\n', partOne('19-input.txt', 449531))
-- log.report('part two test %d\n', partTwo('19-test.txt', 167409079868000))
-- log.report('part two      %d\n', partTwo('19-input.txt', 0))

