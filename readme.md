# devbook-go

Rede social simples (estilo Devbook) escrita em Go, dividida em dois módulos independentes:

- **`server/`** — API REST (Go + `gorilla/mux` + MySQL + JWT). Faz toda a regra de negócio e acesso ao banco.
- **`app/`** — Aplicação web (Go + `gorilla/mux` + HTML/Bootstrap/jQuery). Renderiza as telas e consome a API do `server` via HTTP, autenticando o usuário com cookies seguros (`gorilla/securecookie`).

## Funcionalidades

- Cadastro, login e logout de usuários
- Edição, atualização de senha e exclusão de conta
- Seguir / deixar de seguir outros usuários
- Publicações (posts): criar, editar, excluir, curtir e descurtir
- Listagem de usuários e feed de publicações

## Estrutura

```
app/     -> frontend (views HTML, controllers, router, assets JS/CSS)
server/  -> API (controllers, models, repositories, middlewares, banco MySQL)
```

Cada módulo tem seu próprio `go.mod`, `.env` e `main.go`.

## Banco de dados

Script de criação em `server/sql/sql.sql`. Cria o banco `devbook` com as tabelas:

- `usuarios`
- `seguidores`
- `publicacoes`
- `publicacoes_curtidas`

## Configuração

Cada módulo lê variáveis de um arquivo `.env` na sua raiz.

**`server/.env`**
```
API_PORT=5000
JWT_SECRET=
DB_USUARIO=
DB_SENHA=
DB_NOME=devbook
```

**`app/.env`**
```
PORT=3000
BASEURL_API=http://localhost:5000
HASH_KEY=
BLOCK_KEY=
```

## Rodando

```bash
# 1. Suba o banco e rode server/sql/sql.sql

# 2. API
cd server
go run main.go

# 3. App (frontend)
cd app
go run main.go
```

A aplicação web ficará disponível em `http://localhost:3000` (ou porta definida em `PORT`) e consumirá a API em `BASEURL_API`.
