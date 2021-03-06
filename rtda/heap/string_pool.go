package heap

import (
	"sync"
	"unicode/utf16"
)

// internedStrings:字符串缓冲池
var internedStrings = map[string]*Object{}
var stringMutex sync.Mutex

func JString(loader *ClassLoader, goStr string) *Object {
	// 读取string时得加锁
	stringMutex.Lock()
	if internedStr, ok := internedStrings[goStr]; ok {
		stringMutex.Unlock()
		return internedStr
	}
	stringMutex.Unlock()

	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars, nil,&Monitor{}}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	// 将字符串对象放入缓冲池中
	stringMutex.Lock()
	if internedStr, ok := internedStrings[goStr]; ok {
		stringMutex.Unlock()
		return internedStr
	}
	internedStrings[goStr] = jStr
	stringMutex.Unlock()
	return jStr
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s)
	return utf16.Encode(runes)
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s)
	return string(runes)
}

// Intern:　将字面量放入到字符串缓存中
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}
