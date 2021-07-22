package extended

import "jvmgo/jvm/instructions/base"

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }
