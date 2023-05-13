# Tutorial - Get Started With Go

Este é o primeiro contato com a linaguagem. Nesse ponto do aprendizado estou bem confuso ainda
com toda essa novidade da linguagem.

## Um simples hello world

O arquivo hello.go, contém um simples "Hello World"
[/hello](/hello/hello.go)

## Usando um módulo externo

Agora um hello world um pouco menos simples. Usando um módulo externo.
[/hello2](/hello2/hello.go)

Através do gerenciador de pacotes online do golang, você pode procurar os pacotes que estão disponíveis para download.

Vamos usar o pacote https://pkg.go.dev/rsc.io/quote/v4 para este exemplo.

No terminal dentro do projeto, vamos fazer o download do pacote que ficara disponível e listado dentro do arquivo go.mod

usaremos o seguinte comando para fazer o download do pacote: 
```bash
go get rsc.io/quote/v4
```

Após instala-lo é só rodar o app

```bash
go run .
```
