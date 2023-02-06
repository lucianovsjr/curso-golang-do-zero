# Instalação
* Entrar no site: https://go.dev
* Baixar o instalador
* Olhar as instruções: https://go.dev/doc/install
    * Executar o comando rm e tar separadamente com sudo
* Para desinstalar o go, basta remover a pasta go de /usr/local/go

# Comandos
* Verificar váriaveis de ambiente: `go env`
* Verificar versão do go: `go version`

# Workspace
* GOPATH é a variável de ambiente com o caminho do workspace do GO
* São as pastas bin, pkg e src. Criadas de forma dinamica de acordo com o uso.
    * bin: Binários dos códigos desenvolvidos em GO. Executaveis.
    * pkg: Pacotes externos baixados.
    * src: Código fonte do que está sendo desenvolvido.
