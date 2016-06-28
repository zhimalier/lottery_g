package controllers

import (
	//"fmt"
	"math/rand"
	"strconv"
	"time"

	"lottery_g/models"

	"github.com/astaxie/beego"
)

type LotteryController struct {
	beego.Controller
}

var lotteryResult *models.LotteryResult

func (this *LotteryController) Get() {
	prize := getLotteryLevel()
	this.Data["json"] = map[string]interface{}{"prize": prize}
	this.ServeJSON()
	//c.Data["Email"] = "zhangmi@sobey.com"
	//c.TplName = "index.tpl"
}

func getLotteryLevel() int {
	tNow := time.Now()
	timeNow := tNow.Format("20060102")
	lotteryID, _ := strconv.Atoi(timeNow) //取当天的时间戳年月日作为抽奖活动ID

	//	r := rand.New(rand.NewSource(time.Now().UnixNano())) //第1次随机数作为用户ID使用
	//	userID := r.Intn(100000)
	userID := rand.Intn(100000)
	//fmt.Println("在设定ID范围内的随机userID:", userID)

	//如果该用户已经抽过该次奖，则本次无效
	if lotteryResult.GetResult(lotteryID, userID) {
		//fmt.Println("该用户已经抽过奖了:", userID)
		return -1
	}

	//查询出已抽出奖项列表
	drawedPrizeMap := lotteryResult.GetDrawedLottery(lotteryID)
	//fmt.Println("已抽出奖项列表:", drawedPrizeMap)

	lotteryList := [5]int{1, 1, 2, 3, 10000}
	if drawedPrizeMap != nil {
		for n := range lotteryList {
			count, ok := drawedPrizeMap[strconv.Itoa(n)]
			if ok {
				countStr, _ := count.(string)
				counti, _ := strconv.Atoi(countStr)
				if lotteryList[n] >= counti {
					lotteryList[n] -= counti
				} else {
					lotteryList[n] = 0
				}
			}
		}
	}

	//fmt.Println("减去已抽完的奖项还剩的奖项列表:", lotteryList)

	sum := 0
	for i := range lotteryList {
		sum += lotteryList[i]
	}

	var prize = 4
	//fmt.Println("剩余的奖项总的数量sum:", sum)
	if sum <= 0 { //奖项全部抽完了
		lotteryResult.AddResult(lotteryID, userID, prize)
		return prize
	}

	//randNum := r.Intn(sum) //第2次随机数作为抽奖数
	randNum := rand.Intn(sum) //第2次随机数作为抽奖数
	//fmt.Println("抽奖用的随机数randNum:", randNum)

	for j := range lotteryList {
		if randNum < lotteryList[j] {
			prize = j
			break
		} else {
			randNum -= lotteryList[j]
		}
	}
	//fmt.Println("抽奖结果:", prize)
	lotteryResult.AddResult(lotteryID, userID, prize)
	//beego.Debug(userID, drawedPrizeMap, lotteryList, randNum, prize)

	return prize
}
