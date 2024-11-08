# Decisões de arquitetura
## 1. Arquitetura
Utilizei uma estrutura de pastas bastante utilizada pela comunidade Go, possuindo um pacote interno, onde organizo por feature para proporcionar uma escalabilidade modular. Pensando em um futuro adicionar serviços complementares ao delivery (ex: gerenciamento de usuários), essa separação permite uma organização mais concisa em relação ao negócio do que a abordagem de organização por camada (services, repositories, etc).

Em projetos maiores considero interessante utilizar a abordagem híbrida, organizando primeiramente por funcionalidade de negócio, e em cada funcionalidade faço a organização por camada técnica.

Para a API criei outro pacote onde organizarei as rotas e as funções que serão utilizadas na API. Também criei uma pasta para a documentação da API, que possui um arquivo swagger.json que é utilizado para gerar a documentação da API.

## 2. Utilização de interfaces
Utilizei interfaces para implementação do service e repository para ter facilidade em construir os testes unitários utilizando mocks.

## 3. Repositório
Os repositórios foram construídos pensando apenas em ser uma camada de comunicação com o banco de dados, com cada método que modifica o banco de dados sendo empacotado dentro de uma transação que só finaliza caso não ocorra nenhum erro. Isso proporciona uma consistência maior em casos onde o banco de dados pode receber muitas conexões simultâneas. Utilizei a biblioteca padrão (sql) para comunicação com o banco de dados.

## 4. Service
O service é a camada intermediária que realiza a validação de filtros e regras de negócio. É responsável por chamar os métodos do repositório para enviá-los ao handler.

## 5. Handler
O handler é onde defino quais funções serão utilizadas na API. Também é onde faço a maior parte das validações de erro que podem ser descobertos antes ou após a requisição chegar ao repositório. Utilizei a biblioteca padrão (net/http) para criação das rotas e criei funções utilitárias para reduzir o boilerplate e repetição de código (por exemplo: serializar e desserializar JSON e funções que constroem a resposta de erro da API).