## Mutants Project

### Descr## Mutants Project

### Descripción
Este proyecto tiene como fun saber si un humano es mutante, esto se logre recibiendo
como parámetro un array de Strings que representan cada fila de una tabla de (NxN) con
la secuencia del ADN.

Las letras de los Strings solo pueden ser:
(A,T,C,G), las cuales representan cada base nitrogenada del ADN.

Sabrás si un humano es mutante, si encuentras más de una secuencia de cuatro letras
iguales, de forma oblicua, horizontal o vertical que no sean repetidas, es decir no puede ser
mutante si tiene dos coincidencias horizontales.

### Servicios

### API POST - Detectar Mutantes - http://18.207.247.184:8080/mutant

#### Datos de entrada
Los campos zone_id, latitude, longitude, solo se envian si se quiere aplicar las restricciones
según la ubicación del usuario.

```json
{
  "dna": [
    "ATGCGA",
    "CAGTGC",
    "TTATGT",
    "AGAAGG",
    "CCCCTA",
    "TCACTG"
  ]
}
```

#### Respuesta 200
Respuesta cuando el humano tiene ADN mutante

#### Respuesta 403
Respuesta cuando el humano no tiene ADN mutante

#### Respuesta 422
Respuesta si los datos de entrada no son validos

#### Respuesta 500
Respuesta error inesperado

### API GET - Estadisticas Humanos - http://18.207.247.184:8080/stats

#### Respuesta 200

```json
{
  "count_mutant_dna":13,
  "count_human_dna":30,
  "ratio":0
}
```

### Servicios

### API POST - Detectar Mutantes - http://18.207.247.184:8080/mutant

#### Datos de entrada
Los campos zone_id, latitude, longitude, solo se envian si se quiere aplicar las restricciones
según la ubicación del usuario.

```json
{
  "dna": [
    "ATGCGA",
    "CAGTGC",
    "TTATGT",
    "AGAAGG",
    "CCCCTA",
    "TCACTG"
  ]
}
```

#### Respuesta 200
Respuesta cuando el humano tiene ADN mutante

#### Respuesta 403
Respuesta cuando el humano no tiene ADN mutante

#### Respuesta 422
Respuesta si los datos de entrada no son validos

#### Respuesta 500
Respuesta error inesperado

### API GET - Estadisticas Mutantes - http://18.207.247.184:8080/stats

#### Respuesta 200

```json
{
  "count_mutant_dna":13,
  "count_human_dna":30,
  "ratio":0
}
```
### Arquitectura
Para la solución de este proyecto se tomo la decisión de usar una arquitectura en capas
ya que este nos permite 
