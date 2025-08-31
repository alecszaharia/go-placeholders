parser grammar PlaceholderParser;

options { tokenVocab=PlaceholderLexer; }
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

attribute
  :  INS_WS+ name=ID EQ value=STRING
  ;

simplePlaceholder
  : OPEN placeholderName=ID attribute* CLOSE
  ;

blockPlaceholder
  : blockPlaceholderStart
    blockPlaceholderContent
    blockPlaceholderEnd
  ;

blockPlaceholderStart:
    BLOCK_OPEN placeholderName=ID attribute* CLOSE
    ;
blockPlaceholderEnd:
    BLOCK_END placeholderName=ID CLOSE
    ;
blockPlaceholderContent:
    content
    ;

placeholder
  :  simplePlaceholder | blockPlaceholder
  ;

