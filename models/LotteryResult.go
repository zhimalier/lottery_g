package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER_NAME   = "mysql"
	DATA_SOURCE   = "root:root@tcp(localhost:3306)/lottery?charset=utf8"
	MAX_IDLE_CONN = 5
	MAX_OPEN_CONN = 30
)

type LotteryResult struct {
	Id        int `orm:"pk;auto"`
	Lotteryid int
	Userid    int
	Prize     int
}

func init() {
	orm.RegisterDriver(DRIVER_NAME, orm.DRMySQL)

	orm.RegisterDataBase("default", DRIVER_NAME, DATA_SOURCE, MAX_IDLE_CONN, MAX_OPEN_CONN)

	orm.RegisterModel(new(LotteryResult))

	//orm.Debug = true
}

//根据抽奖活动ID和用户ID查询该用户是否已经抽过奖
func (lottery_Result *LotteryResult) GetResult(lotteryID, userID int) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LotteryResult))
	qs = qs.Filter("lotteryid", lotteryID).Filter("userid", userID)
	ex := qs.Exist()
	return ex

}

func (lottery_Result *LotteryResult) GetPrizeCount(lotteryID, prize int) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LotteryResult))
	qs = qs.Filter("lotteryid", lotteryID).Filter("prize", prize)
	count, _ := qs.Count()
	return count
}

func (lottery_Result *LotteryResult) AddResult(lotteryID, userID, prize int) int64 {
	o := orm.NewOrm()
	result := LotteryResult{Lotteryid: lotteryID, Userid: userID, Prize: prize}
	id, err := o.Insert(&result)
	if err != nil {
		return -1
	}
	return id
}

//查询已经抽出的奖项列表
func (lottery_Result *LotteryResult) GetDrawedLottery(lotteryID int) map[string]interface{} {
	o := orm.NewOrm()
	res := make(orm.Params)
	o.Raw("SELECT prize p,COUNT(*) c FROM lottery_result where lotteryid=? group by prize", lotteryID).RowsToMap(&res, "p", "c")
	return res
}
