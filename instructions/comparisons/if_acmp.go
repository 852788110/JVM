package comparisons

import "jvmgo/jvm/instructions/base"

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

type IF_ACMPNE struct{ base.BranchInstruction }
