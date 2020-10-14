package home

import "mofan-blog/app/model/document"

func List(limit int, page int) []*document.Entity {

	docList, _ := document.Model.Order("create_time desc").All()

	return docList
}
