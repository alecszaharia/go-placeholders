parser grammar PlaceholderParser;

options { tokenVocab=PlaceholderLexer; }
@parser::header { import "strings" }
// ----------------------
// Parser rules
// ----------------------

template
  : content EOF
  ;

content
    : (text | placeholder)*
    ;

text
  : TEXT
  ;

attributeList
  :  attribute+
  ;

attribute
  : ID EQUALS attrValue
  ;

attrValue
  : STRING
  | TRUE
  | FALSE
  | NUMBER
  ;
placeholder
  :
     OPEN INS_WS? placeholderName=ID { !strings.HasPrefix($placeholderName.GetText(), "end_") }? attributeList? INS_WS? CLOSE
     content
     OPENEND placeholderEndName=ID INS_WS?  CLOSE
     { $placeholderName.GetText() == $placeholderEndName.GetText() }?
    |
     OPEN INS_WS? placeholderName=ID attributeList? INS_WS? CLOSE
  ;
