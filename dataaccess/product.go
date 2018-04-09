package dataaccess

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nickchou/gocode/model"
)

var (
	DBHOST  = "10.100.158.96:3500"                //IP地址
	DBNACCT = "TCGreatHolidaySuperGoProduct_test" //用户名
	DBPWD   = "THiKrtercarNIBWYBbTG1lnY"          //密码
	DBNAME  = "TCGreatHolidaySuperGoProduct"      //库名
)

func GetProduct(pro model.SGPProductBaseInfo) model.CalendarModel {

	//要返回的价格日历实体
	calm := model.CalendarModel{ProBaseInfo: pro}
	//获取连接字符串 root:123456@tcp(127.0.0.1:3306)/Test?charset=utf8
	dbconn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", DBNACCT, DBPWD, DBHOST, DBNAME)
	db, err := sql.Open("mysql", dbconn)
	defer db.Close()

	if err == nil {
		//列名必须要和Scan保持一致，否则会报错
		stmt, _ := db.Prepare("select PBIId,Serialid,PBITitle,PBIPlayWay,PBIMinDays,PBIMaxDays,PBINights,PBIAdvanceBookDay,PBIBelong,PBIOldProductId,PBIProductType,PBIProductPackageType,PBIProductOperateType,PBIChannels,PBIOnlineStatus,PBIProductAttribute from SGPProductBaseInfo where PBIOldProductId=? AND PBIIfDel = 0 AND PBIDataFlag =1 AND PBICheckStaus =2 AND PBIOnlineStatus =1 and PBIBelong=?")
		r := stmt.QueryRow(calm.ProBaseInfo.PBIOldProductId, calm.ProBaseInfo.PBIBelong)
		err := r.Scan(&pro.PBIId, &pro.Serialid, &pro.PBITitle, &pro.PBIPlayWay, &pro.PBIMinDays, &pro.PBIMaxDays, &pro.PBINights, &pro.PBIAdvanceBookDay, &pro.PBIBelong, &pro.PBIOldProductId, &pro.PBIProductType, &pro.PBIProductPackageType, &pro.PBIProductOperateType, &pro.PBIChannels, &pro.PBIOnlineStatus, &pro.PBIProductAttribute)
		if err == nil {
			go getproductSingleRelated(db, calm)
		}
	}
	return calm
}
func getproductSingleRelated(db *sql.DB, calm model.CalendarModel) {
	stmt, _ := db.Prepare("SELECT PSRId,PSRProductId,PSRResourceId,PSRType,PSRIsMustChoose,PSRUseDays FROM SGPProductSingleRelated_New WHERE PSRProductId =? and PSRIsMustChoose =1 and PSRDataFlag=1 and PSRIfDel=1")
	rows, _ := stmt.Query(calm.ProBaseInfo.PBIOldProductId)
	defer rows.Close()
	for rows.Next() {
		prosin := model.SGPProductSingleRelated_New{}
		rows.Scan(&prosin.PSRId, &prosin.PSRProductId, &prosin.PSRResourceId, &prosin.PSRType, &prosin.PSRIsMustChoose, &prosin.PSRUseDays)
		append(calm.ProSingleRel, prosin)
		for index, value := range calm.ProSingleRel {
			fmt.Println(index)
		}
	}
}
