# Tutorial - Getting Started With Multi Module Workspaces

***Link para o tutorial***: 

Esse tutorial é bem legal, inicia o trabalho sobre workspaces que facilita muito o trabalho com multimodulos locais
sem a necessidade de apontar estaticamente para a pasta dos módulos

vamos começar baixando uma depenência que vamos usar no primeiro teste

No primeiro módulo hello, criamos um programa típico que mostra uma mensagem revertida usando o módulo stringutil.Reverse

Agora, vamos criar um work

```bash
go work init ./hello
```

Agora, podemos executar o projeto da pasta workspace informando o nome do módulo a ser executado.

Exemplo: 

```bash
go run example.com/hello
```

Executando essa linha no terminal, você pode ver que te encontramos o mesmo resultado de executar de dentro do módulo hello

Vamos agora, fazer o download do módulo stringutils para adicionar mais uma funcionalidade. Porém, não queremos ter que fazer re apontamentos 
para usar o módulo local.

Na pasta de workspace, vamos fazer o clone do pacote https://go.googlesource.com/example

```bash
git clone https://go.googlesource.com/example
```

Após, fazermos download do módulo do git, podemos adicionar o módulo para ser usado local no workspace

```
go work use ./example
```

Nesse ponto, já estamos usando o pacote local, mas isso é invisível à execução. Teste: 

```bash
go run .
```

Vamos agora no módulo baixado, e adicionar uma nova função ao pacote stringutil. 

No tutorial, nós alteramos o hello para adicionar a nova função criada. Mas eu vou criar uma copia do projeto pois quero ter as duas versões.

```bash
cp ./hello hello2 -r
```

vamos adicionar o hello2 também ao workspace, mas vamos mudar o nome do módulo para example.com/hello2 alterando o arquivo go.mod

```bash
go work use ./hello2
```

Agora se executarmos o segundo módulo, temos a saida experada.

```bash
go run example.com/hello2
```

Vamos modificar agora, o módulo hello2 para usar a função toupper criada.

Execute agora novamente e o hello deve sair maiusculo. 

## Conclusão

O Go foi pensado para facilitar o trabalho com multi módulos sem que seja necessário fazer gambiarras ou manter módulos separados dentro do mesmo projeto. Com o workspace conseguimos organizar muito bem nosso módulos.


## Finalizando

Resolvi guardar também uma versão do pacote example. Então vou modificar todo o projeto aqui para apontar para o meu example e não mais o do google

A primeira coisa foi fazer um fork do projeto no github: https://github.com/golang/example

Depois, entramos no projeto example no nosso local e trocamos o remote para o nosso projeto

```bash
git remote remove origin && git add origin git@github.com:brunoom1/goexample.git
```

Agora, vamos alterar os projetos para buscar do nosso módulo

