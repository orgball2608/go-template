package internal

import "time"

type User struct {
	Id                   int       `gorm:"column:id;primaryKey;autoIncrement"`
	Email                string    `gorm:"column:mail_e;type:varchar(80)"`
	Type                 string    `gorm:"column:type__c;type:varchar(18)"`
	Name                 string    `gorm:"column:name;type:varchar(80)"`
	IsMiddleManager      bool      `gorm:"column:ismiddlemanager__c"`
	IsProxy              bool      `gorm:"column:isproxy__c"`
	Password             string    `gorm:"column:password__c;type:varchar(255)"`
	Status               string    `gorm:"column:status__c"`
	DepartmentName       string    `gorm:"column:departmentname__c;type:varchar(255)"`
	LastLoginAt          time.Time `gorm:"column:lastlogin_at__c;type:timestamp"`
	LastChangePasswordAt time.Time `gorm:"column:lastchangepassword_at__c;type:timestamp"`
	CreatedAt            time.Time `gorm:"column:recordcreatetime__c;autoCreateTime;type:timestamp"`
	UpdatedAt            time.Time `gorm:"column:recordupdatetime__c;autoUpdateTime;type:timestamp"`
	ExternalId           string    `gorm:"column:externalid__c;varchar(50)"`
	RecordCreateUserId   string    `gorm:"column:recordcreateuserid__c;varchar(18)"`
	RecordUpdateUserId   string    `gorm:"column:recordupdateuserid__c;varchar(18)"`
}

type UserToken struct {
	Id                 int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserId             int       `gorm:"column:user_id;index:mmsp_t_heroku_usertoken_i_userid;not null"` //replace: mmsp_m_herokuuser__c_id
	Token              string    `gorm:"column:token;type:varchar(255);not null"`
	ExpireAt           time.Time `gorm:"column:expire_at;type:timestamp;not null"` //replace: expiration_time
	IsDeleted          bool      `gorm:"column:is_deleted"`
	ResendTime         time.Time `gorm:"column:resend_time;type:timestamp;not null"`    //replace: token_resendable_time
	RecordCreateUserId string    `gorm:"column:record_create_user_id;type:varchar(18)"` //replace: int
	RecordCreateTime   time.Time `gorm:"column:record_create_time;autoCreateTime;type:timestamp"`
	RecordUpdateUserId string    `gorm:"column:record_update_user_id;type:varchar(18)"` //replace: int
	RecordUpdateTime   time.Time `gorm:"column:record_update_time;autoUpdateTime;type:timestamp"`
}
