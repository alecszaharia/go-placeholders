lexer grammar PlaceholderLexer;

// ------------- Default mode (plain text) -------------

BLOCK_OPEN  : '{{' ' '* 'block_'    -> pushMode(INSIDE); // normal opener
BLOCK_END   : '{{' ' '* 'end_'      -> pushMode(INSIDE); // normal opener
OPEN        : '{{' ' '*           -> pushMode(INSIDE); // normal opener
TEXT        : ~'{'+ ( '{' ~'{' ~'{'* )* ;
LBRACE_AS_TEXT : '{' -> type(TEXT);  // treat single '{' as plain text


// ------------- Inside mode (placeholders) -------------
mode INSIDE; // tokens recognized between OPEN and CLOSE

ID      : [a-zA-Z][a-zA-Z0-9_]* ;
INS_WS  : [ \t\r\n]  ;
EQ      : '=' ;
CLOSE   : ' '* '}}' -> popMode ;
INS_ANY : ~[}] -> type(TEXT); // more specific than '.'
// double-quoted string with escapes

STRING
  : '"'  ( '\\' . | ~["\\\r\n] )* '"'
  | '\'' ( '\\' . | ~['\\\r\n] )* '\''
  ;