package instructionexecutor

import (
	"errors"
	"github.com/MyOrg/Microprocessor/database"
	"github.com/MyOrg/Microprocessor/enum"
	"github.com/MyOrg/Microprocessor/model"
	service "github.com/MyOrg/Microprocessor/operatorfactory"
	"strings"
)

func Execute(instructions []string) ([]model.Register, error) {
	var result []model.Register
	for i, _ := range instructions {
		instruction := strings.Split(instructions[i], " ")
		operatorType := enum.OperatorType(enum.OperatorType_value[instruction[0]])
		operatorHandler, err := service.GetOperatorInstance(operatorType)
		if err != nil {
			return result, errors.New("Invalid instruction")
		}
		operatorHandler.Execute(instruction)
	}
	result = database.GetAllRegister()
	return result, nil
}
