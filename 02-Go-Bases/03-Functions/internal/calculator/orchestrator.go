package calculator

// Orchestrator is a function that returns a MathOperation function
// based on the operator string
func Orchestrator(operator MathOperator) (mo MathOperation, err string) {
	switch operator {
	case OperatorSum:
		mo = Add
	case OperatorSub:
		mo = Subtract
	case OperatorMul:
		mo = Multiply
	case OperatorDiv:
		mo = Divide
	default:
		err = "error: invalid operator"
	}
	return
}