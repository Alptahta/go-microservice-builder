package filewriter

var RepositoryCodes = [][]byte{
	[]byte("package repository\n\n"),
	[]byte("import \"context\"\n\n"),
	[]byte("type Repository struct{ctx context.Context}\n\n"),
	[]byte("func CreateRepository(ctx context.Context) *Repository {return &Repository{ctx:ctx}}\n\n"),
}
