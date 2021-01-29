package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/liuhongdi/unittest04/controller"
	"github.com/liuhongdi/unittest04/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

var mock sqlmock.Sqlmock

//初始化
func init() {
	//创建sqlmock
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if nil != err {
		log.Fatalf("Init sqlmock failed, err %v", err)
	}
	//结合gorm、sqlmock
	global.DBLink, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn: db,
	}), &gorm.Config{})
	if nil != err {
		log.Fatalf("Init DB with sqlmock failed, err %v", err)
	}
}

//测试得到一件商品的信息,启用sqlmock,放回结果集
func TestOneMockRes(t *testing.T) {
	goodsId:=1
	//创建数据库记录
	rows := sqlmock.NewRows([]string{"goodsId", "goodsName", "subject", "price", "stock"}).
		AddRow(2, "Moonii陶瓷月相灯", "这是一件测试商品", "5.32", "33")
	mock.ExpectQuery("^SELECT \\* FROM `goods` WHERE goodsId=\\? ORDER BY `goods`.`goodsId` LIMIT 1").
		WithArgs(goodsId).WillReturnRows(rows)
	//执行
	goods,err:=controller.GoodsOne(goodsId)
	if (err != nil) {
		t.Fatalf("goodsId: %d, err:%v", goodsId, err)
	}else {
			fmt.Println(goods)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

//测试得到一件商品的信息,启用sqlmock,无结果集
func TestOneMockNoRes(t *testing.T) {
	goodsId:=1

	//创建数据库记录
	rows := sqlmock.NewRows([]string{"goodsId", "goodsName", "subject", "price", "stock"})
		//AddRow(2, "Moonii陶瓷月相灯", "这是一件测试商品", "5.32", "33")
	mock.ExpectQuery("^SELECT \\* FROM `goods` WHERE goodsId=\\? ORDER BY `goods`.`goodsId` LIMIT 1").
		WithArgs(goodsId).WillReturnRows(rows)
	//执行
	goods,err:=controller.GoodsOne(goodsId)
	if (err != nil) {
		t.Fatalf("goodsId: %d, err:%v", goodsId, err)
	}else {
			fmt.Println(goods)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

//测试得到一件商品的信息,启用sqlmock,返回数据库错误
func TestOneMockError(t *testing.T) {
	goodsId:=1
	//传递sql
	mock.ExpectQuery("^SELECT \\* FROM `goods` WHERE goodsId=\\? ORDER BY `goods`.`goodsId` LIMIT 1").
		WithArgs(goodsId).WillReturnError(errors.New("some error"))
	//执行
	goods,err:=controller.GoodsOne(goodsId)
	if (err != nil) {
		t.Fatalf("goodsId: %d, err:%v", goodsId, err)
	}else {
		fmt.Println(goods)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
