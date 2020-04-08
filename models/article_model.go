package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id int `json:"id"`

	AdminId int    `json:"admin_id"`
	Cover   string `json:"cover"`
	Title   string `json:"title"`
	Content string `json:"content" orm:"size(4095)"`
	Section int    `json:"section"`

	Hidden     bool      `json:"hidden,omitempty"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`
}

var ArticleSection = []int{
	0,
	1, //科普信息
	2, //视频区
	3, //指南区
	4, //答疑区
	5, //案例区
	6, //故事区
}

func init() {
	orm.RegisterModel(new(Article))
}

func CreateArticle(article *Article) (int, error) {
	id, err := orm.NewOrm().Insert(article)
	return int(id), err
}

func UpdateArticle(article *Article, cols ...string) error {
	_, err := orm.NewOrm().Update(article, cols...)
	return err
}

// 如果返回的 article.Id == 0, 则出错
func GetArticleById(id int, rel ...interface{}) *Article {
	o := orm.NewOrm()
	article := new(Article)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Article)).Filter("id", id).One(article)
	} else {
		_ = o.QueryTable(new(Article)).Filter("id", id).RelatedSel(rel...).One(article)
	}

	return article
}

func DeleteArticlesByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Article)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Article)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetArticles(cond *orm.Condition, page, size int, rel ...interface{}) interface{} {
	if size < 1 {
		size = 10
	}

	o := orm.NewOrm()
	articles := make([]*Article, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Article)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(size, size*(page-1)).
			All(&articles)

		data := make(map[string]interface{})
		total, _ := o.QueryTable(new(Article)).SetCond(cond).Filter("hidden", false).Count()
		data["totalCount"] = total
		data["pageSize"] = size
		data["pageIndex"] = page

		if int(total)%size == 0 {
			data["pageCount"] = int(total) / size
		} else {
			data["pageCount"] = int(total)/size + 1
		}

		data["list"] = articles

		return data
	}

	_, _ = o.QueryTable(new(Article)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&articles)
	return articles
}

func ArticleExist(cond *orm.Condition) bool {
	return orm.NewOrm().QueryTable(new(Article)).SetCond(cond).Exist()
}
