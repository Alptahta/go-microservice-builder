package filewriter

var RepositoryCodes = [][]byte{
	[]byte("package repository\n\n"),
	[]byte("import \"context\"\n\n"),
	[]byte("type Repository struct {\n\tctx context.Context\n}\n\n"),
	[]byte("func CreateRepository(ctx context.Context) *Repository {\n\treturn &Repository{ctx: ctx}\n}\n\n"),
}
