# Stress Test CLI

Este é um aplicativo de linha de comando (CLI) em Go para realizar testes de carga em serviços web. Ele permite que você especifique a URL do serviço, o número total de requisições e a concorrência (número de chamadas simultâneas).

## Requisitos

- [Go](https://golang.org/doc/install) (versão 1.16 ou superior)
- [Docker](https://docs.docker.com/get-docker/) (opcional, para execução via container)

## Como Usar (Sem Docker)

Para compilar e executar a aplicação diretamente em sua máquina:

1.  Navegue até o diretório raiz do projeto:
    ```bash
    cd /home/dg/pos-go/stress-test
    ```
2.  Compile a aplicação:
    ```bash
    go build -o stress-test .
    ```
3.  Execute o teste de carga, substituindo os valores pelos seus:
    ```bash
    ./stress-test --url=http://exemplo.com --requests=1000 --concurrency=10
    ```

## Como Usar (Com Docker)

Para construir a imagem Docker e executar a aplicação em um container:

1.  Navegue até o diretório raiz do projeto:
    ```bash
    cd stress-test
    ```
2.  Construa a imagem Docker (você só precisa fazer isso uma vez ou quando o código mudar):
    ```bash
    docker build -t stress-test-app .
    ```
3.  Execute o container Docker, substituindo os valores pelos seus:
    ```bash
    docker run stress-test-app --url=http://exemplo.com --requests=1000 --concurrency=10
    ```

## Parâmetros CLI

-   `--url`: A URL do serviço web a ser testado (obrigatório).
-   `--requests`: O número total de requisições a serem feitas (padrão: 1).
-   `--concurrency`: O número de chamadas simultâneas (padrão: 1).

## Exemplo de Uso

```bash
./stress-test --url=https://www.google.com --requests=500 --concurrency=50
```

## Relatório de Saída

Após a execução, o aplicativo gerará um relatório com as seguintes informações:

-   **Tempo total gasto**: O tempo total que o teste de carga levou para ser concluído.
-   **Quantidade total de requests realizados**: O número total de requisições enviadas.
-   **Quantidade de requests com status HTTP 200 (OK)**: O número de requisições que retornaram um status HTTP 200.
-   **Distribuição de outros códigos de status HTTP**: Uma lista de outros códigos de status HTTP recebidos e suas respectivas contagens (ex: 404, 500, etc.).
