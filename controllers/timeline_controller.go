package controllers

import "github.com/astaxie/beego/orm"

// @description 获取全部timeline
// @router /timeline/from [get]
func (c *PublicController) GetAllTimelines() {
	sql := `SELECT 'from', count(*) FROM timeline WHERE record_id IN (77, 78) GROUP BY 'from';`

	timelines := make([]orm.ParamsList, 0)
	if _, err := orm.NewOrm().Raw(sql).ValuesList(&timelines); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", timelines)
}
