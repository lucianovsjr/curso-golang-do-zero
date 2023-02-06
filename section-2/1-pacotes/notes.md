## Módulo
* Quando estamos trabalhando com pacotes, precisamos utilizar módulos.
* Conjunto de pacotes.
* Criar um módulo: `go mod init "name"` dentro da pasta.
* Arquivo "go.mod" centraliza as dependências do projeto, faz tipo o papel do package.json.

## build
* `go build` criar um binário do módulo (Executável).

## Pacote
* Conjunto de arquivos que estão na mesma pasta.

## Arquivo
* Letra maiuscula no primeiro caracter de uma variavel, função... para dizer que algo é público.
* Letra minuscula, visivel apenas dentro do próprio pacote.
* Não é possível acessar uma função por exemplo de um módulo que esteja com a inicial minuscula.
* Dentro de outros arquivos do mesmo pacote é possível utilizar uma função por exemplo que esteja com a inicial minuscula.

## install
* Parecido com o build. Gera o binário mas salva o arquivo na raiz do Go.
* Remove o build gerado caso tenha rodado o comando `go build`.

## Pacote externo
* "go.mod" irá fazer o gerenciamento de dependências.
* Install: `go get github.com/badoux/checkmail`.
* O arquivo "go.mod" não é para ser alterado manualmente, as alterações deve ser realizadas através dos comando do Go.
* `go mod tidy` Remove todas as dependências que não estão sendo utilizadas
