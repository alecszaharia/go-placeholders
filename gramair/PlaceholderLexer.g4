lexer grammar PlaceholderLexer;

// ------------- Default mode (plain text) -------------

OPENEND : '{{end_'  -> pushMode(INSIDE); // close-tag opener
OPEN    : '{{'      -> pushMode(INSIDE); // normal opener
TEXT    : ( ~'{' | '{' ~'{' )+ ;                     // any run of non-'{'
//TEXT    : [^{]+;                  // any run of non-'{'
LBRACE_AS_TEXT : '{' -> type(TEXT);  // treat single '{' as plain text


// ------------- Inside mode (placeholders) -------------
mode INSIDE; // tokens recognized between OPEN and CLOSE

ID      : [a-zA-Z_][a-zA-Z0-9_]* ;
INT     : [0-9]+ ;
EQUALS  : '=' ;
INS_WS  : [ \t\r\n]+ -> skip ;
CLOSE   : '}}' -> popMode ;
INS_ANY : . -> type(TEXT); // safety net so weird chars don't error inside {{...}}

// double-quoted string with escapes
STRING
  : '"'  ( '\\' . | ~["\\\r\n] )* '"'
  | '\'' ( '\\' . | ~['\\\r\n] )* '\''
  ;