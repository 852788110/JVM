package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jvmgo/jvm/rtda/heap"
	"jvmgo/jvm/server/model"
	"net/http"
)

func ParseCode(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("classFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	raw := heap.ParseClass(data)

	class := &model.Class{
		Name:           raw.Name(),
		SuperClassName: raw.GetSuperClass(),
		InterfaceNames: raw.GetInterface(),
	}

	var methods []*model.Method

	for _, m := range raw.Methods() {
		method := &model.Method{
			Name:       m.Name(),
			Descriptor: m.Descriptor(),
		}

		var instructions []string
		for _, i := range m.Code() {
			inst := fmt.Sprintf(i.GetName())
			instructions = append(instructions, inst)
		}

		method.Code = &model.Code{Instructions: instructions}

		methods = append(methods, method)
	}

	class.Methods = methods

	jsonfile, err := json.MarshalIndent(class, " ", " ")
	if err != nil {
		return
	}
	_, _ = w.Write(jsonfile)
}
