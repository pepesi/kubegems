package apis

type NotFoundResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BadRequest struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PageSizeOptions struct {
	Page uint `json:"page"`
	Size uint `json:"size"`
}
