package service

import "github.com/MyOrg/Microprocessor/database"

func init() {
	RstOperatorInstance = &rstOperator{}
}

type rstOperator struct {
}

var RstOperatorInstance *rstOperator

func (d rstOperator) Execute(instruction []string) error {
	database.Reset()
	return nil
}
