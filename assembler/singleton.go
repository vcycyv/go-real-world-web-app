package assembler

var (
	BookAsm BookAssembler
)

func init() {
	BookAsm = NewBookAssembler()
}
