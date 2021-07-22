package conversions

import "jvmgo/jvm/instructions/base"

// Convert int to byte
type I2B struct{ base.NoOperandsInstruction }

// Convert int to char
type I2C struct{ base.NoOperandsInstruction }

// Convert int to short
type I2S struct{ base.NoOperandsInstruction }

// Convert int to long
type I2L struct{ base.NoOperandsInstruction }

// Convert int to float
type I2F struct{ base.NoOperandsInstruction }

// Convert int to double
type I2D struct{ base.NoOperandsInstruction }
