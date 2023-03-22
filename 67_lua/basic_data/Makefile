






wrk:
	wrk -t8 -c400 -d 10s -T2s --latency -s new.lua  http://127.0.0.1:4900/v3/c/tts/getAllTTS


wrk2:
	wrk -t4 -c100 -d 30s  --latency -s new.lua  http://127.0.0.1:4900/v3/c/tts/getAllTTS


# -t 8: 8 threads
# -c 400: 400 connections
# -d 10s: 10 seconds
# -T 2s: 2 seconds timeout


wrk3:
	wrk -t4 -c100 -d 30s  --latency -s new.lua  http://127.0.0.1:4900/tts/query


wrk4:
	wrk -t1 -c1 -d 1s  --latency -s realInfo.lua  http://127.0.0.1:4900/v3/u/vehicle/realTimeInfo

cpu:
	go tool pprof -http=:32001 http://127.0.0.1:4900/debug/pprof/profile?seconds=10
