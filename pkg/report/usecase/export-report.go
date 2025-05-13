package usecase

func (u reportUsecase) ExportReport(QFMainID, qfSubplanId *string) (string, error) {
	return u.ReportRepository.ExportReport(QFMainID, qfSubplanId)
}