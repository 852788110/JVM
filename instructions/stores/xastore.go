package stores

import "jvmgo/jvm/instructions/base"

// Store into reference array
type AASTORE struct{ base.NoOperandsInstruction }

// Store into byte or boolean array
type BASTORE struct{ base.NoOperandsInstruction }

// Store into char array
type CASTORE struct{ base.NoOperandsInstruction }

// Store into double array
type DASTORE struct{ base.NoOperandsInstruction }

// Store into float array
type FASTORE struct{ base.NoOperandsInstruction }

// Store into int array
type IASTORE struct{ base.NoOperandsInstruction }

// Store into long array
type LASTORE struct{ base.NoOperandsInstruction }

// Store into short array
type SASTORE struct{ base.NoOperandsInstruction }
