#!/bin/bash

# 功能测试
go test -v .

# 代码覆盖率测试
go test -coverprofile=c.out && \
    go tool cover -html=c.out 


# 性能测试
go test -bench .

