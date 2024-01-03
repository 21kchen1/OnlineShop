package test

import "onlineshop/models"

var testLog = []models.Log {
	{
		UserID: 1,
		Title: "打牌",
		Content: "输了",
	},
	{
		UserID: 1,
		Title: "打牌",
		Content: "赢了",
	},
	{
		UserID: 1,
		Title: "打牌",
		Content: "阿黑啊阿黑，你怎可如此堕落",
	},
	{
		UserID: 1,
		Title: "打牌",
		Content: "好好好",
	},
}

func addLog() {
	for _, i := range testLog {
		models.CreateALog(&i)
	}
}