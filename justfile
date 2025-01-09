default:
  just --list

[group('api')]
generate:
  cd api && go generate ./...

# Open shell connected to dev database
[group('database')]
db-shell:
    @psql --username {{ env_var('POSTGRES_USER') }} --host {{ env_var('POSTGRES_HOSTNAME') }}

migration_dir := justfile_directory() / "api" / "migrations"

# Migrate the database to the latest version
[group('database')]
migrate: (_tern "migrate")

# Migration targets may be a migration number, a positive or negative delta, or
# 0 to revert all migrations.
#
# Migrate to a particular state
[group('database')]
migrate-to target: (_tern "migrate" "--destination" target)

# Create a new migration
[group('database')]
new-migration name: (_tern "new" name)

# Use `tern` to execute migrations from the correct working directory.
[group('database')]
_tern +ARGS:
    #!/usr/bin/env bash
    set -eufo pipefail
    cd {{migration_dir}}
    go run github.com/jackc/tern/v2 {{ARGS}}
