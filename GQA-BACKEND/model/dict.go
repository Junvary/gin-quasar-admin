package model

type SysDict struct {
	GqaModelWithCreatedByAndUpdatedBy
	//ParentCode <==> DictCode
	ParentCode string    `json:"parent_code" gorm:"comment:父字典编码;index"`
	DictCode   string    `json:"dict_code" gorm:"comment:字典编码;not null;uniqueIndex;"`
	DictLabel  string    `json:"dict_label" gorm:"comment:字典名称;not null;"`
	DictExt1   string    `json:"dict_ext_1" gorm:"comment:字典扩展项1;"`
	DictExt2   string    `json:"dict_ext_2" gorm:"comment:字典扩展项2;"`
	DictExt3   string    `json:"dict_ext_3" gorm:"comment:字典扩展项3;"`
	DictExt4   string    `json:"dict_ext_4" gorm:"comment:字典扩展项4;"`
	DictExt5   string    `json:"dict_ext_5" gorm:"comment:字典扩展项5;"`
	Children   []SysDict `json:"children" gorm:"foreignKey:ParentCode;references:DictCode"`
}

type RequestGetDictList struct {
	RequestPageAndSort
	ParentCode string `json:"parent_code"`
	DictCode   string `json:"dict_code"`
	DictLabel  string `json:"dict_label"`
}

type RequestAddDict struct {
	RequestAdd
	ParentCode string `json:"parent_code"`
	DictCode   string `json:"dict_code"`
	DictLabel  string `json:"dict_label"`
	DictExt1   string `json:"dict_ext_1"`
	DictExt2   string `json:"dict_ext_2"`
	DictExt3   string `json:"dict_ext_3"`
	DictExt4   string `json:"dict_ext_4"`
	DictExt5   string `json:"dict_ext_5"`
}
