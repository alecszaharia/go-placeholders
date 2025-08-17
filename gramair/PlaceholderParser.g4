parser grammar PlaceholderParser;

options { tokenVocab=PlaceholderLexer; }

// ----------------------
// Parser rules
// ----------------------

template
  : (text | placeholder)* EOF
  ;

text
  : TEXT
  ;

placeholder
  : simple_placeholder
  | block_placeholder
  ;

simple_placeholder
  : OPEN IDENT attribute_list? CLOSE
  ;

block_placeholder
  : OPEN head=IDENT attribute_list? CLOSE
    template
    OPEN END tail=IDENT CLOSE
    // To enforce matching names (Go target), uncomment:
    { $head.text == $tail.text }?
  ;

attribute_list
  : attribute (COMMA attribute)*
  ;

attribute
  : IDENT EQUALS value
  ;

value
  : INT
  | DQSTRING
  | SQSTRING
  ;
