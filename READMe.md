# 🚀 Order Management System

## 🌟 Descripción
Este es un **microservicio basado en Golang** para la gestión de pedidos, siguiendo una arquitectura **limpia y escalable**. Implementa un backend con **Fiber**, base de datos **PostgreSQL**, y despliegue en **Docker y AWS ECS**.

## 🏗️ Tecnologías Utilizadas
- 🔹 **Golang** con Fiber
- 🛢️ **PostgreSQL** para almacenamiento de datos
- 🐳 **Docker** y **Docker Compose** para contenedores
- ☁️ **AWS ECS** para despliegue en la nube
- 🧪 **Testify** para pruebas unitarias
- 📦 **GitHub Actions** (para CI/CD)

## 📌 Endpoints REST
| Método | Ruta               | Descripción |
|--------|--------------------|-------------|
| `GET`  | `/orders`         | Obtiene todos los pedidos |
| `GET`  | `/orders/:id`     | Obtiene un pedido por ID |
| `POST` | `/orders`         | Crea un nuevo pedido |
| `PUT`  | `/orders/:id`     | Actualiza un pedido |
| `PATCH`| `/orders/:id/status` | Modifica solo el estado de un pedido |
| `DELETE` | `/orders/:id`   | Elimina un pedido |

## 🚀 Instalación y Ejecución
### 📌 Clonar el repositorio
```bash
git clone https://github.com/tu_usuario/order-management-system.git
cd order-management-system


🐳 Usando Docker
Ejecuta:

docker-compose up --build


go mod tidy

Copiar
Editar
go run ./cmd/main.go
🧪 Pruebas Unitarias e Integración
Para ejecutar las pruebas:

bash
Copiar
Editar
go test ./...
📦 Despliegue en AWS

docker build -t order-service .
Subir a AWS ECR
Configurar tarea en AWS ECS
Configurar RDS PostgreSQL en AWS
📜 Licencia
Este proyecto está bajo la licencia MIT.