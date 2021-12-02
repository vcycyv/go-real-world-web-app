package assembler

var (
	BookAss BookAssembler
)

func init() {
	BookAss = NewBookAssembler()
}
