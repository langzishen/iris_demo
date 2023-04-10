package dao

// 用户表
type Users struct {
	Id       uint   `gorm:"column:id;type:mediumint(9) unsigned;primary_key;AUTO_INCREMENT;comment:表id" json:"id"`
	Email    string `gorm:"column:email;type:varchar(200);comment:邮件;NOT NULL" json:"email"`
	Password string `gorm:"column:password;type:varchar(32);comment:密码;NOT NULL" json:"password"`
	EcSalt   string `gorm:"column:ec_salt;type:varchar(11);NOT NULL" json:"ec_salt"`
	Sex      uint   `gorm:"column:sex;type:tinyint(4) unsigned;default:0;comment:0 保密 1 男 2 女;NOT NULL" json:"sex"`
}

func (m *Users) TableName() string {
	return "users"
}
