package controllers

import "github.com/astaxie/beego/orm"

// @description
// @router /timeline/graph1 [get]
func (c *PublicController) GetGraph1Data() {

	type data struct {
		Date   string `json:"date"`
		Source string `json:"source"`
		Num    int    `json:"num"`
	}

	ds := make([]*data, 0)

	sql := `SELECT date, source, count(*) num FROM timeline WHERE record_id IN (77, 78) GROUP BY date, source;`
	if _, err := orm.NewOrm().Raw(sql).QueryRows(&ds); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", ds)
}

// @description
// @router /timeline/countries [get]
func (c *PublicController) GetTimelineCountries() {
	countries := make([]string, 0)

	sql := `SELECT DISTINCT country FROM timeline WHERE country != '' ORDER BY CONVERT(country using gbk);`
	if _, err := orm.NewOrm().Raw(sql).QueryRows(&countries); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", countries)
}
