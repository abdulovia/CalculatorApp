package calculationService

import "gorm.io/gorm"

type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type calculationRepositoryImpl struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calculationRepositoryImpl{db: db}
}

func (r *calculationRepositoryImpl) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *calculationRepositoryImpl) GetAllCalculations() ([]Calculation, error) {
	var calculations []Calculation
	if err := r.db.Find(&calculations).Error; err != nil {
		return []Calculation{}, err
	}
	return calculations, nil
}

func (r *calculationRepositoryImpl) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation
	if err := r.db.First(&calc, "id = ?", id).Error; err != nil {
		return Calculation{}, err
	}
	return calc, nil
}

func (r *calculationRepositoryImpl) UpdateCalculation(calc Calculation) error {
	return r.db.Save(&calc).Error
}

func (r *calculationRepositoryImpl) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
