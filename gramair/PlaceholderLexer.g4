lexer grammar PlaceholderLexer;
@lexer::header {
  var blocksTagStack []string

  func pushTag(name string) { blocksTagStack = append(blocksTagStack, name) }
  func popTag() string {
    if len(blocksTagStack) == 0 { return "" }
    t := blocksTagStack[len(blocksTagStack)-1]
    blocksTagStack = blocksTagStack[:len(blocksTagStack)-1]
    return t
  }
  func peekTag() string {
    if len(blocksTagStack) == 0 { return "" }
    return blocksTagStack[len(blocksTagStack)-1]
  }
  // "{{A}}" -> "A"
  func openNameFromText(s string) string {
    // assumes s starts with "{{" and ends with "}}"
    if len(s) >= 4 { return s[2:len(s)-2] }
    return s
  }
  // "{{end_A}}" -> "A"
  func endNameFromText(s string) string {
    // "{{" (2) + "end_" (4) = 6
    if len(s) >= 8 { return s[6:len(s)-2] }
    return s
  }
}
// ------------- Default mode (plain text) -------------

OPEN_END    : '{{end_'      -> pushMode(INSIDE); // normal opener
OPEN    : '{{'      -> pushMode(INSIDE); // normal opener
TEXT    : ( ~'{' | '{' ~'{' )+ ;                     // any run of non-'{'
//TEXT    : [^{]+;                  // any run of non-'{'
//LBRACE_AS_TEXT : '{' -> type(TEXT);  // treat single '{' as plain text


// ------------- Inside mode (placeholders) -------------
mode INSIDE; // tokens recognized between OPEN and CLOSE



ID      : [a-zA-Z_]+ ;
//INS_WS  : [ \t\r\n]+ -> skip ;
CLOSE   : '}}' -> popMode ;
INS_ANY : . -> type(TEXT); // safety net so weird chars don't error inside {{...}}

// double-quoted string with escapes

STRING
  : '"'  ( '\\' . | ~["\\\r\n] )* '"'
  | '\'' ( '\\' . | ~['\\\r\n] )* '\''
  ;