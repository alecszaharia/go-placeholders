lexer grammar PlaceholderLexer;

// ------------- Default mode (plain text) -------------

TEXT
  : ( ~'{' | '{' ~'{' )+           // any chars that don't start '{{'
  ;

OPEN
  : '{{' -> pushMode(INSIDE)
  ;

// ------------- Placeholder-inside mode -------------
mode INSIDE;

CLOSE     : '}}' -> popMode ;
END       : 'end_' ;
COMMA     : ',' ;
EQUALS    : '=' ;
IDENT     : [a-zA-Z_] [a-zA-Z_0-9]* ;
INT       : [0-9]+ ;

// double-quoted string with escapes
DQSTRING  : '"' ( ~["\\] | '\\' . )* '"' ;
// single-quoted string with escapes
SQSTRING  : '\'' ( ~['\\] | '\\' . )* '\'' ;

// ignore whitespace inside {{ ... }}
WSIN      : [ \t\r\n]+ -> skip ;
