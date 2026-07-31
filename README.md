# 📦 LagerManager — Web-Anwendung zur Lagerverwaltung

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat&logo=mysql)](https://www.mysql.com/)
[![Nginx](https://img.shields.io/badge/Nginx-Reverse%20Proxy-009639?style=flat&logo=nginx)](https://nginx.org/)
[![Grafana](https://img.shields.io/badge/Monitoring-Grafana%20%7C%20Prometheus%20%7C%20Loki-F46800?style=flat&logo=grafana)](https://grafana.com/)

Eine moderne, containerisierte Webanwendung zur zentralen Verwaltung von Lagerbeständen, Standorten und Kategorien. Entwickelt im Rahmen eines Ausbildungsprojekts.

---

## 🚀 Key Features

* **🔍 Interaktive Suche & Verwaltung (CRUD):** Schnelles Finden, Hinzufügen, Bearbeiten und Löschen von Artikeln und Standorten.
* **📸 Foto-Upload:** Unterstützung von Bild-Uploads (inkl. Optimierung für hochauflösende Fotos vom Smartphone).
* **📱 Mobile Design:** Responsive Benutzeroberfläche für die einfache Bedienung auf Smartphones und Tablets.
* **🔒 Netzwerksicherheit & SSL:** Bereitstellung über Nginx Reverse Proxy mit vertrauenswürdiger Let's Encrypt SSL-Verschlüsselung (`https://int.lager.netways.de`).
* **🔄 Automatisierte SQL-Migrationen:** Versionierte Schema-Migrationen und Initialdaten (Seeds) via `golang-migrate`.
* **📊 Observability Stack:** System-Monitoring und Log-Aggregation mit Prometheus, Loki, Promtail und Grafana.
* **🛡️ CI/CD Pipeline:** Automatisierte Testing-, Build- und Security-Scans (Hadolint, Go Syntax Check, Trivy Vulnerability Scanner) über GitLab CI.

---

## 🛠️ Technologiestack

* **Backend:** Go (Golang 1.25)
* **Frontend:** HTML5, CSS3, Go Templates
* **Datenbank:** MySQL 8.0
* **Migrationen:** `golang-migrate`
* **Reverse Proxy / TLS:** Nginx + Certbot (Let's Encrypt)
* **Containerisierung:** Docker & Docker Compose (Multi-Stage Builds)
* **Monitoring & Logs:** Prometheus, Loki, Promtail, Grafana
* **CI/CD:** GitLab CI/CD

---

## 📁 Projektstruktur

```text
.
├── server/                # Haupt-Einstiegspunkt der Go-Anwendung (main.go)
├── internal/           # Interner Anwendungs-Code
│   ├── database/       # DB-Verbindung & golang-migrate Logik
│   └── handlers/       # HTTP Handler (Routing, Uploads, CRUD)
├── migrations/         # SQL-Migrationsdateien (Up/Down Skripte)
├── web/                # HTML-Templates & statische Assets (CSS, Images)
├── static/uploads/     # Speicherort für hochgeladene Artikel-Fotos
├── nginx.conf          # Nginx Reverse Proxy & SSL Konfiguration
├── Dockerfile          # Multi-Stage Docker Build
├── docker-compose.yml  # Docker Compose Orchestrierung
├── .gitlab-ci.yml      # GitLab CI/CD Pipeline Definition
└── README.md