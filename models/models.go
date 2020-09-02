package models

import "time"

// DBCredentials keeps all the necessary information to connect to database
type DBCredentials struct {
	User     string
	Password string
	Name     string
	Address  string
	Port     string
}

// ConnectionString returns the string responsible to connect to database
func (c DBCredentials) ConnectionString() string {
	return "postgres://" + c.User + ":" + c.Password + "@" +
		c.Address + ":" + c.Port + "/" + c.Name + "?sslmode=disable"
}

type money float64

type Usuario struct {
	CodUsuario int
	Email      string
	Nome       string
	Senha      string
}

type Produto struct {
	CodProduto int
	Nome       string
	Descricao  string
	Foto       string
	Marca      string
}

type Comprador struct {
	CodComprador int
	CPF          string
	Prime        bool
	CodUsuario   int
}

type Anunciante struct {
	CodAnunciante int
	CNPJ          string
	Telefone      string
	CodUsuario    int
}

type ProdutoFornecido struct {
	CodProdFornecido int
	Preco            money
	Estoque          int
	Frete            money
	Prime            bool
	CodAnunciante    int
	CodProduto       int
}

type Categoria struct {
	CodCategoria    int
	Nome            string
	Descricao       string
	CodCategoriaPai int
}

type Cupom struct {
	CodCupom int
	Desconto float64
	Validade time.Time
}

type Endereco struct {
	CodEndereco  int
	Nome         string
	Telefone     string
	Pais         string
	Estado       string
	Cidade       string
	CEP          string
	Endereco     string
	Numero       string
	Complemento  string
	CodComprador int
}

type Cartao struct {
	CodCartao     int
	Nome          string
	Numero        string
	DataExpiracao time.Time
	CVV           int
	CodComprador  int
}

type Pedido struct {
	CodPedido    int
	Total        money
	Frete        money
	DataPedido   time.Time
	StatusPedido string // statusPedido IN ('PAGAMENTO PENDENTE', 'CONFIRMADO', 'EM TRANSPORTE', 'ENTREGUE', 'CANCELADO')
	Prime        bool
	CodCupom     int
	CodComprador int
	CodEndereco  int
	CodCartao    int
}

type ProdutoFornecido_Pedido struct {
	CodProdutoFornecido int
	CodPedido           int
	Quantidade          int
	PrecoNegociado      money
}

type Avaliacao struct {
	Nota          int
	Comentario    string
	DataAvaliacao time.Time
	CodProduto    int
	CodComprador  int
}

// ProdutoFornecido_Comprador
type Carrinho struct {
	Quantidade          int
	PrecoNegociado      money
	CodProdutoFornecido int
	CodComprador        int
}

// Produto_Comprador
type ListaDesejos struct {
	CodProduto   int
	CodComprador int
}

type Produto_Cupom struct {
	CodProduto int
	CodCupom   int
}

type Produto_Categoria struct {
	CodProduto   int
	CodCategoria int
}
