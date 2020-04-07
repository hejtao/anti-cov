package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// @description
// @router /timeline/graph1 [get]
func (c *PublicController) GetGraph1Data() {
	country := c.GetString("country")
	if country != "" {
		country = fmt.Sprintf("AND country = '%s' ", country)
	}

	type data struct {
		Date   string `json:"date"`
		Source string `json:"source"`
		Total  int    `json:"total"`
	}

	ds := make([]*data, 0)

	sql := fmt.Sprintf(
		`SELECT date, source, count(*) total FROM timeline WHERE record_id IN (77, 78) %sGROUP BY date, source;`,
		country,
	)

	if _, err := orm.NewOrm().Raw(sql).QueryRows(&ds); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	retData := make(map[string]map[string]int)
	for _, d := range ds {
		if retData[d.Date] == nil {
			retData[d.Date] = map[string]int{
				d.Source: d.Total,
			}
		} else {
			retData[d.Date][d.Source] = d.Total
		}

	}

	c.ReturnSuccess(1, "ok", retData)
}

// @description
// @router /timeline/graph2 [get]
func (c *PublicController) GetGraph2Data() {
	country := c.GetString("country")
	if country != "" {
		country = fmt.Sprintf("AND country = '%s' ", country)
	}

	type data struct {
		Date     string `json:"date"`
		Terminal string `json:"terminal"`
		Total    int    `json:"total"`
	}

	ds := make([]*data, 0)

	sql := fmt.Sprintf(
		`SELECT date, terminal, count(*) total FROM timeline WHERE record_id IN (77, 78) %sGROUP BY date, terminal;`,
		country,
	)

	if _, err := orm.NewOrm().Raw(sql).QueryRows(&ds); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	retData := make(map[string]map[string]int)
	for _, d := range ds {
		if retData[d.Date] == nil {
			retData[d.Date] = map[string]int{
				d.Terminal: d.Total,
			}
		} else {
			retData[d.Date][d.Terminal] = d.Total
		}

	}

	c.ReturnSuccess(1, "ok", retData)
}

// @description
// @router /timeline/graph3 [get]
func (c *PublicController) GetGraph3Data() {
	country := c.GetString("country")
	if country != "" {
		country = fmt.Sprintf("AND country = '%s' ", country)
	}

	type data struct {
		PerHour string `json:"per_hour"`
		Source  string `json:"source"`
		Total   int    `json:"total"`
	}

	ds := make([]*data, 0)

	sql := fmt.Sprintf(
		`SELECT per_hour, source, count(*) total FROM timeline WHERE record_id IN (77, 78) %sGROUP BY per_hour, source;`,
		country,
	)

	if _, err := orm.NewOrm().Raw(sql).QueryRows(&ds); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	retData := make(map[string]map[string]int)
	for _, d := range ds {
		if retData[d.PerHour] == nil {
			retData[d.PerHour] = map[string]int{
				d.Source: d.Total,
			}
		} else {
			retData[d.PerHour][d.Source] = d.Total
		}

	}

	c.ReturnSuccess(1, "ok", retData)
}

// @description
// @router /timeline/graph4 [get]
func (c *PublicController) GetGraph4Data() {
	country := c.GetString("country")
	if country != "" {
		country = fmt.Sprintf("AND country = '%s' ", country)
	}

	type data struct {
		PerHour  string `json:"per_hour"`
		Terminal string `json:"terminal"`
		Total    int    `json:"total"`
	}

	ds := make([]*data, 0)

	sql := fmt.Sprintf(
		`SELECT per_hour, terminal, count(*) total FROM timeline WHERE record_id IN (77, 78) %sGROUP BY per_hour, terminal;`,
		country,
	)

	if _, err := orm.NewOrm().Raw(sql).QueryRows(&ds); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	retData := make(map[string]map[string]int)
	for _, d := range ds {
		if retData[d.PerHour] == nil {
			retData[d.PerHour] = map[string]int{
				d.Terminal: d.Total,
			}
		} else {
			retData[d.PerHour][d.Terminal] = d.Total
		}

	}

	c.ReturnSuccess(1, "ok", retData)
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
