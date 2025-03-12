package service

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
)

type ReportService interface {
	FindByUserId(userId int) ([]*model.Report, error)
	CreateReport(report *model.Report) error
	UpdateReport(report *model.Report) error
	DeleteReport(idReport int) error
}

type reportService struct {
	reportRepository repositories.ReportRepository
}

// CreateReport implements ReportService.
func (r *reportService) CreateReport(report *model.Report) error {
	err := r.reportRepository.CreateReport(report)
	if err != nil {
		return err
	}
	return nil
}

// DeleteReport implements ReportService.
func (r *reportService) DeleteReport(idReport int) error {
	err := r.reportRepository.DeleteReport(idReport)
	if err != nil {
		return err
	}
	return nil
}

// FindByUserId implements ReportService.
func (r *reportService) FindByUserId(userId int) ([]*model.Report, error) {
	result, err := r.reportRepository.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateReport implements ReportService.
func (r *reportService) UpdateReport(report *model.Report) error {
	err := r.reportRepository.UpdateReport(report)
	if err != nil {
		return err
	}
	return nil
}

func NewReportService(reportRepo repositories.ReportRepository) ReportService {
	return &reportService{
		reportRepository: reportRepo,
	}

}
