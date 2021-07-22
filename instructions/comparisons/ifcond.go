package comparisons

import "jvmgo/jvm/instructions/base"

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

type IFNE struct{ base.BranchInstruction }

type IFLT struct{ base.BranchInstruction }

type IFLE struct{ base.BranchInstruction }

type IFGT struct{ base.BranchInstruction }

type IFGE struct{ base.BranchInstruction }
