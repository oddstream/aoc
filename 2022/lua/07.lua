local inputFilename = "07-input.txt"
-- {name, type, size, children, parent}
--
-- $ cd /
-- $ cd ..
-- $ cd <name>
-- $ ls
-- dir <name>
-- <bytes> <name>

local root = {name='root', type='directory', size=0, children={}, parent=nil}

local pwd = root

for line in io.lines(inputFilename) do
	if line:sub(1,1) == '$' then
		if line == '$ cd /' then
			-- print('change to root')
			pwd = root
			-- print('pwd now', pwd.name)
		elseif line == '$ cd ..' then
			-- print('pwd change to parent')
			if pwd.parent then
				pwd = pwd.parent
			else
				error('no parent')
			end
			if not pwd then
				error('no pwd')
			end
			-- print('pwd now', pwd.name)
		elseif line:match("$ cd %a*") then
			local subdir = line:match("$ cd (%a*)")
			-- io.write('pwd change to ')
			for _, child in pairs(pwd.children) do
				if child.name == subdir and child.type == 'directory' then
					-- print(subdir)
					pwd = child
					break
				end
			end
			-- print('pwd now', pwd.name)
		elseif line == '$ ls' then
			-- print('-----------------')
		end
	else
		if line:match("dir %a+") then
			local name = line:match("dir (%a+)")
			-- print('insert directory', name)
			for _, entry in pairs(pwd.children) do
				if entry.name == name and entry.type == 'directory' then
					print('ERROR already a directory called', name)
					break
				end
			end
			table.insert(pwd.children, {name=name, type='directory', size=0, children={}, parent=pwd})
		else
			local bytes, name = line:match("(%d*) (.*)")
			bytes = tonumber(bytes)
			-- print('insert file', name, bytes)
			table.insert(pwd.children, {name=name, type='file', size=bytes, parent=pwd})
		end
	end
end

local function lsall(dir, indent)
	if not dir then
		print('ERROR no dir')
		return
	end
	if not dir.name or dir.name == "" then
		print('ERROR nil dir')
		return
	end
	if not dir.children then
		print('ERROR', dir.name, 'nil children')
		return
	end
	for _, entry in pairs(dir.children) do
		if entry.type == 'file' then
			print(indent, 'f', entry.name, entry.size)
		elseif entry.type == 'directory' then
			print(indent, 'd', entry.name)
			if entry.children and #entry.children > 0 then
				print(#entry.children)
				lsall(entry, indent .. "-")
			end
		end
	end
end

-- lsall(root, "-")

local comp_total = 0
local sizes = {}

local function dir_size(dir)
	local size = 0
	for _, entry in pairs(dir.children) do
		if entry.type == 'file' then
			size = size + entry.size
		elseif entry.type == 'directory' then
			size = size + dir_size(entry)
		end
	end
	table.insert(sizes, size)
	print(dir.name, size)
	if size < 100000 then
		comp_total = comp_total + size
	end
	return size
end

dir_size(root)
-- e 584
-- a 94853
-- d 24933642
-- / 48381165
print('part 1', comp_total) -- 95437, 1423358

local capacity = 70000000
local targetUnusedSpace = 30000000
local usedSpace = 40532950 -- size of root
local unusedSpace = capacity - usedSpace
local bestSize = math.huge

for _, size in pairs(sizes) do
	if unusedSpace + size >= targetUnusedSpace then
		bestSize = math.min(bestSize, size)
	end
end

print('part 2', bestSize)	-- 545729

-- table.sort(sizes, function(a, b) return a < b end)
-- for k, v in pairs(sizes) do
-- 	print(k, v)
-- end