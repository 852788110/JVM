package comparisons

import "jvmgo/jvm/instructions/base"

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

type IF_ICMPNE struct{ base.BranchInstruction }

type IF_ICMPLT struct{ base.BranchInstruction }

type IF_ICMPLE struct{ base.BranchInstruction }

type IF_ICMPGT struct{ base.BranchInstruction }

type IF_ICMPGE struct{ base.BranchInstruction }
