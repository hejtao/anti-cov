package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Timeline struct {
	Id int `json:"id,omitempty"`

	RecordId int `json:"record_id"`

	Date    string `json:"date"`
	Hour    string `json:"hour"`
	Minute  string `json:"minute"`
	Second  string `json:"time"`
	PerHour string `json:"per_hour"`

	Ip        string `json:"ip"`
	Continent string `json:"continent"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`

	Hidden bool `json:"hidden,omitempty"`
}

func init() {
	orm.RegisterModel(new(Timeline))
}

func CreateTimeline(timeline *Timeline) (int, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	timeline.Date = now[:10]
	timeline.Hour = now[:13]
	timeline.Minute = now[:16]
	timeline.Second = now
	timeline.PerHour = now[11:13]
	id, err := orm.NewOrm().Insert(timeline)
	return int(id), err
}

func UpdateTimeline(timeline *Timeline, cols ...string) error {
	_, err := orm.NewOrm().Update(timeline, cols...)
	return err
}

// 如果返回的 timeline.Id == 0, 则出错
func GetTimelineById(id int, rel ...interface{}) *Timeline {
	o := orm.NewOrm()
	timeline := new(Timeline)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Timeline)).Filter("id", id).One(timeline)
	} else {
		_ = o.QueryTable(new(Timeline)).Filter("id", id).RelatedSel(rel...).One(timeline)
	}

	return timeline
}

func DeleteTimelinesByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Timeline)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Timeline)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetTimelines(cond *orm.Condition, page int, rel ...interface{}) interface{} {
	o := orm.NewOrm()
	timelines := make([]*Timeline, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Timeline)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(10, 10*(page-1)).
			All(&timelines)

		data := make(map[string]interface{})
		data["count"], _ = o.QueryTable(new(Timeline)).SetCond(cond).Filter("hidden", false).Count()
		data["timelines"] = timelines

		return data
	}

	_, _ = o.QueryTable(new(Timeline)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&timelines)
	return timelines
}
