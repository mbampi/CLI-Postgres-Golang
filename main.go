package main

import (
	"amazon/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/thatisuday/commando"
)

func main() {
	// Estabelece conexão com o Postgres
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbCred := &models.DBCredentials{
		User:     dbUser,
		Password: dbPass,
		Name:     "amazon_fbd",
		Address:  "127.0.0.1",
		Port:     "5432",
	}
	connStr := dbCred.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Configurações do CLI
	commando.
		SetExecutableName("amazon").
		SetVersion("v1.0.0").
		SetDescription("Programa de linha de comando Amazon FBD. Estabelecendo conexões com o banco de dados Postgres.")

	// Define comandos CLI
	commando.
		Register("ping").
		SetShortDescription("pinga o banco de dados para verificar se a conexão está OK").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("Pinging Database...")
			fmt.Println("Ping")
			err := db.Ping()
			if err != nil {
				fmt.Print("Error: ", err.Error())
			} else {
				fmt.Println("Pong")
			}
		})

	commando.
		Register("usuarios").
		SetShortDescription("Pega todos os usuários").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Usuários...\n\n")
			rows, err := db.Query("SELECT * FROM usuario")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("codUsuario | email | nome | senha")
			for rows.Next() {
				usuario := models.Usuario{}
				err = rows.Scan(&usuario.CodUsuario, &usuario.Email, &usuario.Nome, &usuario.Senha)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(usuario)
			}
		})

	commando.
		Register("produtos").
		SetShortDescription("Pega todos os produtos").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Produtos...\n\n")
			rows, err := db.Query("SELECT * FROM produto")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("CodProduto | Nome | Descricao | Foto | Marca")
			for rows.Next() {
				produto := models.Produto{}
				err = rows.Scan(&produto.CodProduto, &produto.Nome, &produto.Descricao, &produto.Foto, &produto.Marca)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(produto)
			}
		})

	// --------------------------------------------
	// 					CONSULTAS
	// --------------------------------------------

	// (Numero de Produtos vendidos por anunciante)
	commando.
		Register("num_prod_anunciante").
		SetShortDescription("Numero de produtos anunciados por cada anunciante").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("num_prod_anunciante")
			rows, err := db.Query("SELECT usuario.nome, COUNT(codProdFornecido) AS num_produtos " +
				"FROM anunciante JOIN usuario USING(codUsuario) JOIN ProdutoFornecido USING(codanunciante) JOIN Produto USING(codProduto) " +
				"GROUP BY usuario.nome")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Anunciante | Num Produtos")
			for rows.Next() {
				var nomeAnunciante, nomeProduto string
				err = rows.Scan(&nomeAnunciante, &nomeProduto)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(nomeAnunciante, "|", nomeProduto)
			}
		})

	// (Produtos vendidos por mais de um anunciante)
	commando.
		Register("produtos_num_fornecedores").
		AddArgument("num", "Numero de fornecedores minimo", "1").
		SetShortDescription("Pega todos os produtos que tenham pelo menos X anunciantes").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			num := args["num"].Value
			fmt.Println("produtos_num_fornecedores", num)
			query := "SELECT COUNT(DISTINCT codAnunciante) AS fornecedores, produto.nome " +
				"FROM Anunciante NATURAL JOIN ProdutoFornecido NATURAL JOIN Produto " +
				"GROUP BY produto.nome HAVING COUNT(DISTINCT codAnunciante) >= " + num
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Nome Produto | Num Fornecedores")
			for rows.Next() {
				var numFornecedores int
				var nomeProduto string
				err = rows.Scan(&numFornecedores, &nomeProduto)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(nomeProduto, "|", numFornecedores)
			}
		})

	// (Anunciantes que vendem romance)
	commando.
		Register("anunciantes_vendem_categoria").
		AddArgument("categoria", "Categoria", "Romance").
		SetShortDescription("Anunciantes que vendem categoria X").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			categoria := args["categoria"].Value
			fmt.Println("anunciantes_vendem_categoria", categoria)
			query := "SELECT DISTINCT codAnunciante " +
				"FROM anunciante NATURAL JOIN produtoFornecido NATURAL JOIN produto NATURAL JOIN produto_categoria " +
				"WHERE codCategoria IN (" +
				"SELECT DISTINCT codCategoria " +
				"FROM categoria " +
				"WHERE categoria.nome = '" + categoria + "'" +
				");"

			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Cod Anunciante")
			for rows.Next() {
				var codAnunciante int
				err = rows.Scan(&codAnunciante)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(codAnunciante)
			}
		})

	// (Anunciantes que vendem mangás e brinquedos)
	commando.
		Register("anunciantes_duas_categorias").
		AddArgument("categoria1", "Categoria 1", "Mangas").
		AddArgument("categoria2", "Categoria 2", "Brinquedos").
		SetShortDescription("Anunciantes que vendem categoria X e Y").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			categoria1 := args["categoria1"].Value
			categoria2 := args["categoria2"].Value
			fmt.Println("anunciantes_duas_categorias", categoria1, categoria2)
			query := "SELECT DISTINCT usuario.nome " +
				"FROM usuario JOIN anunciante USING(codUsuario) JOIN produtofornecido USING (codAnunciante) JOIN produto USING(codProduto) JOIN produto_categoria USING (codProduto) " +
				"WHERE codcategoria IN (" +
				"SELECT DISTINCT codcategoria " +
				"FROM categoria " +
				"WHERE categoria.nome = '" + categoria1 + "'" +
				") " +
				"INTERSECT " +
				"SELECT DISTINCT usuario.nome " +
				"FROM usuario JOIN anunciante USING(codUsuario) JOIN produtofornecido USING (codAnunciante) JOIN produto USING(codProduto) JOIN produto_categoria USING (codProduto) " +
				"WHERE codcategoria IN (" +
				"SELECT DISTINCT codcategoria " +
				"FROM categoria " +
				"WHERE categoria.nome = '" + categoria2 + "'" +
				")"

			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Nome Anunciante")
			for rows.Next() {
				var nomeAnunciante string
				err = rows.Scan(&nomeAnunciante)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(nomeAnunciante)
			}
		})

	// (O Nome dos anunciantes que anunciaram exatamente os mesmos produtos que algum outro anunciante)
	commando.
		Register("anunciantes_mesmos_produtos").
		AddArgument("anunciante", "Anunciantes que tenham exatamente os mesmos produtos que este anunciante", "Livros Online").
		SetShortDescription("Todas os anunciantes que tem exatamente os mesmos produtos que o anunciante X").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			nomeAnunciante := args["anunciante"].Value
			fmt.Println("anunciantes_mesmos_produtos", nomeAnunciante)
			query := "SELECT usuario.nome " +
				"FROM usuario NATURAL JOIN anunciante a1 " +
				"WHERE usuario.nome <> '" + nomeAnunciante + "' " +
				"AND NOT EXISTS (" +
				"SELECT * " +
				"FROM produto NATURAL JOIN produtofornecido JOIN anunciante USING (codanunciante) JOIN usuario USING (codUsuario) " +
				"WHERE codProduto NOT IN (" +
				"SELECT DISTINCT codProduto " +
				"FROM produtoFornecido " +
				"WHERE codAnunciante = a1.codAnunciante" +
				") " +
				"AND usuario.nome = '" + nomeAnunciante + "'" +
				") " +
				"AND NOT EXISTS (" +
				"SELECT * " +
				"FROM produtoFornecido " +
				"WHERE codProduto NOT IN (" +
				"SELECT codProduto " +
				"FROM produto NATURAL JOIN produtofornecido JOIN anunciante USING (codanunciante) JOIN usuario USING (codUsuario) " +
				"WHERE usuario.nome = '" + nomeAnunciante + "'" +
				") " +
				"AND codAnunciante = a1.codAnunciante" +
				")"
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Nome Anunciante")
			for rows.Next() {
				var nomeAnunciante string
				err = rows.Scan(&nomeAnunciante)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(nomeAnunciante)
			}
		})

	// (Usuários prime que fizeram compras)
	commando.
		Register("usuario_prime_pedido").
		SetShortDescription("Usuarios Amazon Prime que realizaram pedido").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("anunciantes_mesmos_produtos")
			query := "SELECT DISTINCT cpf, codUsuario, nome " +
				"FROM NotaFiscal " +
				"WHERE prime = true"
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("CPF | Cod Usuário | Nome Comprador")
			for rows.Next() {
				var codUsuario int
				var cpf, nomeComprador string
				err = rows.Scan(&cpf, &codUsuario, &nomeComprador)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(cpf, "|", codUsuario, "|", nomeComprador)
			}
		})

	// (Valor total de uma compra feita por determinado usuário)
	commando.
		Register("usuario_valor_total").
		AddArgument("codUsuario", "Código do usuário", "1").
		SetShortDescription("Valor total de compra feita pelo usuário de código X").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			codUsuario := args["codUsuario"].Value
			fmt.Println("usuario_valor_total", codUsuario)
			query := "SELECT DISTINCT codUsuario, codPedido, (total + frete) AS valorFinal " +
				"FROM NotaFiscal " +
				"WHERE codUsuario = " + codUsuario
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Cod Usuário | Cod Pedido | Valor Final")
			for rows.Next() {
				var codUsuario, codPedido int
				var valorFinal string
				err = rows.Scan(&codUsuario, &codPedido, &valorFinal)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(codUsuario, "|", codPedido, "|", valorFinal)
			}
		})

	// (Produtos adquiridos por um determinado usuario)
	commando.
		Register("produto_adquirido_usuario").
		AddArgument("usuario", "Nome do usuário", "Jorge Da Silva").
		SetShortDescription("Todos os produtos adquiridos pelo usuário de nome X").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			nomeUsuario := args["usuario"].Value
			fmt.Println("nomeUsuario", nomeUsuario)
			query := "SELECT usuario.nome, produto.nome " +
				"FROM produto NATURAL JOIN produtofornecido NATURAL JOIN produtofornecido_pedido " +
				"NATURAL JOIN pedido NATURAL JOIN comprador JOIN usuario USING(codusuario) " +
				"WHERE usuario.nome = $1"
			rows, err := db.Query(query, nomeUsuario)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Nome Usuario | Nome Produto")
			for rows.Next() {
				var nomeComprador, nomeProduto string
				err = rows.Scan(&nomeComprador, &nomeProduto)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(nomeComprador, "|", nomeProduto)
			}
		})

	// (Produto disponível)
	commando.
		Register("produto_estoque").
		AddArgument("estoque", "Produtos que tenham mais que este numero de estoque", "1").
		SetShortDescription("Todos os produtos que tem estoque maior que X").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			estoque, _ := strconv.Atoi(args["estoque"].Value)
			fmt.Println("produto_estoque", estoque)
			query := "SELECT produto.nome, codAnunciante, estoque " +
				"FROM Produto NATURAL JOIN ProdutoFornecido NATURAL JOIN Anunciante " +
				"WHERE estoque >= $1"
			rows, err := db.Query(query, estoque)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("CodAnunciante | NomeProduto | Estoque")
			for rows.Next() {
				var codAnunciante, estoque int
				var nomeProduto string
				err = rows.Scan(&nomeProduto, &codAnunciante, &estoque)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(codAnunciante, "|", nomeProduto, "|", estoque)
			}
		})

	// (Produtos com baixo número de estrelas)
	commando.
		Register("avaliacao_nota").
		AddArgument("nota", "Avaliações que tenham nota menor ou igual a este numero", "4").
		SetShortDescription("Todas as avaliações que tem nota menor ou igual a X").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			nota, _ := strconv.Atoi(args["nota"].Value)
			fmt.Println("avaliacao_nota", nota)
			query := "SELECT DISTINCT usuario.nome, produto.nome, nota " +
				"FROM produto NATURAL JOIN avaliacao NATURAL JOIN comprador JOIN usuario USING(codUsuario) " +
				"WHERE nota <= $1"
			rows, err := db.Query(query, nota)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Nome Comprador | Nome Produto | Nota")
			for rows.Next() {
				var nota int
				var nomeComprador, nomeProduto string
				err = rows.Scan(&nomeComprador, &nomeProduto, &nota)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(nomeComprador, "|", nomeProduto, "|", nota)
			}
		})

	commando.Parse(nil)
}
