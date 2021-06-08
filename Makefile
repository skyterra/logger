.PHONY: pre-install test cover cover-html cover-func clean

pre-install:
	go get -u github.com/gin-gonic/gin && \
	go get -u github.com/onsi/ginkgo/ginkgo

# 执行单元测试
test:
	go test ./...

# 统计覆盖率
cover:
	go test ./... -coverprofile cover.profile

# 打开浏览器显示覆盖统计信息
cover-html:
	go tool cover -html=cover.profile

# 显示函数覆盖统计
cover-func:
	go tool cover -func=cover.profile

# 清理
clean:
	rm -rf bin
	rm cover.profile