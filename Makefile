export GO= O111MODULE=on GOSUMDB=off GOPROXY=https://goproxy.cn,direct go


linux-bin:
	GOOS=linux $(GO) build -tags netgo -o service service.go 

image:linux-bin
	docker build -f ./Dockerfile -t yeahgo_compare_price_service/service:1.1 .
	
devbuild:linux-bin
	docker build -f Dockerfile-ENV -t registry.cn-shenzhen.aliyuncs.com/yongdaxing/golang-mall-service:DEV  ./
	
prodbuild:linux-bin
	docker build -f Dockerfile-ENV -t registry.cn-shenzhen.aliyuncs.com/yongdaxing/golang-mall-service:prod  ./

clean:
	rm -f service