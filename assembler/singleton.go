package assembler

var (
	PostAss PostAssembler
)

func init() {
	PostAss = NewPostAssembler()
}
