package controllers

// @description 测试test类型接口
// @router /get [get]
func (c *PublicController) TestGet() {

	c.ReturnSuccess(1, "ok", nil)
}
