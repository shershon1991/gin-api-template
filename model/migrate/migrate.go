/**
 * @Description mysql迁移
 **/
package migrate

import (
	"fmt"
	"gorm.io/gorm"
	"shershon1991/gin-api-template/global"
	"shershon1991/gin-api-template/model/entity/user"
)

// 设置表信息
func setTableOption(tableComment string) *gorm.DB {
	value := fmt.Sprintf("ENGINE=InnoDB COMMENT='%s'", tableComment)
	return global.GvaMysqlClient.Set("gorm:table_options", value)
}

// 用户相关表
func userTable() {
	// 用户账号表
	_ = setTableOption("用户表").AutoMigrate(&user.User{})
	// 用户信息表
	_ = setTableOption("用户信息表").AutoMigrate(&user.UserInfo{})
}

// 数据表迁移
func AutoMigrate() {
	// 创建用户相关表
	userTable()
}
