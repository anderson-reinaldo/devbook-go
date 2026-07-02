package models

type DadosHome struct {
	Publicacoes       []Publicacao
	IDUsuarioLogado   uint64
	NickUsuarioLogado string
}

type DadosPerfil struct {
	Usuario         Usuario
	Publicacoes     []Publicacao
	QtdSeguidores   int
	QtdSeguindo     int
	IDUsuarioLogado uint64
	Seguindo        bool
	EhDono          bool
}
