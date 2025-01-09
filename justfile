default:
  just --list

[group('api')]
generate:
  @cd api && go generate ./...
