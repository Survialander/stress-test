# Teste de carga

### Descrição

Programa que realiza um teste de carga em determinada url e gera um relatório contendo os códigos de resposta, número de requests realizadas e tempo utilizado para execução.

### Como utilizar

Para conseguir executar precisamos ter o [``Docker``](https://www.docker.com/) instalado.

Na raiz do projeto basta rodar os seguintes comandos:
```
// Cria imagem docker

docker build -t stress_test
```
```
// Executa imagem

docker run --rm stress_test --url=https://g1.com.br --requests=100 --concurrency=10
```

Variáveis utilizadas no comando:

``--url``: Url que receberá as requisições.

``--requests``: Número de requests que serão feitas.

``--concurrency``: Número de threads utilizadas para realizar o número de requests.
