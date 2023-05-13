# Tutorial - Criar um módulo em GO

***Link do tutorial***: https://go.dev/doc/tutorial/getting-started.html

Começamos pela criação do módulo greetings. Esse módulo será usado pelo modo hello para mostrar a mensagem.

Vamos então copiar agora o módulo hello do primeiro exemplo para a pasta do tutorial para modificar a chamada para usar o novo módulo.

Quando alteramos o programa antigo e adicionamos o novo módulo, ele fica assim: 

```go
package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	fmt.Println(greetings.Hello("Gabriel Mendonça"))
}
```

Mais ao executar, o go não é capaz de encontrar nosso módulo. Por isso vamos ter que apontar nosso módulo example.com/greetings para nosso projeto local. Fazemos isso assim: 

```bash
go mod edit -replace example.com/greetings=./../greetings
```

Assim, o que fizemos aqui foi alterar o caminho para encontrar o módulo que pretendemos utilizar. Dizemos que sempre que chamarmos o example.com/greetings que devemos buscar na verdade nos arquivos locais. Isso é muito bom, por que nosso código se mantém intacto quando publicarmos nosso modo e mudarmos a origem.

Porém, se você executar agora o projeto 
```bash
go run .
```
Verá que ainda sim, não temos um resultado que esperávamos.

É que precisamos ainda, fazer com que o go importe de fato esse modulo para o uso. Nesse ponto anterior, só apontamos o módulo e podemos contatar abrindo o arquivo go.mod

```bash
cat go.mod
module example.com/hello

go 1.20

replace example.com/greetings => ./../greetings
```

Precisamos rodar para atualizar os pacotes o comando 
```bash
go mod tidy
```

E após teremos no nosso arquivo go.mod a seguinte linha: 

```
require example.com/greetings v0.0.0-00010101000000-000000000000
```

Agora podemos continuar e rodar nosso programa que deve retornar a saudação definida no greetings modulo.

Assim, finalizo aqui mais esse tutoria.