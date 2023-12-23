local voucherID = ARGV[1]
local userID = ARGV[2]
local orderID = ARGV[3]
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

redis.call("xadd","stream.orders","*", "voucherID", voucherID, "userID", userID, "orderID", orderID)

return 1
