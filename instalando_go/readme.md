# Instalando o GO

Estou usando o Ubuntu, uma distribuição linux. E não tive nenhum problema para realizar a instalação
em minha distro.

Foi bem fácil e rápido de usar. 

Basicamente o que eu faço é: 

- Baixaro o arquivo tar
``` bash
wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
```

- Movo a posta para a pasta onde costumo colocar os meus binários instalados manualmente e descompactar

```bash
mv go1.20.4.linux-amd64.tar.gz /opt && cd /opt && tar -xvf go1.20.4.linux-amd64.tar.gz go
```

- Agora, costumo criar links simbólicos para a pasta de executáveis locais
```bash
cd /usr/local/bin && ln -s /opt/go/bin/go && ln -s /opt/go/bin/gofmt
```

Testa no terminal e seja feliz com o sua nova lang instalada.

