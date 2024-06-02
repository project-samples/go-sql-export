package app

import (
	"context"
	"database/sql"
	"path/filepath"
	"time"

	"github.com/core-go/io/export"
	"github.com/core-go/io/writer"
	"github.com/core-go/io/writer/formatter"
	_ "github.com/lib/pq"
)

type ApplicationContext struct {
	Export func(ctx context.Context) (int64, error)
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	db, err := sql.Open(cfg.Sql.Driver, cfg.Sql.DataSourceName)
	if err != nil {
		return nil, err
	}
	formatWriter, err := formatter.NewFixedLengthFormatter[User]()
	if err != nil {
		return nil, err
	}
	writer, err := writer.NewFileWriter(GenerateFileName)
	if err != nil {
		return nil, err
	}
	exportService, err := export.NewExporter(db, BuildQuery, formatWriter.Format, writer.Write, writer.Close)
	if err != nil {
		return nil, err
	}
	return &ApplicationContext{
		Export: exportService.Export,
	}, nil
}

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" format:"%011s" length:"11" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Username    string     `json:"username" gorm:"column:username" bson:"username" length:"10" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Test        string     `json:"test" gorm:"-" bson:"username" length:"0" format:"-" dynamodbav:"test" firestore:"test" validate:"required,username,max=100"`
	Email       *string    `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" length:"31" validate:"email,max=100"`
	Phone       string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" length:"20" validate:"required,phone,max=18"`
	Status      bool       `json:"status" gorm:"column:status" true:"1" false:"0" bson:"status" dynamodbav:"status" format:"%5s" length:"5" firestore:"status" validate:"required"`
	CreatedDate *time.Time `json:"createdDate" gorm:"column:createdDate" bson:"createdDate" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate" firestore:"createdDate" validate:"required"`
}

func BuildQuery(ctx context.Context) (string, []interface{}) {
	query := "select id, username, email, phone, status, createdDate from userimport"
	params := make([]interface{}, 0)
	return query, params
}

func GenerateFileName() string {
	fileName := time.Now().Format("20060102150405") + ".csv"
	fullPath := filepath.Join("export", fileName)
	writer.DeleteFile(fullPath)
	return fullPath
}
