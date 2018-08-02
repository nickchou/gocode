package model

import (
	"time"
)

type CalendarModel struct {
	ProBaseInfo         SGPProductBaseInfo
	ProSingleRel        []SGPProductSingleRelated_New
	ProBaseInfoExt      SGPProductBaseInfoExtends_New
	ProTripResource     []SGPProductTripResource
	ProPriceStrategy    []SGPProductPriceStrategy_New
	ProDefPriceStrategy []SGPProductPriceUnifiedStrategy_New
	ProTripStaticRes    []SGPProductTripStaticResource
}

//单产品资源关系表
type SGPProductSingleRelated_New struct {
	PSRId           int64
	PSRProductId    int64
	PSRResourceId   int64
	PSRType         int
	PSRIsMustChoose int
	PSRUseDays      string
}

//产品扩展信息表
type SGPProductBaseInfoExtends_New struct {
	PBIEId                 int64
	PBIEProductId          int64
	PBIENewWarehouseProfit float32
	PBIEPurchasingType     int
	PBIDays                int
}

//产品资源信息表
type SGPProductTripResource struct {
	PTRId                  int64
	PTRProductId           int64
	PTRTripNo              int
	PTRTrafficType         int
	PTRDeptId              int
	PTRDestId              int
	PTRDays                int
	PTRNights              int
	PTRContainResource     string
	PTRTrafficResourceType int
	PTRHotelResourceType   int
}

//销售策略
type SGPProductPriceStrategy_New struct {
	PPSId            int64
	PPSProductId     int64
	PPSSaleDate      time.Time
	PPSUpgradeId     int64
	PPSProfitType    int
	PPSStrategyType  int
	PPSAdultValue    float32
	PPSAdultMantissa float32
	PPSChildValue    float32
	PPSChildMantissa float32
	PPSRoomValue     float32
	PPSRoomMantissa  float32
}

//默认销售策略
type SGPProductPriceUnifiedStrategy_New struct {
	PPUSId            int64
	PPUSProductId     int64
	PPUSProfitType    int
	PPUSStrategyType  int
	PPUSAdultValue    float32
	PPUSAdultMantissa float32
	PPUSChildValue    float32
	PPUSChildMantissa float32
	PPUSRoomValue     float32
	PPUSRoomMantissa  float32
	PPUSUpgradeId     int64
}

//静态资源信息
type SGPProductTripStaticResource struct {
	PTSRId            int64
	PTSRPTRId         int64
	PTSRProductId     int64
	PTSRResourceId    int64
	PTSRResourceType  int
	PTSRBelongTo      int
	PTSRSourceFrom    int
	PTSRSellOrder     int
	PTSRCheckInNights string
}
