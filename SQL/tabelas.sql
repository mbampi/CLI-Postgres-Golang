
CREATE TABLE usuario(
    codUsuario SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(100),
    nome VARCHAR(80) NOT NULL,
    senha VARCHAR(40) NOT NULL
);

CREATE TABLE comprador(
    codComprador SERIAL NOT NULL PRIMARY KEY,
    CPF CHAR(11) NOT NULL UNIQUE,
    prime BOOLEAN NOT NULL,
    codUsuario SERIAL NOT NULL,
    FOREIGN KEY(codUsuario) REFERENCES usuario
);

CREATE TABLE anunciante(
    codAnunciante SERIAL NOT NULL PRIMARY KEY,
    CNPJ CHAR(14) NOT NULL UNIQUE,
    telefone VARCHAR(40),
    codUsuario SERIAL NOT NULL,
    FOREIGN KEY(codUsuario) REFERENCES usuario
);

CREATE TABLE produto(
    codProduto SERIAL NOT NULL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao VARCHAR(1000),
    foto VARCHAR(1000) NOT NULL,
    marca VARCHAR(80) NOT NULL
);

CREATE TABLE produtofornecido(
    codProdFornecido SERIAL NOT NULL PRIMARY KEY,
    preco MONEY NOT NULL,
    estoque INT,
    frete MONEY NOT NULL,
    prime BOOLEAN NOT NULL,
    codAnunciante SERIAL NOT NULL,
    codProduto SERIAL NOT NULL,
    FOREIGN KEY(codAnunciante) REFERENCES anunciante,
    FOREIGN KEY(codProduto) REFERENCES produto
);

CREATE TABLE categoria(
    codCategoria SERIAL NOT NULL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao VARCHAR(1000) NOT NULL,
    codCategoriaPai INTEGER,
    FOREIGN KEY(codCategoriaPai) REFERENCES categoria
);

CREATE TABLE cupom(
    codCupom SERIAL NOT NULL PRIMARY KEY,
    desconto NUMERIC(5, 2),
    validade TIMETZ
);

CREATE TABLE endereco(
    codEndereco SERIAL PRIMARY KEY,
    nome VARCHAR(80) NOT NULL,
    telefone VARCHAR(40),
    pais VARCHAR(100),
    estado VARCHAR(100),
    cidade VARCHAR(100),
    CEP VARCHAR(20),
    endereco VARCHAR(500),
    numero VARCHAR(5),
    complemento VARCHAR(100),
    codComprador SERIAL NOT NULL,
    FOREIGN KEY(codComprador) REFERENCES comprador
);

CREATE TABLE cartao(
    codCartao SERIAL PRIMARY KEY,
    nome VARCHAR(80) NOT NULL,
    numero VARCHAR(19) NOT NULL,
    dataExpiracao TIMETZ,
    CVV SERIAL NOT NULL,
    codComprador SERIAL NOT NULL,
    FOREIGN KEY(codComprador) REFERENCES comprador
);

CREATE TABLE pedido(
    codPedido SERIAL NOT NULL PRIMARY KEY,
    total MONEY NOT NULL,
    frete MONEY NOT NULL,
    dataPedido TIMETZ,
    statusPedido VARCHAR CHECK (statusPedido IN ('PAGAMENTO PENDENTE', 'CONFIRMADO', 'EM TRANSPORTE', 'ENTREGUE', 'CANCELADO')),
    prime BOOLEAN NOT NULL,
    codCupom INTEGER,
    codComprador SERIAL NOT NULL,
    codEndereco SERIAL NOT NULL,
    codCartao SERIAL NOT NULL,
    FOREIGN KEY(codCupom) REFERENCES cupom,
    FOREIGN KEY(codComprador) REFERENCES comprador,
    FOREIGN KEY(codEndereco) REFERENCES endereco,
    FOREIGN KEY(codCartao) REFERENCES cartao
);

CREATE TABLE produtofornecido_pedido(
    codProdutoFornecido SERIAL NOT NULL,
    codPedido SERIAL NOT NULL,
    quantidade INT,
    precoNegociado MONEY,
    FOREIGN KEY(codProdutoFornecido) REFERENCES produtofornecido,
    FOREIGN KEY(codPedido) REFERENCES pedido
);

CREATE TABLE avaliacao(
    nota INT,
    comentario VARCHAR(1000),
    dataAvaliacao TIMETZ,
    codProduto SERIAL NOT NULL,
    codComprador SERIAL NOT NULL,
    FOREIGN KEY(codProduto) REFERENCES produto,
    FOREIGN KEY(codComprador) REFERENCES comprador
);

-- ProdutoFornecido_Comprador
CREATE TABLE carrinho(
    quantidade INT,
    precoNegociado MONEY,
    codProdutoFornecido SERIAL NOT NULL,
    codComprador SERIAL NOT NULL,
    FOREIGN KEY(codProdutoFornecido) REFERENCES produtofornecido,
    FOREIGN KEY(codComprador) REFERENCES comprador
);

-- Produto_Comprador
CREATE TABLE listaDesejos(
    codProduto SERIAL NOT NULL,
    codComprador SERIAL NOT NULL,
    FOREIGN KEY(codProduto) REFERENCES produto,
    FOREIGN KEY(codComprador) REFERENCES comprador
);

CREATE TABLE produto_cupom(
    codProduto SERIAL NOT NULL,
    codCupom SERIAL NOT NULL,
    FOREIGN KEY(codProduto) REFERENCES produto,
    FOREIGN KEY(codCupom) REFERENCES cupom
);

CREATE TABLE produto_categoria(
    codProduto SERIAL NOT NULL,
    codCategoria SERIAL NOT NULL,
    FOREIGN KEY(codProduto) REFERENCES produto,
    FOREIGN KEY(codCategoria) REFERENCES categoria
);
