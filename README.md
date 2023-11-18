# golang-001-calcular-taxa
Curso FullCycle

##  Comandos Docker

Criar o container:
``` bash
docker-compose up -d --build
```

Consultar Container:
``` bash
docker-compose ps
```

Acessar o container do app:
``` bash
docker-compose exec app bash
```

Acessar o container do mysql:
```
docker-compose exec app mysql
```

## Database

``` sql
CREATE TABLE `orders` (
    `id` varchar(255) NOT NULL,
    `price` float NOT NULL,
    `tax` float NOT NULL,
    `final_price` float NOT NULL,
    PRIMARY KEY (`id`)
);
```