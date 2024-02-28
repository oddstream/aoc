-- https://adventofcode.com/2023/day/19 Aplenty

-- https://sankeymatic.com/build/

local log = require 'log'

---split library function
---@param str string eg "a<2006:qkg,m>2090:A,rfg"
---@param pat string
---@return table eg the above would be split into three strings
local function split(str, pat)
	local t = {}  -- NOTE: use {n = 0} in Lua-5.0
	local fpat = "(.-)" .. pat
	local last_end = 1
	local s, e, capture = str:find(fpat, 1)
	while s do
	   if s ~= 1 or capture ~= "" then
		  table.insert(t, capture)
	   end
	   last_end = e+1
	   s, e, capture = str:find(fpat, last_end)
	end
	if last_end <= #str then
	   capture = str:sub(last_end)
	   table.insert(t, capture)
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

---turn rules in a string to a list of parsed rules
---what we COULD do is turn the rules into an on-the-fly
---Lua function
---@param rules string like s>2770:qs,m<1801:hdj,R
---@return table[] each containing var, cmp, num, dst
local function parseRules(rules)
	local out = {}
	for _, rule in ipairs(split(rules, ',')) do
		local dst = rule:match'^(%a+)$'	-- could be another rule or A or R
		if dst then
			out[#out+1] = {dst=dst}
		else
			local var, cmp, num
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
	-- a rule has .cmp .dst .num .var
	-- vars are .x .m .a .s
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
			-- key{filter[], default}
			-- where filter is char op number : key
			local workflowName, rules
			workflowName, rules = line:match'(%a+){(.+)}'
			if workflowName and rules then
				workflows[workflowName] = parseRules(rules)
			else
				print('input error', line)
			end
		else
			-- turn a string like {x=787,m=2655,a=1222,s=2876}
			-- into a Lua table the quick and dirty way
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

--[[
	test1
x	1416	2662 (1246)
m	838		2090 (1252)
a	1716	3333 (1617)
s	537		3448 (2911)

	input
x	138		3936
m	98		3854
a	11		3910
s	100		3971
]]

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

---@param rule table
local function printRule(rule)
	print(rule.var, rule.cmp, rule.num, rule.dst)
end

---utility function to deep copy a table
---@param obj table
---@return table
local function copy(obj)
	if type(obj) ~= 'table' then return obj end
	local res = {}
	for k, v in pairs(obj) do res[copy(k)] = copy(v) end
	return res
end

---@param workflows table[] of workflows, each has has .cmp .dst .num .var
---@param current string node name; starts with "in", ends with "A" or "R"
---@param varRanges table[] x m a s .left .right
local function runWorkflowRanges(workflows, current, varRanges)
	if current == 'A' then
		local prod = 1
		for _, v in pairs(varRanges) do
			prod = prod * (v.right - v.left)
		end
		return prod
	elseif current == 'R' then
		return 0
	end
	local sum = 0
	for _, rule in ipairs(workflows[current]) do
		-- printRule(rule)
		if not rule.cmp then
			sum = sum + runWorkflowRanges(workflows, rule.dst, varRanges)
			return sum
		else
			local v = varRanges[rule.var]
			if rule.cmp == '<' then
				if v.right <= rule.num then
					sum = sum + runWorkflowRanges(workflows, rule.dst, varRanges)
					return sum
				elseif v.left >= rule.num then
				else
					local childVarRanges = copy(varRanges)
					childVarRanges[rule.var] = {left=v.left, right=rule.num}
					sum = sum + runWorkflowRanges(workflows, rule.dst, childVarRanges)
					if v.right == rule.num then
						return sum
					end
					childVarRanges[rule.var] = {left=rule.num, right = v.right}
					varRanges = childVarRanges
				end
			elseif rule.cmp == '>' then
				if v.left > rule.num then
					sum = sum + runWorkflowRanges(workflows, rule.dst, varRanges)
					return sum
				elseif v.right <= rule.num + 1 then
				else
					local childVarRanges = copy(varRanges)
					childVarRanges[rule.var] = {left=rule.num + 1, right=v.right}
					sum = sum + runWorkflowRanges(workflows, rule.dst, childVarRanges)
					if v.left == rule.num + 1 then
						return sum
					end
					childVarRanges[rule.var] = {left=v.left, right=rule.num+1}
					varRanges = childVarRanges
				end
			end
		end
	end
	return sum
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
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

	local varRanges = {
		x = {left=1, right=4001},
		m = {left=1, right=4001},
		a = {left=1, right=4001},
		s = {left=1, right=4001},
	}
	local result = runWorkflowRanges(workflows, "in", varRanges)
	if expected and result ~= expected then
		log.error('expected %d, got %d (%d)\n', expected, result, expected - result)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test %d\n', partOne('19-test.txt', 19114))
log.report('part one      %d\n', partOne('19-input.txt', 449531))
-- log.report('part two test %d\n', partTwo('19-test.txt', 167409079868000))
log.report('part two      %d\n', partTwo('19-input.txt', 122756210763577))

-- 4000*4000*4000*4000 = 256000000000000, 15 digits

--[[
$ time luajit 19.lua
Lua 5.1
part one      449531
part two      122756210763577

real	0m0.043s
user	0m0.031s
sys	0m0.004s
]]