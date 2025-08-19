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
    : (text | block)*
    ;

text
  : TEXT
  ;

attribute_list
  :  attribute+
  ;

attribute
  : ID EQUALS value
  ;

value
  : INT
  | STRING
  ;
block
  :
     OPEN INS_WS? blockName=ID { !strings.HasPrefix($blockName.GetText(), "end_") }? attribute_list? INS_WS? CLOSE
     content
     OPENEND blockEndName=ID INS_WS?  CLOSE
     { $blockName.GetText() == $blockEndName.GetText() }?
    |
     OPEN INS_WS? blockName=ID attribute_list? INS_WS? CLOSE
  ;
