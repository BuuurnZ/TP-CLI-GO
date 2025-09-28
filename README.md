# LogAnalyzer - Analyse de Logs Distribuée

## Description

LogAnalyzer est un outil CLI développé en Go pour analyser des fichiers de logs en parallèle. Il permet de centraliser l'analyse de multiples logs et d'en extraire des informations clés, tout en gérant les erreurs de manière robuste.

## Fonctionnalités

- **Analyse concurrentielle** : Traitement en parallèle de plusieurs logs via des goroutines
- **Gestion d'erreurs robuste** : Erreurs personnalisées avec `errors.Is()` et `errors.As()`
- **Interface CLI moderne** : Utilisation de Cobra pour une interface en ligne de commande intuitive
- **Export JSON** : Génération de rapports au format JSON avec horodatage automatique
- **Architecture modulaire** : Code organisé en packages logiques

## Installation

```bash
# Cloner le repository
git clone <repository-url>
cd loganalyzer

# Installer les dépendances
go mod tidy

# Compiler le programme
go build -o loganalyzer
```

## Utilisation

### Commande principale

```bash
./loganalyzer --help
```

### Commande d'analyse

```bash
# Analyse basique
./loganalyzer analyze -c config.json

# Analyse avec export JSON
./loganalyzer analyze -c config.json -o rapport.json
```

### Options disponibles

- `-c, --config` : Fichier de configuration JSON (requis)
- `-o, --output` : Fichier de sortie JSON (optionnel)

## Configuration

Le fichier de configuration JSON doit suivre ce format :

```json
[
  {
    "id": "web-server-1",
    "path": "/var/log/nginx/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "/var/log/my_app/errors.log",
    "type": "custom-app"
  }
]
```

### Champs de configuration

- `id` : Identifiant unique pour le log
- `path` : Chemin vers le fichier de log (absolu ou relatif)
- `type` : Type de log (informations uniquement)

## Format de sortie

### Affichage console

```
=== web-server-1 ===
Chemin: /var/log/nginx/access.log
Statut: OK
Message: Analyse terminée avec succès.

=== invalid-path ===
Chemin: /non/existent/log.log
Statut: FAILED
Message: Fichier introuvable.
Erreur: stat /non/existent/log.log: no such file or directory
```

### Export JSON

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "/var/log/nginx/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "stat /non/existent/log.log: no such file or directory"
  }
]
```

## Architecture

```
loganalyzer/
├── cmd/                    # Commandes CLI
│   ├── root.go            # Commande racine
│   └── analyze.go         # Commande d'analyse
├── internal/              # Packages internes
│   ├── config/            # Gestion des configurations
│   ├── analyzer/          # Logique d'analyse
│   └── reporter/          # Export des résultats
├── main.go                # Point d'entrée
└── go.mod                 # Dépendances
```

### Packages

- **config** : Lecture et parsing des fichiers de configuration JSON
- **analyzer** : Logique d'analyse des logs avec gestion d'erreurs personnalisées
- **reporter** : Export des résultats en JSON avec horodatage automatique

## Gestion des erreurs

Le programme implémente deux types d'erreurs personnalisées :

1. **FileNotFoundError** : Fichier introuvable ou inaccessible
2. **ParsingError** : Erreur lors du parsing des données

Ces erreurs sont gérées proprement avec `errors.Is()` et `errors.As()`.

## Fonctionnalités bonus

- **Horodatage automatique** : Les fichiers d'export incluent la date (format AAMMJJ)
- **Création automatique de dossiers** : Les répertoires d'export sont créés automatiquement
- **Traitement concurrentiel** : Utilisation de goroutines et channels pour la performance

## Exemples d'utilisation

### Analyse simple

```bash
./loganalyzer analyze -c config.json
```

### Analyse avec export

```bash
./loganalyzer analyze -c config.json -o rapports/analyse_2024.json
```

Le fichier sera automatiquement nommé `250928_analyse_2024.json` (avec la date du jour).

## Développement

### Prérequis

- Go 1.24.3 ou supérieur
- Cobra CLI framework

### Structure du code

Le code suit les bonnes pratiques Go avec :
- Packages bien organisés
- Gestion d'erreurs robuste
- Interface CLI intuitive
- Tests et documentation

## Équipe de développement

- **Développeur principal** : Étudiant M2
- **Framework** : Go avec Cobra CLI
- **Architecture** : Modulaire avec packages internes

## Licence

Projet académique - TP M2
