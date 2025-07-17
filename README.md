# Proyecto Bootcamp W17-G2

Este proyecto es una API RESTful desarrollada en Go para la gestión de un sistema de almacenes, productos, empleados, compradores, vendedores y órdenes de compra. El sistema permite realizar operaciones CRUD sobre las distintas entidades y generar reportes específicos según los requerimientos del negocio.

## Descripción

La aplicación está estructurada en módulos internos que gestionan la lógica de negocio, el acceso a datos y la exposición de endpoints HTTP. Utiliza una base de datos MySQL y sigue buenas prácticas de diseño como la separación en capas (handlers, services, repositories).

## Requerimientos

El desarrollo se realizó cumpliendo **6 requerimientos** principales, cada uno abordando funcionalidades específicas del sistema:

1. **Gestión de Localidades y Vendedores** - David Garcia
2. **Gestión de Almacenes y Transportistas** - Samantha Bello
3. **Gestión de Secciones y Lotes de Productos** - Nicolas Monroy
4. **Gestión de Productos y Registros de Productos** - Sebastian Martinez
5. **Gestión de Empleados y Órdenes de Entrada** - Luis Carlos Medina
6. **Gestión de Compradores y Órdenes de Compra** - Miguel Parra

Cada requerimiento incluye endpoints para crear, consultar, actualizar y eliminar registros, así como reportes y validaciones de negocio.

## Estructura del Proyecto

- `cmd`: Archivo principal para iniciar la aplicación.
- `internal`: Lógica de negocio, controladores, servicios y repositorios.
- `pkg`: Modelos y utilidades comunes.
- `docs`: Documentación y scripts SQL para inicialización y carga de datos.
- `dev.env`: Variables de entorno para configuración local.

## Instalación y Ejecución

1. Clona el repositorio.
2. Configura la base de datos y las variables en `dev.env`.
3. Ejecuta el proyecto con:

```sh
go run cmd/main.go
```
