package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

type Publicacao struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacao {
	return &Publicacao{db}
}

// Criar cria uma publicacao
func (repository *Publicacao) Criar(publicacao models.Publicacao) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// BuscarPublicacoes busca todas as publicacoes dos usuarios que o usuario que fez a requisição segue e do propio usuario
func (repository *Publicacao) Buscar(usuarioId uint64) ([]models.Publicacao, error) {
	stmt, err := repository.db.Prepare(`
		SELECT  t2.id, t2.titulo, t2.conteudo, t3.nick, t2.curtidas, t2.criadoEm 
		FROM seguidores AS t1
		JOIN publicacoes AS t2 ON t1.usuario_id = t2.autor_id
		JOIN usuarios  AS t3 ON t2.autor_id = t3.id
		WHERE t1.seguidor_id = ? OR t1.usuario_id = ?
		ORDER BY t2.criadoEm DESC, t1.seguidor_id ASC;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(usuarioId, usuarioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publicacoes []models.Publicacao

	for rows.Next() {
		var publicacao models.Publicacao

		if err = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// BuscarPublicacoes busca todas as publicacoes de um usuario
func (repository *Publicacao) BuscarPublicacoesUsuario(usuarioId uint64) ([]models.Publicacao, error) {
	stmt, err := repository.db.Prepare(`
		SELECT  t1.id, t1.titulo, t1.conteudo, t2.nick, t1.curtidas, t1.criadoEm 
		FROM publicacoes AS t1
		JOIN usuarios AS t2 ON t1.autor_id = t2.id
		WHERE t1.autor_id = ?
		ORDER BY t1.criadoEm DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(usuarioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publicacoes []models.Publicacao

	for rows.Next() {
		var publicacao models.Publicacao

		if err = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// BuscarPorId busca uma publicacao por id
func (repository *Publicacao) BuscarPorId(id uint64) (models.Publicacao, error) {
	stmt, err := repository.db.Prepare(`
		SELECT t1.id, t1.autor_id, t1.titulo,t1.conteudo, t2.nick, t1.curtidas, t1.criadoEm 
		FROM publicacoes AS t1
		JOIN usuarios AS t2 ON t1.autor_id = t2.id
		WHERE t1.id = ?
	`)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer stmt.Close()

	var publicacao models.Publicacao

	if err = stmt.QueryRow(id).Scan(
		&publicacao.ID,
		&publicacao.AutorID,
		&publicacao.Titulo,
		&publicacao.Conteudo,
		&publicacao.AutorNick,
		&publicacao.Curtidas,
		&publicacao.CriadoEm,
	); err != nil {
		return models.Publicacao{}, fmt.Errorf("Erro ao buscar a publicação ERROR :: %s", err)
	}

	return publicacao, nil
}

// Atualizar atualiza uma publicacao
func (repository *Publicacao) Atualizar(id uint64, publicacao models.Publicacao) error {
	stmt, err := repository.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(publicacao.Titulo, publicacao.Conteudo, id); err != nil {
		return err
	}

	return nil
}

// Deletar deleta uma publicacao
func (repository *Publicacao) Deletar(id uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM publicacoes WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository *Publicacao) Curtir(usuarioID, publicacaoID uint64) error {
	stmt, err := repository.db.Prepare(`
		INSERT INTO publicacoes_curtidas (usuario_id, publicacao_id)
		VALUES (?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuarioID, publicacaoID); err != nil {
		return err
	}

	_, err = repository.db.Exec(`
		UPDATE publicacoes
		SET curtidas = curtidas + 1
		WHERE id = ?
	`, publicacaoID)

	return err
}

func (repository *Publicacao) Descurtir(usuarioID, publicacaoID uint64) error {
	stmt, err := repository.db.Prepare(`
		DELETE FROM publicacoes_curtidas
		WHERE usuario_id = ? AND publicacao_id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuarioID, publicacaoID); err != nil {
		return err
	}

	_, err = repository.db.Exec(`
		UPDATE publicacoes
		SET curtidas = curtidas - 1
		WHERE id = ? AND curtidas > 0
	`, publicacaoID)

	return err
}
