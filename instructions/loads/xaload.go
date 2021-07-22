package loads

import (
	"jvmgo/jvm/instructions/base"
)

// Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }

// Load char from array
type CALOAD struct{ base.NoOperandsInstruction }

// Load double from array
type DALOAD struct{ base.NoOperandsInstruction }

// Load float from array
type FALOAD struct{ base.NoOperandsInstruction }

// Load int from array
type IALOAD struct{ base.NoOperandsInstruction }

// Load long from array
type LALOAD struct{ base.NoOperandsInstruction }

// Load short from array
type SALOAD struct{ base.NoOperandsInstruction }
