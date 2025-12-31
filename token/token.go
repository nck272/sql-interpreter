package token

import (
	"strings"
)

type Token struct {
	Type    string
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier & Literals
	IDENT     = "IDENT"
	CHARACTER = "CHARACTER"
	VARCHAR   = "VARCHAR"
	CLOB      = "CLOB"
	NCHAR     = "NCHAR"
	NVARCHAR  = "NVARCHAR"
	INTEGER   = "INTEGER"
	INT       = "INT"
	SMALLINT  = "SMALLINT"
	BIGINT    = "BIGINT"
	DECIMAL   = "DECIMAL"
	FLOAT     = "FLOAT"
	REAL      = "REAL"
	DOUBLE    = "DOUBLE"
	PRECISION = "PRECISION"
	BINARY    = "BINARY"
	VARBINARY = "VARBINARY"
	BLOB      = "BLOB"
	DATE      = "DATE"
	TIME      = "TIME"
	TIMESTAMP = "TIMESTAMP"
	INTERVAL  = "INTERVAL"
	BOOLEAN   = "BOOLEAN"
	ARRAY     = "ARRAY"
	ROW       = "ROW"
	JSON      = "JSON"
	XML       = "XML"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// DDL Keywords
	CREATE     = "CREATE"
	ALTER      = "ALTER"
	DROP       = "DROP"
	TABLE      = "TABLE"
	VIEW       = "VIEW"
	INDEX      = "INDEX"
	SCHEMA     = "SCHEMA"
	DATABASE   = "DATABASE"
	CONSTRAINT = "CONSTRAINT"
	PRIMARY    = "PRIMARY"
	KEY        = "KEY"
	FOREIGN    = "FOREIGN"
	REFERENCES = "REFERENCES"
	UNIQUE     = "UNIQUE"
	CHECK      = "CHECK"
	DEFAULT    = "DEFAULT"

	// DML Keywords
	SELECT    = "SELECT"
	FROM      = "FROM"
	WHERE     = "WHERE"
	INSERT    = "INSERT"
	INTO      = "INTO"
	VALUES    = "VALUES"
	UPDATE    = "UPDATE"
	SET       = "SET"
	DELETE    = "DELETE"
	JOIN      = "JOIN"
	INNER     = "INNER"
	OUTER     = "OUTER"
	LEFT      = "LEFT"
	RIGHT     = "RIGHT"
	FULL      = "FULL"
	ON        = "ON"
	USING     = "USING"
	GROUP     = "GROUP"
	BY        = "BY"
	HAVING    = "HAVING"
	ORDER     = "ORDER"
	ASC       = "ASC"
	DESC      = "DESC"
	DISTINCT  = "DISTINCT"
	UNION     = "UNION"
	INTERSECT = "INTERSECT"
	EXCEPT    = "EXCEPT"
	ALLS      = "ALLS"

	// Conditions
	CASE    = "CASE"
	WHEN    = "WHEN"
	THEN    = "THEN"
	ELSE    = "ELSE"
	END     = "END"
	IF      = "IF"
	TRUE    = "TRUE"
	FALSE   = "FALSE"
	NULL    = "NULL"
	IS      = "IS"
	NOT     = "NOT"
	AND     = "AND"
	OR      = "OR"
	BETWEEN = "BETWEEN"
	IN      = "IN"
	EXISTS  = "EXISTS"
	LIKE    = "LIKE"
	ANY     = "ANY"
	SOME    = "SOME"
	ALL     = "ALL"
	EQ      = "="
	NOT_EQ  = "!="

	// Aggregate
	COUNT    = "COUNT"
	SUM      = "SUM"
	AVG      = "AVG"
	MIN      = "MIN"
	MAX      = "MAX"
	STDDEV   = "STDDEV"
	VARIANCE = "VARIANCE"

	// Transaction
	BEGIN       = "BEGIN"
	COMMIT      = "COMMIT"
	ROLLBACK    = "ROLLBACK"
	SAVEPOINT   = "SAVEPOINT"
	TRANSACTION = "TRANSACTION"

	// Session & Security
	USER          = "USER"
	CURRENT_USER  = "CURRENT_USER"
	SESSION_USER  = "SESSION_USER"
	SYSTEM_USER   = "SYSTEM_USER"
	ROLE          = "ROLE"
	AUTHORIZATION = "AUTHORIZATION"

	// Other
	OVER      = "OVER"
	PARTITION = "PARTITION"
	ROWS      = "ROWS"
	RANGE     = "RANGE"
	UNBOUNDED = "UNBOUNDED"
	PRECEDING = "PRECEDING"
	FOLLOWING = "FOLLOWING"
	CURRENT   = "CURRENT"
	AS        = "AS"
	WITH      = "WITH"
	RECURSIVE = "RECURSIVE"
	NULLIF    = "NULLIF"
	COALESCE  = "COALESCE"
	CAST      = "CAST"
	EXTRACT   = "EXTRACT"
	POSITION  = "POSITION"
	SUBSTRING = "SUBSTRING"
	TRIM      = "TRIM"
	UPPER     = "UPPER"
	LOWER     = "LOWER"
	LENGTH    = "LENGTH"
	ABS       = "ABS"
	MOD       = "MOD"
	CEIL      = "CEIL"
	FLOOR     = "FLOOR"
	ROUND     = "ROUND"
	POWER     = "POWER"
	SQRT      = "SQRT"
	LOG       = "LOG"
	EXP       = "EXP"
	SIN       = "SIN"
	COS       = "COS"
)

