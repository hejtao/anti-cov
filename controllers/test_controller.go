package controllers

import "antiCov-server/utils"

// @description 测试test类型接口
// @router /get [get]
func (c *PublicController) TestGet() {

	c.ReturnSuccess(1, "ok", utils.GetGeoWithIp(c.GetString("ip")))
}
