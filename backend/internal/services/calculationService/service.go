package calculationService

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vjeantet/govaluate"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id string, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calculationServiceImpl struct {
	repo CalculationRepository
}

func NewCalculationService(repo CalculationRepository) CalculationService {
	return &calculationServiceImpl{repo: repo}
}

func (s *calculationServiceImpl) evaluateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	res, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", res), nil
}

func (s *calculationServiceImpl) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.evaluateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc := Calculation{
		ID:         uuid.NewString(), // gen new uuid
		Expression: expression,
		Result:     result,
	}
	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, nil
}

func (s *calculationServiceImpl) GetAllCalculations() ([]Calculation, error) {
	return s.repo.GetAllCalculations()
}

func (s *calculationServiceImpl) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calculationServiceImpl) UpdateCalculation(id string, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.evaluateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc.Expression = expression
	calc.Result = result
	if err = s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, nil
}

func (s *calculationServiceImpl) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}
