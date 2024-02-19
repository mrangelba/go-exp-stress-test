# Desafio Stress Test - Pós Go Expert

## Como executar

### Executar com docker run

```sh
docker run mrangelba/go-exp-stress-test [flags]
```

### Flags
- `-u, --url`: URL do serviço a ser testado (obrigatório).
- `-r, --requests`: Número total de solicitações a serem enviadas (padrão 10).
- `-c, --concurrency`: Número de chamadas simultâneas (padrão 2).
- `-v, --verbose`: Exibir detalhes das requisições (padrão: false)

### Exemplo
Para fazer um teste de stress para o endereço `http://example.com` com 50 solicitações e 5 chamadas simultâneas:

```sh
docker run mrangelba/go-exp-stress-test:latest -u http://example.com -r 50 -c 5
```
Ou

```sh
docker run mrangelba/go-exp-stress-test:latest --url http://example.com --request 50 --concurrency 5
```

### Como buildar e rodar localmente

Na pasta raiz do projeto:

```sh
go mod tidy
go run ./cmd/cli/main.go [flags]
```
Ou

```sh
docker build -f build/package/Dockerfile . -t <imagem>
docker run <imagem> [flags]
```

## Saídas
```sh
----------------------------------------------------------------------
Relatório de execução
----------------------------------------------------------------------
Tempo total gasto na execução: 2.549357023s
Quantidade total de requests realizados: 1000
----------------------------------------------------------------------
Quantidade de requests com status HTTP 200: 900
Quantidade de requests com status HTTP 400: 10
Quantidade de requests com status HTTP 403: 90
----------------------------------------------------------------------
```
Com o parâmetro `-v` ou `--verbose`
```sh
> go run ./cmd/cli/main.go --url https://www.example.com -v

[2] Request 1 - URL https://www.example.com
[1] Request 1 - URL https://www.example.com
[1] Request 1 - Status: 200
[1] Request 2 - URL https://www.example.com
[2] Request 1 - Status: 200
[2] Request 2 - URL https://www.example.com
[1] Request 2 - Status: 200
[1] Request 3 - URL https://www.example.com
[2] Request 2 - Status: 200
[2] Request 3 - URL https://www.example.com
[1] Request 3 - Status: 200
[1] Request 4 - URL https://www.example.com
[2] Request 3 - Status: 200
[2] Request 4 - URL https://www.example.com
[1] Request 4 - Status: 200
[1] Request 5 - URL https://www.example.com
[2] Request 4 - Status: 200
[2] Request 5 - URL https://www.example.com
[1] Request 5 - Status: 200
[2] Request 5 - Status: 200



----------------------------------------------------------------------
Relatório de execução
----------------------------------------------------------------------
Tempo total gasto na execução: 2.237265942s
Quantidade total de requests realizados: 10
----------------------------------------------------------------------
Quantidade de requests com status HTTP 200: 10
----------------------------------------------------------------------
```
