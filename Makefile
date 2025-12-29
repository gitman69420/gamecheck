# ---------- ENV FILES ----------
DEV_ENV_FILE := .env.local
PROD_ENV_FILE := .env

# ---------- PATHS ----------
SQLC_OUTPUT_DIR := db

# ---------- UTILS ----------
define load_env
	if [ -f $(1) ]; then \
		set -a; \
		. $(1); \
		set +a; \
	fi
endef

.PHONY: dev

sqlc-gen:
	@if [ ! -d "$(SQLC_OUTPUT_DIR)" ]; then \
		echo "db directory doesn't exist, running sqlc generate..."; \
		sqlc generate; \
		echo "Done!"; \
	else \
		echo "db directory exists, skipping sqlc generate"; \
	fi

dev: sqlc-gen
	@$(call load_env, $(DEV_ENV_FILE)); \
	air