package instructions

import "fmt"
import "jvmgo/jvm/instructions/base"
import . "jvmgo/jvm/instructions/comparisons"
import . "jvmgo/jvm/instructions/constants"
import . "jvmgo/jvm/instructions/control"
import . "jvmgo/jvm/instructions/conversions"
import . "jvmgo/jvm/instructions/extended"
import . "jvmgo/jvm/instructions/loads"
import . "jvmgo/jvm/instructions/math"
import . "jvmgo/jvm/instructions/references"
import . "jvmgo/jvm/instructions/stack"
import . "jvmgo/jvm/instructions/stores"

// NoOperandsInstruction singletons
var (
	nop          = &NOP{NoOperandsInstruction: base.NoOperandsInstruction{Name: "nop"}}
	aconst_null  = &ACONST_NULL{NoOperandsInstruction: base.NoOperandsInstruction{Name: "aconst_null"}}
	iconst_m1    = &ICONST_M1{NoOperandsInstruction: base.NoOperandsInstruction{Name: "iconst_m1"}}
	iconst_0     = &ICONST_0{base.NoOperandsInstruction{Name: "iconst 0"}}
	iconst_1     = &ICONST_1{base.NoOperandsInstruction{Name: "iconst 1"}}
	iconst_2     = &ICONST_2{base.NoOperandsInstruction{Name: "iconst 2"}}
	iconst_3     = &ICONST_3{base.NoOperandsInstruction{Name: "iconst 3"}}
	iconst_4     = &ICONST_4{base.NoOperandsInstruction{Name: "iconst 4"}}
	iconst_5     = &ICONST_5{base.NoOperandsInstruction{Name: "iconst 5"}}
	lconst_0     = &LCONST_0{base.NoOperandsInstruction{Name: "lconst 0"}}
	lconst_1     = &LCONST_1{base.NoOperandsInstruction{Name: "lconst 1"}}
	fconst_0     = &FCONST_0{base.NoOperandsInstruction{Name: "fconst 0"}}
	fconst_1     = &FCONST_1{base.NoOperandsInstruction{Name: "fconst 1"}}
	fconst_2     = &FCONST_2{base.NoOperandsInstruction{Name: "fconst 2"}}
	dconst_0     = &DCONST_0{base.NoOperandsInstruction{Name: "dconst 0"}}
	dconst_1     = &DCONST_1{base.NoOperandsInstruction{Name: "dconst 1"}}
	iload_0      = &ILOAD_0{base.NoOperandsInstruction{Name: "iload 0"}}
	iload_1      = &ILOAD_1{base.NoOperandsInstruction{Name: "iload 1"}}
	iload_2      = &ILOAD_2{base.NoOperandsInstruction{Name: "iload 2"}}
	iload_3      = &ILOAD_3{base.NoOperandsInstruction{Name: "iload 3"}}
	lload_0      = &LLOAD_0{base.NoOperandsInstruction{Name: "lload 0"}}
	lload_1      = &LLOAD_1{base.NoOperandsInstruction{Name: "lload 1"}}
	lload_2      = &LLOAD_2{base.NoOperandsInstruction{Name: "lload 2"}}
	lload_3      = &LLOAD_3{base.NoOperandsInstruction{Name: "lload 3"}}
	fload_0      = &FLOAD_0{base.NoOperandsInstruction{Name: "fload 2"}}
	fload_1      = &FLOAD_1{base.NoOperandsInstruction{Name: "fload 1"}}
	fload_2      = &FLOAD_2{base.NoOperandsInstruction{Name: "fload 2"}}
	fload_3      = &FLOAD_3{base.NoOperandsInstruction{Name: "fload 3"}}
	dload_0      = &DLOAD_0{base.NoOperandsInstruction{Name: "dload 0"}}
	dload_1      = &DLOAD_1{base.NoOperandsInstruction{Name: "dload 1"}}
	dload_2      = &DLOAD_2{base.NoOperandsInstruction{Name: "dload 2"}}
	dload_3      = &DLOAD_3{base.NoOperandsInstruction{Name: "dload 3"}}
	aload_0      = &ALOAD_0{base.NoOperandsInstruction{Name: "aload 0"}}
	aload_1      = &ALOAD_1{base.NoOperandsInstruction{Name: "aload 1"}}
	aload_2      = &ALOAD_2{base.NoOperandsInstruction{Name: "aload 2"}}
	aload_3      = &ALOAD_3{base.NoOperandsInstruction{Name: "aload 3"}}
	iaload       = &IALOAD{base.NoOperandsInstruction{Name: "iaload"}}
	laload       = &LALOAD{base.NoOperandsInstruction{Name: "laload"}}
	faload       = &FALOAD{base.NoOperandsInstruction{Name: "faload"}}
	daload       = &DALOAD{base.NoOperandsInstruction{Name: "daload"}}
	aaload       = &AALOAD{base.NoOperandsInstruction{Name: "aaload"}}
	baload       = &BALOAD{base.NoOperandsInstruction{Name: "baload"}}
	caload       = &CALOAD{base.NoOperandsInstruction{Name: "caload"}}
	saload       = &SALOAD{base.NoOperandsInstruction{Name: "saload"}}
	istore_0     = &ISTORE_0{base.NoOperandsInstruction{Name: "istore 0"}}
	istore_1     = &ISTORE_1{base.NoOperandsInstruction{Name: "istore 1"}}
	istore_2     = &ISTORE_2{base.NoOperandsInstruction{Name: "istore 2"}}
	istore_3     = &ISTORE_3{base.NoOperandsInstruction{Name: "istore 3"}}
	lstore_0     = &LSTORE_0{base.NoOperandsInstruction{Name: "lstore 0"}}
	lstore_1     = &LSTORE_1{base.NoOperandsInstruction{Name: "lstore 1"}}
	lstore_2     = &LSTORE_2{base.NoOperandsInstruction{Name: "lstore 2"}}
	lstore_3     = &LSTORE_3{base.NoOperandsInstruction{Name: "lstore 3"}}
	fstore_0     = &FSTORE_0{base.NoOperandsInstruction{Name: "fstore 0"}}
	fstore_1     = &FSTORE_1{base.NoOperandsInstruction{Name: "fstore 1"}}
	fstore_2     = &FSTORE_2{base.NoOperandsInstruction{Name: "fstore 2"}}
	fstore_3     = &FSTORE_3{base.NoOperandsInstruction{Name: "fstore 3"}}
	dstore_0     = &DSTORE_0{base.NoOperandsInstruction{Name: "dstore 0"}}
	dstore_1     = &DSTORE_1{base.NoOperandsInstruction{Name: "dstore 1"}}
	dstore_2     = &DSTORE_2{base.NoOperandsInstruction{Name: "dstore 2"}}
	dstore_3     = &DSTORE_3{base.NoOperandsInstruction{Name: "dstore 3"}}
	astore_0     = &ASTORE_0{base.NoOperandsInstruction{Name: "astore 0"}}
	astore_1     = &ASTORE_1{base.NoOperandsInstruction{Name: "astore 1"}}
	astore_2     = &ASTORE_2{base.NoOperandsInstruction{Name: "astore 2"}}
	astore_3     = &ASTORE_3{base.NoOperandsInstruction{Name: "astore 3"}}
	iastore      = &IASTORE{base.NoOperandsInstruction{Name: "iastore"}}
	lastore      = &LASTORE{base.NoOperandsInstruction{Name: "lastore"}}
	fastore      = &FASTORE{base.NoOperandsInstruction{Name: "fastore"}}
	dastore      = &DASTORE{base.NoOperandsInstruction{Name: "dastore"}}
	aastore      = &AASTORE{base.NoOperandsInstruction{Name: "aastore"}}
	bastore      = &BASTORE{base.NoOperandsInstruction{Name: "bastore"}}
	castore      = &CASTORE{base.NoOperandsInstruction{Name: "castore"}}
	sastore      = &SASTORE{base.NoOperandsInstruction{Name: "sastore"}}
	pop          = &POP{base.NoOperandsInstruction{Name: "pop"}}
	pop2         = &POP2{base.NoOperandsInstruction{Name: "pop2"}}
	dup          = &DUP{base.NoOperandsInstruction{Name: "dup"}}
	dup_x1       = &DUP_X1{base.NoOperandsInstruction{Name: "dup x1"}}
	dup_x2       = &DUP_X2{base.NoOperandsInstruction{Name: "dup x2"}}
	dup2         = &DUP2{base.NoOperandsInstruction{Name: "dup2"}}
	dup2_x1      = &DUP2_X1{base.NoOperandsInstruction{Name: "dup2 x1"}}
	dup2_x2      = &DUP2_X2{base.NoOperandsInstruction{Name: "dup2 x2"}}
	swap         = &SWAP{base.NoOperandsInstruction{Name: "swap"}}
	iadd         = &IADD{base.NoOperandsInstruction{Name: "iadd"}}
	ladd         = &LADD{base.NoOperandsInstruction{Name: "ladd"}}
	fadd         = &FADD{base.NoOperandsInstruction{Name: "fadd"}}
	dadd         = &DADD{base.NoOperandsInstruction{Name: "dadd"}}
	isub         = &ISUB{base.NoOperandsInstruction{Name: "isub"}}
	lsub         = &LSUB{base.NoOperandsInstruction{Name: "lsub"}}
	fsub         = &FSUB{base.NoOperandsInstruction{Name: "fsub"}}
	dsub         = &DSUB{base.NoOperandsInstruction{Name: "dsub"}}
	imul         = &IMUL{base.NoOperandsInstruction{Name: "imul"}}
	lmul         = &LMUL{base.NoOperandsInstruction{Name: "lmul"}}
	fmul         = &FMUL{base.NoOperandsInstruction{Name: "fmul"}}
	dmul         = &DMUL{base.NoOperandsInstruction{Name: "dmul"}}
	idiv         = &IDIV{base.NoOperandsInstruction{Name: "idiv"}}
	ldiv         = &LDIV{base.NoOperandsInstruction{Name: "ldiv"}}
	fdiv         = &FDIV{base.NoOperandsInstruction{Name: "fdiv"}}
	ddiv         = &DDIV{base.NoOperandsInstruction{Name: "ddiv"}}
	irem         = &IREM{base.NoOperandsInstruction{Name: "irem"}}
	lrem         = &LREM{base.NoOperandsInstruction{Name: "lrem"}}
	frem         = &FREM{base.NoOperandsInstruction{Name: "frem"}}
	drem         = &DREM{base.NoOperandsInstruction{Name: "drem"}}
	ineg         = &INEG{base.NoOperandsInstruction{Name: "ineg"}}
	lneg         = &LNEG{base.NoOperandsInstruction{Name: "lneg"}}
	fneg         = &FNEG{base.NoOperandsInstruction{Name: "fneg"}}
	dneg         = &DNEG{base.NoOperandsInstruction{Name: "dneg"}}
	ishl         = &ISHL{base.NoOperandsInstruction{Name: "ishl"}}
	lshl         = &LSHL{base.NoOperandsInstruction{Name: "lshl"}}
	ishr         = &ISHR{base.NoOperandsInstruction{Name: "ishr"}}
	lshr         = &LSHR{base.NoOperandsInstruction{Name: "lshr"}}
	iushr        = &IUSHR{base.NoOperandsInstruction{Name: "iushr"}}
	lushr        = &LUSHR{base.NoOperandsInstruction{Name: "lushr"}}
	iand         = &IAND{base.NoOperandsInstruction{Name: "iand"}}
	land         = &LAND{base.NoOperandsInstruction{Name: "land"}}
	ior          = &IOR{base.NoOperandsInstruction{Name: "ior"}}
	lor          = &LOR{base.NoOperandsInstruction{Name: "lor"}}
	ixor         = &IXOR{base.NoOperandsInstruction{Name: "ixor"}}
	lxor         = &LXOR{base.NoOperandsInstruction{Name: "lxor"}}
	i2l          = &I2L{base.NoOperandsInstruction{Name: "i2l"}}
	i2f          = &I2F{base.NoOperandsInstruction{Name: "i2f"}}
	i2d          = &I2D{base.NoOperandsInstruction{Name: "i2d"}}
	l2i          = &L2I{base.NoOperandsInstruction{Name: "l2i"}}
	l2f          = &L2F{base.NoOperandsInstruction{Name: "l2f"}}
	l2d          = &L2D{base.NoOperandsInstruction{Name: "l2d"}}
	f2i          = &F2I{base.NoOperandsInstruction{Name: "l2i"}}
	f2l          = &F2L{base.NoOperandsInstruction{Name: "f2l"}}
	f2d          = &F2D{base.NoOperandsInstruction{Name: "f2d"}}
	d2i          = &D2I{base.NoOperandsInstruction{Name: "d2i"}}
	d2l          = &D2L{base.NoOperandsInstruction{Name: "d2l"}}
	d2f          = &D2F{base.NoOperandsInstruction{Name: "d2f"}}
	i2b          = &I2B{base.NoOperandsInstruction{Name: "i2b"}}
	i2c          = &I2C{base.NoOperandsInstruction{Name: "i2c"}}
	i2s          = &I2S{base.NoOperandsInstruction{Name: "i2s"}}
	lcmp         = &LCMP{base.NoOperandsInstruction{Name: "lcmp"}}
	fcmpl        = &FCMPL{base.NoOperandsInstruction{Name: "fcmpl"}}
	fcmpg        = &FCMPG{base.NoOperandsInstruction{Name: "fcmpg"}}
	dcmpl        = &DCMPL{base.NoOperandsInstruction{Name: "dcmpl"}}
	dcmpg        = &DCMPG{base.NoOperandsInstruction{Name: "dcmpg"}}
	ireturn      = &IRETURN{base.NoOperandsInstruction{Name: "ireturn"}}
	lreturn      = &LRETURN{base.NoOperandsInstruction{Name: "lreturn"}}
	freturn      = &FRETURN{base.NoOperandsInstruction{Name: "freturn"}}
	dreturn      = &DRETURN{base.NoOperandsInstruction{Name: "dreturn"}}
	areturn      = &ARETURN{base.NoOperandsInstruction{Name: "areturn"}}
	_return      = &RETURN{base.NoOperandsInstruction{Name: "return"}}
	arraylength  = &ARRAY_LENGTH{base.NoOperandsInstruction{Name: "array length"}}
	athrow       = &ATHROW{base.NoOperandsInstruction{Name: "athrow"}}
	monitorenter = &MONITOR_ENTER{base.NoOperandsInstruction{Name: "monitor enter"}}
	monitorexit  = &MONITOR_EXIT{base.NoOperandsInstruction{Name: "monitor exit"}}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return &BIPUSH{Name: "bipush"}
	case 0x11:
		return &SIPUSH{Name: "sipush"}
	case 0x12:
		return &LDC{base.Index8Instruction{Name: "ldc"}}
	case 0x13:
		return &LDC_W{base.Index16Instruction{Name: "ldc_w"}}
	case 0x14:
		return &LDC2_W{base.Index16Instruction{Name: "ldc2_w"}}
	case 0x15:
		return &ILOAD{base.Index8Instruction{Name: "iload"}}
	case 0x16:
		return &LLOAD{base.Index8Instruction{Name: "lload"}}
	case 0x17:
		return &FLOAD{base.Index8Instruction{Name: "fload"}}
	case 0x18:
		return &DLOAD{base.Index8Instruction{Name: "dload"}}
	case 0x19:
		return &ALOAD{base.Index8Instruction{Name: "aload"}}
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x2e:
		return iaload
	case 0x2f:
		return laload
	case 0x30:
		return faload
	case 0x31:
		return daload
	case 0x32:
		return aaload
	case 0x33:
		return baload
	case 0x34:
		return caload
	case 0x35:
		return saload
	case 0x36:
		return &ISTORE{base.Index8Instruction{Name: "istore"}}
	case 0x37:
		return &LSTORE{base.Index8Instruction{Name: "lstore"}}
	case 0x38:
		return &FSTORE{base.Index8Instruction{Name: "fstore"}}
	case 0x39:
		return &DSTORE{base.Index8Instruction{Name: "dstore"}}
	case 0x3a:
		return &ASTORE{base.Index8Instruction{Name: "astore"}}
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x4f:
		return iastore
	case 0x50:
		return lastore
	case 0x51:
		return fastore
	case 0x52:
		return dastore
	case 0x53:
		return aastore
	case 0x54:
		return bastore
	case 0x55:
		return castore
	case 0x56:
		return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &IINC{Name: "iinc"}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return &IFEQ{base.BranchInstruction{Name: "ifeq"}}
	case 0x9a:
		return &IFNE{base.BranchInstruction{Name: "ifne"}}
	case 0x9b:
		return &IFLT{base.BranchInstruction{Name: "iflt"}}
	case 0x9c:
		return &IFGE{base.BranchInstruction{Name: "ifge"}}
	case 0x9d:
		return &IFGT{base.BranchInstruction{Name: "ifgt"}}
	case 0x9e:
		return &IFLE{base.BranchInstruction{Name: "ifle"}}
	case 0x9f:
		return &IF_ICMPEQ{base.BranchInstruction{Name: "if icmpeq"}}
	case 0xa0:
		return &IF_ICMPNE{base.BranchInstruction{Name: "if icmpne"}}
	case 0xa1:
		return &IF_ICMPLT{base.BranchInstruction{Name: "if icmplt"}}
	case 0xa2:
		return &IF_ICMPGE{base.BranchInstruction{Name: "if icmpge"}}
	case 0xa3:
		return &IF_ICMPGT{base.BranchInstruction{Name: "if icmpgt"}}
	case 0xa4:
		return &IF_ICMPLE{base.BranchInstruction{Name: "if icmple"}}
	case 0xa5:
		return &IF_ACMPEQ{base.BranchInstruction{Name: "if acmpeq"}}
	case 0xa6:
		return &IF_ACMPNE{base.BranchInstruction{Name: "if acmpne"}}
	case 0xa7:
		return &GOTO{base.BranchInstruction{Name: "goto"}}
	// case 0xa8:
	// 	return &JSR{}
	// case 0xa9:
	// 	return &RET{}
	case 0xaa:
		return &TABLE_SWITCH{Name: "table switch"}
	case 0xab:
		return &LOOKUP_SWITCH{Name: "lookup switch"}
	case 0xac:
		return ireturn
	case 0xad:
		return lreturn
	case 0xae:
		return freturn
	case 0xaf:
		return dreturn
	case 0xb0:
		return areturn
	case 0xb1:
		return _return
	case 0xb2:
		return &GET_STATIC{base.Index16Instruction{Name: "get static"}}
	case 0xb3:
		return &PUT_STATIC{base.Index16Instruction{Name: "put static"}}
	case 0xb4:
		return &GET_FIELD{base.Index16Instruction{Name: "get field"}}
	case 0xb5:
		return &PUT_FIELD{base.Index16Instruction{Name: "put field"}}
	case 0xb6:
		return &INVOKE_VIRTUAL{base.Index16Instruction{Name: "invoke virtual"}}
	case 0xb7:
		return &INVOKE_SPECIAL{Index16Instruction: base.Index16Instruction{Name: "invoke special"}}
	case 0xb8:
		return &INVOKE_STATIC{Index16Instruction: base.Index16Instruction{Name: "invoke static"}}
	case 0xb9:
		return &INVOKE_INTERFACE{Name: "invoke interface"}
	// case 0xba:
	// 	return &INVOKE_DYNAMIC{}
	case 0xbb:
		return &NEW{Index16Instruction: base.Index16Instruction{Name: "new"}}
	case 0xbc:
		return &NEW_ARRAY{Name: "new array"}
	case 0xbd:
		return &ANEW_ARRAY{Index16Instruction: base.Index16Instruction{Name: "anew array"}}
	case 0xbe:
		return arraylength
	case 0xbf:
		return athrow
	case 0xc0:
		return &CHECK_CAST{Index16Instruction: base.Index16Instruction{Name: "check cast"}}
	case 0xc1:
		return &INSTANCE_OF{Index16Instruction: base.Index16Instruction{Name: "instance of"}}
	case 0xc2:
		return monitorenter
	case 0xc3:
		return monitorexit
	case 0xc4:
		return &WIDE{Name: "wide"}
	case 0xc5:
		return &MULTI_ANEW_ARRAY{Name: "multi anew array"}
	case 0xc6:
		return &IFNULL{BranchInstruction: base.BranchInstruction{Name: "ifnull"}}
	case 0xc7:
		return &IFNONNULL{BranchInstruction: base.BranchInstruction{Name: "ifnonnull"}}
	case 0xc8:
		return &GOTO_W{Name: "goto_w"}
	// case 0xc9:
	// 	return &JSR_W{}
	// case 0xca: breakpoint
	// case 0xff: impdep2
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
