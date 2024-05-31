CREATE TABLE test (
  id bigint PRIMARY KEY
);

--bun:split

ALTER TABLE test ADD COLUMN name varchar(100);

--bun:split

ALTER TABLE test ADD COLUMN email varchar(100);
