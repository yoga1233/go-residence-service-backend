package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type ReportRepository interface {
	FindByUserId(userId int) ([]*model.Report, error)
	CreateReport(report *model.Report) error
	UpdateReport(report *model.Report) error
	DeleteReport(reportId int) error
}

type reportRepository struct {
	db *gorm.DB
}

// CreateReport implements ReportRepository.
func (r *reportRepository) CreateReport(report *model.Report) error {
	if err := r.db.Create(&report).Error; err != nil {
		return err
	}
	return nil
}

// DeleteReport implements ReportRepository.
func (r *reportRepository) DeleteReport(reportId int) error {
	if err := r.db.Delete(&model.Report{}, reportId).Error; err != nil {
		return err
	}
	return nil
}

// FindByUserId implements ReportRepository.
func (r *reportRepository) FindByUserId(userId int) ([]*model.Report, error) {

	var reports []*model.Report
	if err := r.db.Where("user_id = ?", userId).Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

// UpdateReport implements ReportRepository.
func (r *reportRepository) UpdateReport(report *model.Report) error {
	if err := r.db.Save(&report).Error; err != nil {
		return err
	}
	return nil
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{
		db: db,
	}
}
