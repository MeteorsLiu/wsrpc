package builder

type ExportStruct struct {
	A string
}
type FuncA_Args struct {
	a string
}
type FuncB_Args struct {
	b string
}

type FuncC_Args struct {
	c string
}

func (e *ExportStruct) FuncA(a *FuncA_Args) {

}

func (e *ExportStruct) FuncB(b *FuncB_Args) {

}

func (e *ExportStruct) FuncC(c *FuncC_Args) {

}

func (e ExportStruct) FuncNotPtr(c *FuncC_Args) {

}
