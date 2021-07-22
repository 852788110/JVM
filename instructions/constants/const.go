package constants

import "jvmgo/jvm/instructions/base"

// Push null
type ACONST_NULL struct{ base.NoOperandsInstruction }

// Push double
type DCONST_0 struct{ base.NoOperandsInstruction }

type DCONST_1 struct{ base.NoOperandsInstruction }

// Push float
type FCONST_0 struct{ base.NoOperandsInstruction }

type FCONST_1 struct{ base.NoOperandsInstruction }

type FCONST_2 struct{ base.NoOperandsInstruction }

// Push int constant
type ICONST_M1 struct{ base.NoOperandsInstruction }

type ICONST_0 struct{ base.NoOperandsInstruction }

type ICONST_1 struct{ base.NoOperandsInstruction }

type ICONST_2 struct{ base.NoOperandsInstruction }

type ICONST_3 struct{ base.NoOperandsInstruction }

type ICONST_4 struct{ base.NoOperandsInstruction }

type ICONST_5 struct{ base.NoOperandsInstruction }

// Push long constant
type LCONST_0 struct{ base.NoOperandsInstruction }

type LCONST_1 struct{ base.NoOperandsInstruction }
