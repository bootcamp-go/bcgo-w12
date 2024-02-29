package calculator

var (
	ErrMsgInvalidOperator = "invalid operator"
)

// Orchestrator is a function that returns a MathOperation function
// based on the operator string
func Orchestrator(operator MathOperator) (mo MathOperation, err string) {
	switch operator {
	case OperatorSum:
		mo = Add
		// mo = func(i1, i2 int) (int, string) {
		// 	return i1 + i2, ""
		// }
	case OperatorSub:
		mo = Subtract
	case OperatorMul:
		mo = Multiply
	case OperatorDiv:
		mo = Divide
	default:
		err = ErrMsgInvalidOperator
	}
	return
}