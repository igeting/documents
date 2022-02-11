# postgres command

## connect
```
psql -h <host> -p <port> -d <dbname> -U <uname> -W <password>
```

## dump and restore
```
pg_dump -h <host> -p <port> -d <dbname> -U <uname> -W <password> -Fc > test.dmp
pg_restore -h <host> -p <port> -d <dbname> -U <uname> -W <password> < test.dmp
psql -h <host> -p <port> -d <dbname> -U <uname> -W <password> -f test.dmp
```

## display

### list database
```
\l <dbname>
```

### connect database
```
\c <dbname>
```

### list table
```
\dt
```

### list index
```
\di
```

## role and user

### create user
```
create role <rname> password <password>;
create user <uname> password <password>;

```

### change password
```
\password <uname>
alter user <uname> password <password>;
```

### config user
```
alter role <rname> login;
alter user <uname> superuser|nosuperuser;
alter user <uname> createdb|nocreatedb;
alter user <uname> createrole|nocreaterole;
alter user <uname> createuser|nocreateuser;
alter user <uname> inherit|noinherit;
alter user <uname> login|nologin;
alter user <uname> replication|noreplication;
```

### config connection
```
alter user <uname> connection limit 10|-1;
```

### config group
```
create group <gname>;
grant <gname> to <rname1>,<rname2>;
revoke <gname> from <rname1>,<rname2>;
grant all on database <dbname> to <gname>;
revoke all on database <dbname> from <gname>;
grant all on all tables in schema <public> to <gname>;
revoke all on all tables in schema <public> from <gname>;
```

### drop group
```
drop owned by <gname> cascade;
drop user <gname>;
```

### drop user
```
revoke all on database <dbname> from <uname>;
revoke all on all tables in schema <public> from <uname>;
alter table <tbname> owner to <newuser>;
drop user <uname>;
```

## drop user other way
```
drop owned by <uname> cascade;
drop user <uname>;
```

### create database owner
```
create database <dbname> owner <uname>;
```

### change owner
```
alter database <dbname> owner to <uname>;
alter table <tbname> owner to <uname>;
```

## grant 

### grant schema
```
grant usage on schema <public> to <uname>;
```

### grant database all privileges
```
grant all on database <dbname> to <uname>;
```

### revoke database all privileges
```
revoke all on database <dbname> from <uname>;
```

### grant database connect
```
grant connect on database <dbname> to <uname>;
```

### grant one table all privileges 
```
grant all on table <tbname> to <uname>;
```

### revoke one table all privileges
```
revoke all on table <tbname> from <uname>;
```

### grant all tables one operation
```
grant <select> on all tables in schema <public> to <uname>;
```

### revoke all tables one operation
```
revoke <select> on all tables in schema <public> from <uname>;
```

### grant one table one operation
```
grant <update> on table <tbname> to <uname>;
```

### revoke one table one operation
```
revoke <update> on table <tbname> from <uname>;
```

### grant one table one operation for all user
```
grant <select> on table <tbname> to public;
```

### revoke one table one operation for all user
```
revoke <select> on table <tbname> from public;
```
