local function gcd(a, b)
    return a == 0 and b or gcd(b % a, a)
end

local function lcm(a, b)
    return a * b / gcd(a, b)
end

local function lcm_array(arr)
    return #arr == 1 and arr[1] or lcm(arr[1], lcm_array{table.unpack(arr, 2)})
end

local arr = {2, 3, 4, 5, 6}
print(lcm_array(arr)) -- Output: 60
