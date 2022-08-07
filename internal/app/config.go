package app

import "github.com/core-go/log/zap"

type Config struct {
	Sql DBConfig   `mapstructure:"sql"`
	Log log.Config `mapstructure:"log"`
}
type DBConfig struct {
	DataSourceName string `mapstructure:"data_source_name" json:"dataSourceName,omitempty" gorm:"column:datasourcename" bson:"dataSourceName,omitempty" dynamodbav:"dataSourceName,omitempty" firestore:"dataSourceName,omitempty"`
	Driver         string `mapstructure:"driver" json:"driver,omitempty" gorm:"column:driver" bson:"driver,omitempty" dynamodbav:"driver,omitempty" firestore:"driver,omitempty"`
}
