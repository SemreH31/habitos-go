# 📈 Habit Tracker - Go Project

Este es un proyecto de control de hábitos desarrollado en **Go**, diseñado para funcionar de forma local con persistencia en **SQLite**. El proyecto forma parte de un plan de aprendizaje anual enfocado en **Redes y Ciberseguridad**.

## 🚀 Características (Semana 1)
- [x] Estructura de proyecto profesional (Standard Go Layout).
- [x] Conexión a base de datos persistente con SQLite.
- [x] Gestión de configuración mediante variables de entorno (`.env`).
- [ ] Interfaz web (En progreso).
- [ ] Notificaciones vía Telegram Bot API (En progreso).

## 📁 Estructura del Proyecto
- `cmd/server/`: Punto de entrada de la aplicación.
- `internal/database/`: Lógica de conexión y esquemas de base de datos.
- `internal/habits/`: Gestión de la lógica de negocio de los hábitos.
- `web/`: Archivos frontend (HTML/CSS/JS).

## 🛠️ Tecnologías utilizadas
- **Lenguaje:** Go 1.25+
- **Base de Datos:** SQLite (vía `glebarez/go-sqlite` para soporte Pure Go sin CGO).
- **Configuración:** `joho/godotenv` para seguridad de credenciales.

## ⚙️ Configuración

1. Clone el repositorio:
   ```bash
   git clone <https://github.com/SemreH31/habitos-go>