package parameter

type ServicePackage struct {
	PackageName    string  `json:"package_name"`    // 服务包名称
	City           string  `json:"city"`            // 所属城市
	Sales          int     `json:"sales"`           // 销量
	Price          float64 `json:"price"`           // 售价
	OriginalPrice  float64 `json:"original_price"`  // 原始价格
	DurationMonths int     `json:"duration_months"` // 服务时长(按月计算)
	Desc           string  `json:"desc"`            // 描述
}

type PkgList struct {
	List []ServicePackage `json:"list"` //服务包列表
}
