// @BeeOverwrite YES
// @BeeGenerateTime 20200820_195417
package api

type ResponseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
