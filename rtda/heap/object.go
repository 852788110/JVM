package heap

type Object struct {
	class *Class
	data  interface{} // Slots for Object, []int32 for int[] ...
	extra interface{}
}

// getters & setters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Data() interface{} {
	return self.data
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}
func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}