var keywords = map[string]string{
	"illegal":       ILLEGAL,
	"eof":           EOF,
	"ident":         IDENT,
	"character":     CHARACTER,
	"varchar":       VARCHAR,
	"clob":          CLOB,
	"nchar":         NCHAR,
	"nvarchar":      NVARCHAR,
	"integer":       INTEGER,
	"int":           INT,
	"smallint":      SMALLINT,
	"bigint":        BIGINT,
	"decimal":       DECIMAL,
	"float":         FLOAT,
	"real":          REAL,
	"double":        DOUBLE,
	"precision":     PRECISION,
	"binary":        BINARY,
	"varbinary":     VARBINARY,
	"blob":          BLOB,
	"date":          DATE,
	"time":          TIME,
	"timestamp":     TIMESTAMP,
	"interval":      INTERVAL,
	"boolean":       BOOLEAN,
	"array":         ARRAY,
	"row":           ROW,
	"json":          JSON,
	"xml":           XML,
	"create":        CREATE,
	"alter":         ALTER,
	"drop":          DROP,
	"table":         TABLE,
	"view":          VIEW,
	"index":         INDEX,
	"schema":        SCHEMA,
	"database":      DATABASE,
	"constraint":    CONSTRAINT,
	"primary":       PRIMARY,
	"key":           KEY,
	"foreign":       FOREIGN,
	"references":    REFERENCES,
	"unique":        UNIQUE,
	"check":         CHECK,
	"default":       DEFAULT,
	"select":        SELECT,
	"from":          FROM,
	"where":         WHERE,
	"insert":        INSERT,
	"into":          INTO,
	"values":        VALUES,
	"update":        UPDATE,
	"set":           SET,
	"delete":        DELETE,
	"join":          JOIN,
	"inner":         INNER,
	"outer":         OUTER,
	"left":          LEFT,
	"right":         RIGHT,
	"full":          FULL,
	"on":            ON,
	"using":         USING,
	"group":         GROUP,
	"by":            BY,
	"having":        HAVING,
	"order":         ORDER,
	"asc":           ASC,
	"desc":          DESC,
	"distinct":      DISTINCT,
	"union":         UNION,
	"intersect":     INTERSECT,
	"except":        EXCEPT,
	"alls":          ALLS,
	"case":          CASE,
	"when":          WHEN,
	"then":          THEN,
	"else":          ELSE,
	"end":           END,
	"if":            IF,
	"true":          TRUE,
	"false":         FALSE,
	"null":          NULL,
	"is":            IS,
	"not":           NOT,
	"and":           AND,
	"or":            OR,
	"between":       BETWEEN,
	"in":            IN,
	"exists":        EXISTS,
	"like":          LIKE,
	"any":           ANY,
	"some":          SOME,
	"all":           ALL,
	"count":         COUNT,
	"sum":           SUM,
	"avg":           AVG,
	"min":           MIN,
	"max":           MAX,
	"stddev":        STDDEV,
	"variance":      VARIANCE,
	"begin":         BEGIN,
	"commit":        COMMIT,
	"rollback":      ROLLBACK,
	"savepoint":     SAVEPOINT,
	"transaction":   TRANSACTION,
	"user":          USER,
	"current_user":  CURRENT_USER,
	"session_user":  SESSION_USER,
	"system_user":   SYSTEM_USER,
	"role":          ROLE,
	"authorization": AUTHORIZATION,
	"over":          OVER,
	"partition":     PARTITION,
	"rows":          ROWS,
	"range":         RANGE,
	"unbounded":     UNBOUNDED,
	"preceding":     PRECEDING,
	"following":     FOLLOWING,
	"current":       CURRENT,
	"as":            AS,
	"with":          WITH,
	"recursive":     RECURSIVE,
	"nullif":        NULLIF,
	"coalesce":      COALESCE,
	"cast":          CAST,
	"extract":       EXTRACT,
	"position":      POSITION,
	"substring":     SUBSTRING,
	"trim":          TRIM,
	"upper":         UPPER,
	"lower":         LOWER,
	"length":        LENGTH,
	"abs":           ABS,
	"mod":           MOD,
	"ceil":          CEIL,
	"floor":         FLOOR,
	"round":         ROUND,
	"power":         POWER,
	"sqrt":          SQRT,
	"log":           LOG,
	"exp":           EXP,
	"sin":           SIN,
	"cos":           COS,
}

func LookupIdentifier(literal string) string {
	ident := strings.ToLower(literal)
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
