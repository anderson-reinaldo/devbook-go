package repositories

import (
	"database/sql"
	"devbook/src/models"
	"errors"
)

type Usuario struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuario {
	return &Usuario{db}
}

// Criar cria um usuario
func (repository *Usuario) Criar(usuario models.Usuario) (int64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Buscar busca todos os usuarios
func (repository *Usuario) Buscar(nameOrNick string) ([]models.Usuario, error) {
	nameOrNick = "%" + nameOrNick + "%"
	stmt, err := repository.db.Prepare("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario

		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorId busca um usuario
func (repository *Usuario) BuscarPorId(id uint64) (models.Usuario, error) {
	stmt, err := repository.db.Prepare("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?")
	if err != nil {
		return models.Usuario{}, err
	}
	defer stmt.Close()

	var usuario models.Usuario
	if err = stmt.QueryRow(id).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Nick,
		&usuario.Email,
		&usuario.CriadoEm,
	); err != nil {
		return models.Usuario{}, errors.New("Usuário não encontrado")
	}

	return usuario, nil
}

// BuscarPorEmail busca um usuario
func (repository *Usuario) BuscarPorEmail(email string) (models.Usuario, error) {
	stmt, err := repository.db.Prepare("SELECT id, senha FROM usuarios WHERE email = ?")
	if err != nil {
		return models.Usuario{}, err
	}
	defer stmt.Close()

	var usuario models.Usuario
	if err = stmt.QueryRow(email).Scan(
		&usuario.ID,
		&usuario.Senha,
	); err != nil {
		return models.Usuario{}, errors.New("Usuário não encontrado")
	}

	return usuario, nil
}

// Atualizar atualiza um usuario
func (repository *Usuario) Atualizar(id uint64, usuario models.Usuario) error {
	stmt, err := repository.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); err != nil {
		return err
	}

	return nil
}

// Deletar deleta um usuario
func (repository *Usuario) Deletar(id uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

// Seguir segue um usuario
func (repository *Usuario) Seguir(usuarioId, seguidorId uint64) error {
	stmt, err := repository.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuarioId, seguidorId); err != nil {
		return err
	}

	return nil
}

// PararDeSeguir para de seguir um usuario
func (repository *Usuario) PararDeSeguir(usuarioId, seguidorId uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuarioId, seguidorId); err != nil {
		return err
	}

	return nil
}

// Seguidores busca os seguidores de um usuario
func (repository *Usuario) Seguidores(usuarioID uint64) ([]models.Usuario, error) {

	stmt, err := repository.db.Prepare(`
		SELECT t2.id, t2.nome,t2.nick,t2.email,t2.criadoEm 
		FROM seguidores AS t1
		INNER JOIN usuarios AS t2 ON t1.seguidor_id = t2.id
		WHERE t1.usuario_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario

		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Seguindo busca os usuarios que um usuario está seguindo
func (repository *Usuario) Seguindo(usuarioID uint64) ([]models.Usuario, error) {
	stmt, err := repository.db.Prepare(`
		SELECT t2.id, t2.nome,t2.nick,t2.email,t2.criadoEm 
		FROM seguidores AS t1
		INNER JOIN usuarios AS t2 ON t1.usuario_id = t2.id
		WHERE t1.seguidor_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario

		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSenha busca a senha atual do usuario
func (repository *Usuario) BuscarSenha(usuarioID uint64) (string, error) {
	stmt, err := repository.db.Prepare("SELECT senha FROM usuarios WHERE id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var usuario models.Usuario
	if err = stmt.QueryRow(usuarioID).Scan(
		&usuario.Senha,
	); err != nil {
		return "", errors.New("Usuário não encontrado")
	}

	return usuario.Senha, nil
}

// AtualizarSenha atualiza a senha de um usuario
func (repository *Usuario) AtualizarSenha(id uint64, senha string) error {
	stmt, err := repository.db.Prepare("UPDATE usuarios SET senha = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(senha, id); err != nil {
		return err
	}

	return nil
}
