print(package.cpath)

local md5 = require 'md5'

-- https://en.wikipedia.org/wiki/MD5
-- MD5("The quick brown fox jumps over the lazy dog") = 9e107d9d372bb6826bd81d3542a419d6
local message = 'The quick brown fox jumps over the lazy dog'

print(message)
local md5_as_data  = md5.sum(message)       -- returns raw bytes
local md5_as_hex   = md5.sumhexa(message)   -- returns a hex string
print(md5_as_hex)
local md5_as_hex2  = md5.tohex(md5_as_data) -- returns the same string as md5_as_hex
print(md5_as_hex2)
if md5_as_hex == '9e107d9d372bb6826bd81d3542a419d6' then
	print('ok')
else
	print('not ok')
end