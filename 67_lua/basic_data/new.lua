wrk.method = "POST"
wrk.body = '{"token":"SI9VoL5sNRv98JRUnkJa","wid":"xiaojing","page":1,"limit":1}'
wrk.headers["Content-Type"] = "application/json"
--function request()
--    return wrk.format('POST', nil, nil, body)
--end
response = function(status, headers, body)
    print(body) --调试用，正式测试时需要关闭，因为解析response非常消耗资源
end
-- table 索引可以是数字，也可以是字符串

-- function wrk.format(method, path, headers, body)

