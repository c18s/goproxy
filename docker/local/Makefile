build:
	docker build -t local/goproxy:latest -f Dockerfile ./../..

http:
	docker run -p 33080:33080 local/goproxy:latest http -t tcp -p ":33080" -a "user1:pass1"
