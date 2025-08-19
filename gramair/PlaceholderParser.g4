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
block
  :
     OPEN INS_WS? blockName=ID { !strings.HasPrefix($blockName.GetText(), "end_") }? attributeList? INS_WS? CLOSE
     content
     OPENEND blockEndName=ID INS_WS?  CLOSE
     { $blockName.GetText() == $blockEndName.GetText() }?
    |
     OPEN INS_WS? blockName=ID attributeList? INS_WS? CLOSE
  ;
