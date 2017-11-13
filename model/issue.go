package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// IssueEntity structure for database
type IssueEntity struct {
	gorm.Model
	Schedule  int64  `gorm:"column:schedule"`
	StationID int    `gorm:"column:station_id"`
	State     string `gorm:"column:state"`
	CodeTrain string `gorm:"column:code_train"`
}

// ToEntity transform pub/sub issue to the IssueEntity
func (i *Issue) ToEntity() (*IssueEntity, error) {
	schedule, err := time.Parse("02/01/2006 15:04 -0700", i.Schedule)

	if err != nil {
		return nil, err
	}

	return &IssueEntity{
		StationID: i.StationID,
		State:     i.State,
		CodeTrain: i.Code,
		Schedule:  schedule.Unix(),
	}, nil
}

// TableName for IssueEntity
func (e IssueEntity) TableName() string {
	return "issue"
}

// Persist issue in database
func (e *IssueEntity) Persist() error {
	return db.Save(e).Error
}
