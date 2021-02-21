package rtda

import "jvmgo/jvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
