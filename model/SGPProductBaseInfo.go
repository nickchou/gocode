package model

type SGPProductBaseInfo struct {
	PBIId                 int64
	Serialid              string
	PBITitle              string
	PBIPlayWay            int
	PBIMinDays            int
	PBIMaxDays            int
	PBINights             int
	PBIAdvanceBookDay     int
	PBIBelong             int
	PBIOldProductId       int
	PBIProductType        int
	PBIProductPackageType int
	PBIProductOperateType int
	PBIChannels           int64
	PBIOnlineStatus       int
	PBIProductAttribute   int
}
