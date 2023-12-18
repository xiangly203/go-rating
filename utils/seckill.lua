local voucherId = ARGV[1]
local userId = ARGV[2]

local stockKey = "seckill:stock:" .. voucherId
local orderKey = "seckill:order:" .. userId

-- Check stock
if (tonumber(redis.call('get', stockKey)) <= 0) then
    return 2
end

-- Check is order
if(redis.call("SISMEMBER", orderKey, userId) == 1) then
    return 3
end

redis.call("incrby", stockKey,-1)
redis.call("SADD", orderKey, userId)
