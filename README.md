# Mini-CRM CLI

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go. Ce projet illustre les bonnes pratiques de développement Go, incluant :

- Une architecture en packages découplés
- L'injection de dépendances via les interfaces
- La création d'une CLI professionnelle avec Cobra
- La gestion de configuration externe avec Viper
- Plusieurs couches de persistance avec GORM et SQLite

## Fonctionnalités

- **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer des contacts
- **Interface en ligne de commande** : Commandes et sous-commandes claires et standardisées
- **Configuration externe** : Le comportement de l'application peut être modifié sans recompiler
- **Persistance des données** : Support de multiples backends de stockage :
  - **GORM/SQLite** : Une base de données SQL robuste contenue dans un simple fichier
  - **Fichier JSON** : Une sauvegarde simple et lisible
  - **En mémoire** : Un stockage éphémère pour les tests

## Installation

1. Clonez le repository :
```bash
git clone <repository-url>
cd TP-CLI
```

2. Installez les dépendances :
```bash
go mod tidy
```

3. Compilez l'application :
```bash
go build -o minicrm
```

## Configuration

Le fichier `config.yaml` permet de configurer le type de stockage :

```yaml
storage:
  type: "gorm"        # Options: "gorm", "json", "memory"
  database: "contacts.db"
  json_file: "contacts.json"
```

### Types de stockage

- **`gorm`** : Utilise SQLite avec GORM pour une persistance robuste
- **`json`** : Sauvegarde les données dans un fichier JSON lisible
- **`memory`** : Stockage temporaire en mémoire (données perdues à la fermeture)

## Utilisation

### Commandes disponibles

```bash
# Afficher l'aide
./minicrm --help

# Ajouter un nouveau contact
./minicrm add --name "Jean Dupont" --email "jean@example.com" --phone "0123456789" --company "Acme Corp"

# Lister tous les contacts
./minicrm list

# Mettre à jour un contact
./minicrm update 1 --phone "0987654321"

# Supprimer un contact
./minicrm delete 1
```

### Options des commandes

#### Commande `add`
- `--name, -n` : Nom du contact (obligatoire)
- `--email, -e` : Email du contact (obligatoire)
- `--phone, -p` : Téléphone du contact
- `--company, -c` : Entreprise du contact

#### Commande `update`
- `--name, -n` : Nouveau nom
- `--email, -e` : Nouvel email
- `--phone, -p` : Nouveau téléphone
- `--company, -c` : Nouvelle entreprise

## Architecture

Le projet suit une architecture modulaire avec séparation des responsabilités :

```
├── cmd/                    # Commandes CLI avec Cobra
│   ├── add.go             # Commande d'ajout
│   ├── delete.go          # Commande de suppression
│   ├── list.go            # Commande de liste
│   ├── root.go            # Commande racine
│   └── update.go          # Commande de mise à jour
├── internal/
│   ├── config/            # Gestion de la configuration
│   │   └── config.go      # Chargement config avec Viper
│   ├── models/            # Modèles de données
│   │   └── contact.go      # Structure Contact avec GORM
│   └── stores/            # Couche de persistance
│       ├── gorm_store.go  # Implémentation GORM/SQLite
│       ├── json_store.go  # Implémentation JSON
│       ├── memory_store.go # Implémentation mémoire
│       └── storer.go      # Interface Storer
├── config.yaml            # Configuration de l'application
├── main.go                # Point d'entrée
└── go.mod                  # Dépendances Go
```

### Interfaces et injection de dépendances

Le projet utilise l'interface `Storer` pour découpler la logique métier de la persistance :

```go
type Storer interface {
    Create(contact *models.Contact) error
    GetByID(id uint) (*models.Contact, error)
    GetAll() ([]models.Contact, error)
    Update(contact *models.Contact) error
    Delete(id uint) error
    GetByEmail(email string) (*models.Contact, error)
}
```

Cette approche permet de :
- Changer facilement de backend de stockage
- Tester la logique métier avec des mocks
- Maintenir un code propre et modulaire

## Dépendances

- **Cobra** : Framework CLI professionnel
- **Viper** : Gestion de configuration
- **GORM** : ORM pour Go avec support SQLite
- **SQLite** : Base de données embarquée

## Exemples d'utilisation

### Basculer entre les types de stockage

1. **Utiliser SQLite** :
```yaml
storage:
  type: "gorm"
  database: "contacts.db"
```

2. **Utiliser JSON** :
```yaml
storage:
  type: "json"
  json_file: "contacts.json"
```

3. **Utiliser la mémoire** :
```yaml
storage:
  type: "memory"
```

### Workflow typique

```bash
# Ajouter quelques contacts
./minicrm add --name "Alice" --email "alice@example.com" --company "Tech Corp"
./minicrm add --name "Bob" --email "bob@example.com" --phone "0123456789"

# Lister les contacts
./minicrm list

# Mettre à jour un contact
./minicrm update 1 --phone "0987654321"

# Supprimer un contact
./minicrm delete 2
```

## Développement

### Ajouter un nouveau type de stockage

1. Implémenter l'interface `Storer` dans `internal/stores/`
2. Ajouter le cas dans `config.NewStore()`
3. Mettre à jour la documentation

### Tests

```bash
# Compiler et tester
go build -o minicrm
./minicrm --help
```

## Licence

Ce projet est développé dans le cadre d'un TP d'architecture logicielle en Go.