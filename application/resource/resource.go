package resource

type MessageReceivingResource struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type MessageSendingResource struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
