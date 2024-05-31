-- Rodar essa query para criar as tablas inicial de migracao
DROP TABLE IF EXISTS bun_migration_locks;
CREATE TABLE bun_migration_locks (
    id SERIAL PRIMARY KEY UNIQUE,
    version int,
    table_name VARCHAR,
    locked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DELETE FROM bun_migration_locks WHERE id <> 0;

DROP TABLE IF EXISTS bun_migrations;
CREATE TABLE bun_migrations (
  id SERIAL PRIMARY KEY UNIQUE,
  name VARCHAR(255) NOT NULL,
  group_id INTEGER NOT NULL,
  migrated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);