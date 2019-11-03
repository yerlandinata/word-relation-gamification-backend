pg_dump -U postgres --column-inserts > db/seed/seed.sql
cat db/seed/seed.sql | grep -E "INSERT" > db/seed/seed.sql
