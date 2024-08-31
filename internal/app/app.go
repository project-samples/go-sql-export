package app

import (
	"context"
	"database/sql"
	"path/filepath"
	"time"

	"github.com/core-go/io/export"
	w "github.com/core-go/io/writer"
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
	transformer, err := w.NewFixedLengthTransformer[User]()
	if err != nil {
		return nil, err
	}
	writer, err := w.NewFileWriter(GenerateFileName)
	if err != nil {
		return nil, err
	}
	exporter, err := export.NewExporter[User](db, BuildQuery, transformer.Transform, writer.Write, writer.Close)
	if err != nil {
		return nil, err
	}
	return &ApplicationContext{
		Export: exporter.Export,
	}, nil
}

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" format:"%011s" length:"11" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Username    string     `json:"username" gorm:"column:username" bson:"username" length:"10" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Email       string     `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" length:"31" validate:"email,max=100"`
	Phone       string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" length:"20" validate:"required,max=18"`
	Status      bool       `json:"status" gorm:"column:status" true:"1" false:"0" bson:"status" dynamodbav:"status" format:"%5s" length:"5" firestore:"status"`
	CreatedDate *time.Time `json:"createdDate" gorm:"column:createdDate" bson:"createdDate" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate" firestore:"createdDate" validate:"required"`

	Username1    string     `json:"username1" gorm:"column:username1" bson:"username1" length:"10" dynamodbav:"username1" firestore:"username1" validate:"required,username,max=100"`
	Email1       string     `json:"email1" gorm:"column:email1" bson:"email" dynamodbav:"email1" firestore:"email1" length:"31" validate:"email,max=100"`
	Phone1       string     `json:"phone1" gorm:"column:phone1" bson:"phone1" dynamodbav:"phone1" firestore:"phone1" length:"20" validate:"required,max=18"`
	Status1      bool       `json:"status1" gorm:"column:status1" true:"1" false:"0" bson:"status1" dynamodbav:"status1" format:"%5s" length:"5" firestore:"status1"`
	CreatedDate1 *time.Time `json:"createdDate1" gorm:"column:createdDate1" bson:"createdDate1" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate1" firestore:"createdDate1" validate:"required"`

	Username2    string     `json:"username2" gorm:"column:username2" bson:"username2" length:"10" dynamodbav:"username2" firestore:"username2" validate:"required,username,max=200"`
	Email2       string     `json:"email2" gorm:"column:email2" bson:"email" dynamodbav:"email2" firestore:"email2" length:"31" validate:"email,max=200"`
	Phone2       string     `json:"phone2" gorm:"column:phone2" bson:"phone2" dynamodbav:"phone2" firestore:"phone2" length:"20" validate:"required,max=28"`
	Status2      bool       `json:"status2" gorm:"column:status2" true:"2" false:"0" bson:"status2" dynamodbav:"status2" format:"%5s" length:"5" firestore:"status2"`
	CreatedDate2 *time.Time `json:"createdDate2" gorm:"column:createdDate2" bson:"createdDate2" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate2" firestore:"createdDate2" validate:"required"`

	Username3    string     `json:"username3" gorm:"column:username3" bson:"username3" length:"10" dynamodbav:"username3" firestore:"username3" validate:"required,username,max=300"`
	Email3       string     `json:"email3" gorm:"column:email3" bson:"email" dynamodbav:"email3" firestore:"email3" length:"31" validate:"email,max=300"`
	Phone3       string     `json:"phone3" gorm:"column:phone3" bson:"phone3" dynamodbav:"phone3" firestore:"phone3" length:"20" validate:"required,max=38"`
	Status3      bool       `json:"status3" gorm:"column:status3" true:"3" false:"0" bson:"status3" dynamodbav:"status3" format:"%5s" length:"5" firestore:"status3"`
	CreatedDate3 *time.Time `json:"createdDate3" gorm:"column:createdDate3" bson:"createdDate3" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate3" firestore:"createdDate3" validate:"required"`

	Username4    string     `json:"username4" gorm:"column:username4" bson:"username4" length:"10" dynamodbav:"username4" firestore:"username4" validate:"required,username,max=400"`
	Email4       string     `json:"email4" gorm:"column:email4" bson:"email" dynamodbav:"email4" firestore:"email4" length:"31" validate:"email,max=400"`
	Phone4       string     `json:"phone4" gorm:"column:phone4" bson:"phone4" dynamodbav:"phone4" firestore:"phone4" length:"20" validate:"required,max=48"`
	Status4      bool       `json:"status4" gorm:"column:status4" true:"4" false:"0" bson:"status4" dynamodbav:"status4" format:"%5s" length:"5" firestore:"status4"`
	CreatedDate4 *time.Time `json:"createdDate4" gorm:"column:createdDate4" bson:"createdDate4" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate4" firestore:"createdDate4" validate:"required"`

	Username5    string     `json:"username5" gorm:"column:username5" bson:"username5" length:"10" dynamodbav:"username5" firestore:"username5" validate:"required,username,max=500"`
	Email5       string     `json:"email5" gorm:"column:email5" bson:"email" dynamodbav:"email5" firestore:"email5" length:"31" validate:"email,max=500"`
	Phone5       string     `json:"phone5" gorm:"column:phone5" bson:"phone5" dynamodbav:"phone5" firestore:"phone5" length:"20" validate:"required,max=58"`
	Status5      bool       `json:"status5" gorm:"column:status5" true:"5" false:"0" bson:"status5" dynamodbav:"status5" format:"%5s" length:"5" firestore:"status5"`
	CreatedDate5 *time.Time `json:"createdDate5" gorm:"column:createdDate5" bson:"createdDate5" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate5" firestore:"createdDate5" validate:"required"`

	Username6    string     `json:"username6" gorm:"column:username6" bson:"username6" length:"10" dynamodbav:"username6" firestore:"username6" validate:"required,username,max=600"`
	Email6       string     `json:"email6" gorm:"column:email6" bson:"email" dynamodbav:"email6" firestore:"email6" length:"31" validate:"email,max=600"`
	Phone6       string     `json:"phone6" gorm:"column:phone6" bson:"phone6" dynamodbav:"phone6" firestore:"phone6" length:"20" validate:"required,max=68"`
	Status6      bool       `json:"status6" gorm:"column:status6" true:"6" false:"0" bson:"status6" dynamodbav:"status6" format:"%5s" length:"5" firestore:"status6"`
	CreatedDate6 *time.Time `json:"createdDate6" gorm:"column:createdDate6" bson:"createdDate6" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate6" firestore:"createdDate6" validate:"required"`

	Username7    string     `json:"username7" gorm:"column:username7" bson:"username7" length:"10" dynamodbav:"username7" firestore:"username7" validate:"required,username,max=700"`
	Email7       string     `json:"email7" gorm:"column:email7" bson:"email" dynamodbav:"email7" firestore:"email7" length:"31" validate:"email,max=700"`
	Phone7       string     `json:"phone7" gorm:"column:phone7" bson:"phone7" dynamodbav:"phone7" firestore:"phone7" length:"20" validate:"required,max=78"`
	Status7      bool       `json:"status7" gorm:"column:status7" true:"7" false:"0" bson:"status7" dynamodbav:"status7" format:"%5s" length:"5" firestore:"status7"`
	CreatedDate7 *time.Time `json:"createdDate7" gorm:"column:createdDate7" bson:"createdDate7" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate7" firestore:"createdDate7" validate:"required"`

	Username8    string     `json:"username8" gorm:"column:username8" bson:"username8" length:"10" dynamodbav:"username8" firestore:"username8" validate:"required,username,max=800"`
	Email8       string     `json:"email8" gorm:"column:email8" bson:"email" dynamodbav:"email8" firestore:"email8" length:"31" validate:"email,max=800"`
	Phone8       string     `json:"phone8" gorm:"column:phone8" bson:"phone8" dynamodbav:"phone8" firestore:"phone8" length:"20" validate:"required,max=88"`
	Status8      bool       `json:"status8" gorm:"column:status8" true:"8" false:"0" bson:"status8" dynamodbav:"status8" format:"%5s" length:"5" firestore:"status8"`
	CreatedDate8 *time.Time `json:"createdDate8" gorm:"column:createdDate8" bson:"createdDate8" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate8" firestore:"createdDate8" validate:"required"`

	Username9    string     `json:"username9" gorm:"column:username9" bson:"username9" length:"10" dynamodbav:"username9" firestore:"username9" validate:"required,username,max=900"`
	Email9       string     `json:"email9" gorm:"column:email9" bson:"email" dynamodbav:"email9" firestore:"email9" length:"31" validate:"email,max=900"`
	Phone9       string     `json:"phone9" gorm:"column:phone9" bson:"phone9" dynamodbav:"phone9" firestore:"phone9" length:"20" validate:"required,max=98"`
	Status9      bool       `json:"status9" gorm:"column:status9" true:"9" false:"0" bson:"status9" dynamodbav:"status9" format:"%5s" length:"5" firestore:"status9"`
	CreatedDate9 *time.Time `json:"createdDate9" gorm:"column:createdDate9" bson:"createdDate9" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate9" firestore:"createdDate9" validate:"required"`

	Username10    string     `json:"username10" gorm:"column:username10" bson:"username10" length:"10" dynamodbav:"username10" firestore:"username10" validate:"required,username,max=1000"`
	Email10       string     `json:"email10" gorm:"column:email10" bson:"email" dynamodbav:"email10" firestore:"email10" length:"31" validate:"email,max=1000"`
	Phone10       string     `json:"phone10" gorm:"column:phone10" bson:"phone10" dynamodbav:"phone10" firestore:"phone10" length:"20" validate:"required,max=108"`
	Status10      bool       `json:"status10" gorm:"column:status10" true:"10" false:"0" bson:"status10" dynamodbav:"status10" format:"%5s" length:"5" firestore:"status10"`
	CreatedDate10 *time.Time `json:"createdDate10" gorm:"column:createdDate10" bson:"createdDate10" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate10" firestore:"createdDate10" validate:"required"`

	Username11    string     `json:"username11" gorm:"column:username11" bson:"username11" length:"10" dynamodbav:"username11" firestore:"username11" validate:"required,username,max=1100"`
	Email11       string     `json:"email11" gorm:"column:email11" bson:"email" dynamodbav:"email11" firestore:"email11" length:"31" validate:"email,max=1100"`
	Phone11       string     `json:"phone11" gorm:"column:phone11" bson:"phone11" dynamodbav:"phone11" firestore:"phone11" length:"20" validate:"required,max=118"`
	Status11      bool       `json:"status11" gorm:"column:status11" true:"11" false:"0" bson:"status11" dynamodbav:"status11" format:"%5s" length:"5" firestore:"status11"`
	CreatedDate11 *time.Time `json:"createdDate11" gorm:"column:createdDate11" bson:"createdDate11" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate11" firestore:"createdDate11" validate:"required"`

	Username12    string     `json:"username12" gorm:"column:username12" bson:"username12" length:"10" dynamodbav:"username12" firestore:"username12" validate:"required,username,max=1200"`
	Email12       string     `json:"email12" gorm:"column:email12" bson:"email" dynamodbav:"email12" firestore:"email12" length:"31" validate:"email,max=1200"`
	Phone12       string     `json:"phone12" gorm:"column:phone12" bson:"phone12" dynamodbav:"phone12" firestore:"phone12" length:"20" validate:"required,max=128"`
	Status12      bool       `json:"status12" gorm:"column:status12" true:"12" false:"0" bson:"status12" dynamodbav:"status12" format:"%5s" length:"5" firestore:"status12"`
	CreatedDate12 *time.Time `json:"createdDate12" gorm:"column:createdDate12" bson:"createdDate12" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate12" firestore:"createdDate12" validate:"required"`

	Username13    string     `json:"username13" gorm:"column:username13" bson:"username13" length:"10" dynamodbav:"username13" firestore:"username13" validate:"required,username,max=1300"`
	Email13       string     `json:"email13" gorm:"column:email13" bson:"email" dynamodbav:"email13" firestore:"email13" length:"31" validate:"email,max=1300"`
	Phone13       string     `json:"phone13" gorm:"column:phone13" bson:"phone13" dynamodbav:"phone13" firestore:"phone13" length:"20" validate:"required,max=138"`
	Status13      bool       `json:"status13" gorm:"column:status13" true:"13" false:"0" bson:"status13" dynamodbav:"status13" format:"%5s" length:"5" firestore:"status13"`
	CreatedDate13 *time.Time `json:"createdDate13" gorm:"column:createdDate13" bson:"createdDate13" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate13" firestore:"createdDate13" validate:"required"`

	Username14    string     `json:"username14" gorm:"column:username14" bson:"username14" length:"10" dynamodbav:"username14" firestore:"username14" validate:"required,username,max=1400"`
	Email14       string     `json:"email14" gorm:"column:email14" bson:"email" dynamodbav:"email14" firestore:"email14" length:"31" validate:"email,max=1400"`
	Phone14       string     `json:"phone14" gorm:"column:phone14" bson:"phone14" dynamodbav:"phone14" firestore:"phone14" length:"20" validate:"required,max=148"`
	Status14      bool       `json:"status14" gorm:"column:status14" true:"14" false:"0" bson:"status14" dynamodbav:"status14" format:"%5s" length:"5" firestore:"status14"`
	CreatedDate14 *time.Time `json:"createdDate14" gorm:"column:createdDate14" bson:"createdDate14" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate14" firestore:"createdDate14" validate:"required"`
}

func BuildQuery(ctx context.Context) (string, []interface{}) {
	query := "select * from userimport"
	params := make([]interface{}, 0)
	return query, params
}

func GenerateFileName() string {
	fileName := time.Now().Format("20060102150405") + ".csv"
	fullPath := filepath.Join("export", fileName)
	w.DeleteFile(fullPath)
	return fullPath
}
