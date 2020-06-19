package model

const (
	TransactionLogTypeUpdate = 1
	TransactionLogTypeInsert = 2
)

// 事务日志
type TransactionLog struct {
	Id           int64  `db:"id"`
	Tid          string `db:"tid"`
	Type         int    `db:"type"`
	BeforeCol    string `db:"before_col"` // 变更前镜像，用于回滚
	AfterCol     string `db:"after_col"`  // 变更后的镜像
	Table        string `db:"table_name"`
	Connection   string `db:"connection_str"`
	PrimaryKey   string `db:"primary_key"`   // db 主键
	PrimaryValue string `db:"primary_value"` // db 主键
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
