module sdk-demo-go

go 1.16

require (
    github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.817 // indirect
    github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid v1.0.817 // indirect
)

replace (
    github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid v1.0.817  => ./tencentcloud/tdid
)
