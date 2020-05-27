package models

import (
	"cloudPlatformDemo/middleware"
	"cloudPlatformDemo/utils"
	"xorm.io/builder"
)

type CostValidator struct {
	ProductLine int64 `validate:"required,oneof=1 2 3"`
	ProductType int64 `validate:"required,oneof=1 2 3 4 5 6"`
	IpType 		string `validate:"oneof=IPV4 IPV6"`
}

type Cost struct {
	Id 			        int64 	//默认自增主键
	ProductLine         int8 	`xorm:"tinyint(4) notnull default(1) 'product_line'"`
	CostCode            string  `xorm:"varchar(255) not null default('') 'cost_code'"`
	ProductType         int8	`xorm:"tinyint(4) notnull default(1) 'product_type'"`
	IdcId		        int64	`xorm:"int(11) notnull default(0) 'idc_id'"`
	IdcNo		        string	`xorm:"varchar(255) not null default('') 'idc_no'"`
	CdnNo		        string	`xorm:"varchar(255) not null default('') 'cdn_no'"`
	Price		        float64	`xorm:"decimal(10,2) not null default(0.00) 'price'"`
	Unit		        string	`xorm:"varchar(255) not null default('') 'unit'"`
	RatedCurrent        int64   `xorm:"int(11) notnull default(0) 'rated_current'"`
	ServiceProvider     string	`xorm:"varchar(255) not null default('') 'service_provider'"`
	Supplier		    string	`xorm:"varchar(255) not null default('') 'supplier'"`
	ServerDepreciation  float64	`xorm:"decimal(10,2) not null default(0.00) 'server_depreciation'"`
	NetworkDepreciation float64	`xorm:"decimal(10,2) not null default(0.00) 'network_depreciation'"`
	CabinetCost         float64	`xorm:"decimal(10,2) not null default(0.00) 'cabinet_cost'"`
	OtherCost 			float64	`xorm:"decimal(10,2) not null default(0.00) 'other_cost'"`
	PowerConsumption 	float64	`xorm:"decimal(10,2) not null default(0.00) 'power_consumption'"`
	IpType 				string	`xorm:"varchar(255) not null default('') 'ip_type'"`
	LineType 			string	`xorm:"varchar(255) not null default('') 'line_type'"`
	ResourceName 		string	`xorm:"varchar(255) not null default('') 'resource_name'"`
	CreatedErp 			string	`xorm:"varchar(255) not null default('') 'created_erp'"`
	UpdatedErp 			string	`xorm:"varchar(255) not null default('') 'updated_erp'"`
	CreatedAt 			int64   `xorm:"int(11) notnull default(0) 'created_at'"`
	UpdatedAt 			int64	`xorm:"int(11) notnull default(0) 'updated_at'"`
	DeletedAt  			int64	`xorm:"int(11) notnull default(0) 'deleted_at'"`
	IsDel  				int8	`xorm:"tinyint(4) notnull default(0) 'is_del'"`
}

func GetCostList(params *CostValidator) ([]Cost, int64, error) {
	orm := utils.Db
	if err := orm.Sync2(new(Cost)); err != nil {
		middleware.Log.Error("fail to sync database: %v\n", err)
		return nil, 0, err
	}
	costArr := make([]Cost, 0)
	count, err := orm.Where(
		builder.Eq{
			"product_line": params.ProductLine,
			"product_type": params.ProductType,
		}).Desc("created_at").FindAndCount(&costArr)
	if err != nil {
		middleware.Log.Error()
		return nil, 0, err
	}
	middleware.Log.Info("count is:",count)
	return costArr, count, nil
}
