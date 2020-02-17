# Crear la imagen optimizada:
- docker build -t hello-docker:2.0 . -f Dockerfile.optimizado
# Correr el contenedor exponiendo puertos y la carpeta public 
- docker run --name servidor -d -p 8001:8001 -v ${PWD}/public/:/app/public hello-server:1.0

## Imagen con 6 MB corriendo como servidor estático   👑