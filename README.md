

# Meli Quasar Challenge
## Operación Fuego de Quasar
Repositorio contenedor de la solución para el challenge propuesto por mercado libre denominado 
[Operación Fuego de Quasar.](https://github.com/bbasstyle/meli/blob/main/doc/backend-operacion-fuego-de-quasar.pdf)


> Han Solo ha sido recientemente nombrado General de l a Alianza
> Rebelde y busca dar un gran golpe contra el Imperio Galáctico para
> reavivar la llama de la resistencia.

> El servicio de inteligencia rebelde ha detectado un llamado de auxilio de
> una nave portacarga imperial a la deriva en un campo de asteroides. El
> manifiesto de la nave es ultra clasificado, pero se rumorea que
> transporta raciones y armamento para una legión entera.

![mill](./doc/img/mil-falcon.jpg)
_________________
## Análisis del problema
### Problema Planteado
Obtener la ubicación de un punto desconocido teniendo la ubicación de otros 3 puntos y su distancia hasta el cuarto punto en cuestión. Para resolver este problema, después de mucho buscar por la web terminé documentandome acerca del concepto de trilateración;

*La trilateración es un método matemático para determinar las posiciones relativas de objetos usando la geometría de triángulos de forma análoga la triangulación. A diferencia de esta, que usa medidas de ángulo, la trilateración usa las localizaciones conocidas de dos o más puntos de referencia, y la distancia medida entre el sujeto y cada punto de referencia. Para determinar de forma única y precisa la localización relativa de un punto en un plano bidimensional usando solo trilateración, se necesitan generalmente al menos 3 puntos de referencia.* [ref]( https://amp.blog.buy-es.com/1849965/1/trilateracion.html)

### Ejemplo gráfico
![possible example](./doc/img/graphically-possible-example.JPG)

Teniendo más claridad de cómo resolver el problema, me puse manos a la obra.
_________________
## Diseño de la solución

### Diagrama de despliegue
![deploy diagram](./doc/img/challenge-meli-aws-deploy-diagram.png)

### Diagrama de secuencia
![sequence diagram](./doc/img/challenge-meli-sequence.png)

_________________
## Ambientación Local

### Prerequisitos Generales
- contar con una cuenta activa con privilegios en los servicios aws referenciados en la sección [`Tools Box`](#Tools-Box)
- command line
- git
- aws cli
- go 
- make command
- editor de texto o "ide"

### Template de despliegue de configuración 
El archivo [template.yml](./aws-stack/template.yml) contiene la configuración con los recursos del stack aws.
- AWS::Serverless::Api
- AWS::Serverless::Function
- AWS::ApiGateway::UsagePlan
- AWS::DynamoDB::Table
- AWS::Cognito::UserPool
- AWS::Cognito::UserPoolResourceServer
- AWS::Cognito::UserPoolClient
- AWS::Cognito::UserPoolDomain

### Archivo de configuración 
El archivo [.env](./aws-stack/.env) contiene las propiedades a utilizar en el despligue del stack aws

- `AWS_ACCOUNT_ID=id cuenta aws a utilizar`
- `AWS_BUCKET_NAME=nombre del bucket s3 a utilizar para los despliegues`
- `AWS_STACK_NAME=nombre del stack`
- `AWS_REGION=región donde se instalarán los artefactos`

### Comandos de ayuda
Para simplificar el trabajo en ambiente local se disponibiliza un MakeFile que apoya con la ejecución de comandos;
- **make clean** (limpia directorio dist)
- **make build** (_clean_ + go build)
- **make install** (instala dependencias)
- **make configure** (creación del bucket en s3 que contendrá el zip del código a desplegar)
- **make test** (ejecución de pruebas unitarias) _algunos test requieren que el stack esté desplegado_
- **make put-satellites-into-db** (inserta los 3 registros correspondiente a los satellites en dynamodb)
- **make package** (_build_ + cloudformation package)
- **make deploy** (cloudformation deploy)
- **make describe** (cloudformation describe stack)

_________________
## Despliegue en AWS

### comandos para desplegar el stack en aws
- **make configure** (sólo una única vez y antes del primer despliegue del stack)
- **make build** 
- **make package**
- **make deploy**
- **put-satellites-into-db** (sólo una única vez y debe ser ejecutado luego de haber desplegado correctamente el stack)
- **make describe**

_________________
## URL API Quasar

### Postman
Para probar la API creada se disponibiliza un set de archivos [collection](./postman/) de postman donde se encuentran las peticiones a los distintos endpoints generados.

### Pasos para realizar pruebas sobre la API.

- Obtener un token válido a través del _request_ creado bajo la carpeta 'cognito'. Este _request_ ya cuenta con la configuración correspondiente.
- Agregar el _jwt_ obtenido en el paso anterior en el _header_ 'Authorization' de los _requests_ disponibles en los directorios 'topsecret' & 'topsecret_split'
- En el archivo de [variables globales](./postman/quasar.postman_globals.json) se encuentra el id de la API generada, el cual es utilizado en los _request_ generados.

![postman](./doc/img/postman-view.JPG)

De igual forma con el comando **make describe** se puede obtener la url de API generada.

https://z7imxrjh6e.execute-api.us-east-2.amazonaws.com/Prod
_________________
## Tools Box
- golang 1.16
- make 4.2.1
- aws cli 1.18.69
- aws
    - sam
    - lambda
    - api gateway
    - cloudwatch
    - cloudformation
    - s3
    - dynamo
    - cognito
- curl
- postman
- visual studio code
- generador MD online https://dillinger.io/
- formateador json http://jsonviewer.stack.hu/
_________________
## TODO
- Agregar swagger al proyecto

_________________
## Autor
- Bastián Bastías Sánchez, Software Architect.
_________________
### Referencias
- https://www.physicsforums.com/threads/how-to-calculate-2d-trilateration-step-by-step.874246/
- https://intellipaat.com/community/14464/2d-trilateration
- https://www.researchgate.net/publication/265336167_A_Novel_Trilateration_Algorithm_for_Localization_of_a_TransmitterReceiver_Station_in_a_2D_Plane_Using_Analytical_Geometry
- https://www.physicsforums.com/threads/2d-trilateration-with-3-sensors.680652/
- https://stackoverflow.com/questions/2813615/trilateration-using-3-latitude-and-longitude-points-and-3-distances
- https://gis.stackexchange.com/questions/66/trilateration-using-3-latitude-longitude-points-and-3-distances
- https://stackoverflow.com/questions/29656921/trilateration-2d-algorithm-implementation
- https://gis.stackexchange.com/questions/40660/trilateration-algorithm-for-n-amount-of-points?newreg=8c203a449c8442a09ac5891e2b72d66e
- https://math.stackexchange.com/questions/884807/find-x-location-using-3-known-x-y-location-using-trilateration
- https://math.stackexchange.com/questions/100448/finding-location-of-a-point-on-2d-plane-given-the-distances-to-three-other-know
- http://es.onlinemschool.com/math/assistance/cartesian_coordinate/p_length/
- https://es.symbolab.com/solver/equation-calculator/

