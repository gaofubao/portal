### 1. "健康检查"

1. route definition

- Url: /health
- Method: GET
- Request: `-`
- Response: `-`

2. request definition



3. response definition


### 2. "交流"

1. route definition

- Url: /api/v1/ai/chat
- Method: POST
- Request: `ChatReq`
- Response: `ChatResp`

2. request definition



```golang
type ChatReq struct {
	Question string `json:"question" form:"question"`
}
```


3. response definition



```golang
type ChatResp struct {
	Answer string `json:"answer" form:"answer"`
}
```

