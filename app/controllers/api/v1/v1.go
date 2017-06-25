package v1

import (
	"github.com/revel/revel"
	"github.com/shiro16/golang-bbs/app/utils"
	"net/http"
)

//埋め込みによる revel.Controllerをラップした　ApiV1Controllerを定義する
type ApiV1Controller struct {
	*revel.Controller
}

// エラーの際に返す　Json用の構造体
type ErrorResponse struct {
	Code int		`json:"code"`
	Message string	`json:"message"`
}

//正常な際に返す　Json用の構造体　（今回は１種類で統一する）
type Response struct {
	Result interface{} `json:"results"`
}

//引数として渡されて interface　にリクエストのjsonの値を格納する
func (c *ApiV1Controller) BindParams(s interface{} ) error {
	return utils.JsonDecode(c.Request.Body, s)
}

// Bad Request Error
func (c *ApiV1Controller) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := ErrorResponse{c.Response.Status, s}
	return  c.RenderJSON(r)
}

// Not Found Error
func (c *ApiV1Controller) HandleNotFoundError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Internal Server Error
func (c *ApiV1Controller) HandleInternalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}