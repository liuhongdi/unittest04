package controller

import (
	"github.com/liuhongdi/unittest04/global"
	"github.com/liuhongdi/unittest04/model"
)

func GoodsOne(goodsId int) (*model.Goods, error) {
	goodsOne:=&model.Goods{}
	err := global.DBLink.Where("goodsId=?",goodsId).First(&goodsOne).Error
	//fmt.Println(err)
	if (err != nil) {
		return nil,err
	} else {
		return goodsOne,nil
	}
}
