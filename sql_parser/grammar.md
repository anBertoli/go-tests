

## rSQL syntactic grammar
https://forcedotcom.github.io/phoenix/index.html#expression
https://pkg.go.dev/database/sql#Rows.ColumnTypes

### Expression
```
expression   → 	 comparison (("AND" | "OR") comparison)*

comparison   →   summand (">" | ">=" | "<" | "<=" | "==" | "!=") summand
                 | summand ("LIKE" | "NOT LIKE") summand
                 | summand ("IN" | "NOT IN") list
                 | "NOT" expression

list         →   "(" (summand + ",")+ ")"
                 | "(" select_list ")"

summand      → 	 factor (("-" | "+") factor)*

factor       →	 term (("/" | "*") term)*

term	     → 	 NUMBER | STRING | "true" | "false" | "null"
                 | ident 
                 | "(" expression ")"
                 | case

case         →   CASE (WHEN expression THEN expression)+
```

### Select statement

```
select       →   "SELECT" col_list "FROM" table "WHERE" expression "LIMIT" NUMBER
select_list  →   "SELECT" col      "FROM" table "WHERE" expression "LIMIT" NUMBER

col_list     →   "(" (col + ",")+ ")"
col          →   ident
table        →   ident
```

### Update statement

```
update       →   "UPDATE" table "SET" col_updates "WHERE" expression

col_updates  →   (col "=" expression ",")+
col          →   ident
table        →   ident
```

#### Other

```
* = 0 or more
+ = 1 or more
```
