package bean

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
