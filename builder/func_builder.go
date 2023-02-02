package builder

import (
	"fmt"
	"strconv"
	"strings"
)

type Func struct {
	Name         string
	Params       []string
	Return       []string
	ExportStruct string
	hasResult    bool
}

const (
	ParamsModel = `
	type %s_Params struct {
		%s
	}
	`
	ResultModel = `
	type %s_Result struct {
		%s
	}
	`
	CallerModelNotReturn = `
	func (w *__ws_gen) %s(bindata json.RawMessage) (ret json.RawMessage, err error) {
		var p %s_Params
		ret = nil
		if err = json.Unmarshal(bindata, &p); err != nil {
			return
		}
		err = %s.%s(%s)
		return
	}
	`
	CallerModelHasReturn = `
	func (w *__ws_gen) %s(bindata json.RawMessage) (ret json.RawMessage, err error) {
		var p %s_Params
		var r %s_Result
		if err = json.Unmarshal(bindata, &p); err != nil {
			return
		}
		%s = %s.%s(%s)
		// fetch the result
		%s
		ret, err = json.Marshal(&p)
		return 
	}
	`
)

func params(prefix string, s []string) string {
	var generateCallParams strings.Builder
	generateCallParams.WriteString(prefix + s[0])
	for _, p := range s[1:] {
		generateCallParams.WriteString(", ")
		generateCallParams.WriteString(prefix + p)
	}
	return generateCallParams.String()
}

func NewFunc(name, export string, params []string, returns []string) *Func {
	hasReturn := false
	if returns != nil {
		hasReturn = true
	}
	return &Func{
		Name:         name,
		Params:       params,
		ExportStruct: export,
		Return:       returns,
		hasResult:    hasReturn,
	}
}

func (f *Func) generateParams() string {
	return fmt.Sprintf(ParamsModel, f.Name, strings.Join(f.Params, "\n"))
}

func (f *Func) generateReturn() string {
	return fmt.Sprintf(ResultModel, f.Name, strings.Join(f.Return, "\n"))
}

func (f *Func) tempReturnVar() (string, []string) {
	gen := make([]string, len(f.Return))
	for i := 0; i < len(f.Return); i++ {
		gen = append(gen, "r"+strconv.Itoa(i))
	}
	return strings.Join(gen, ", "), gen
}

func (f *Func) fetchResult(gr []string) string {
	var gen strings.Builder
	gen.WriteString("r." + f.Return[0] + " = " + gr[0])
	for i, p := range f.Return[1:] {
		gen.WriteString(", ")
		gen.WriteString("r." + p + " = " + gr[i+1])
	}
	return gen.String()

}

func (f *Func) generateCaller() string {
	if f.hasResult {
		retGen, rets := f.tempReturnVar()
		return fmt.Sprintf(CallerModelHasReturn,
			f.Name,
			f.Name,
			f.Name,
			retGen,
			f.ExportStruct,
			f.Name,
			f.fetchResult(rets),
			params("p.", f.Params))
	}
	return fmt.Sprintf(CallerModelNotReturn,
		f.Name,
		f.Name,
		f.ExportStruct,
		f.Name,
		params("p.", f.Params))
}
func (f *Func) String() string {
	if f.hasResult {
		return f.generateParams() + "\n" + f.generateReturn() + "\n" + f.generateCaller()
	}
	return f.generateParams() + "\n" + f.generateCaller()
}
