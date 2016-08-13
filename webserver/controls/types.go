package controls

type ErrorMessage struct {
	Result     string `json:"result"`
	Note       string `json:"note"`
	ResultCode int    `json:"result_code"`
}