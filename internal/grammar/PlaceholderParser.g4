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
  : name=ID EQ value=STRING
  ;

simplePlaceholder
  : OPEN placeholderName=ID INS_WS* attribute* CLOSE
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
    INS_WS* content INS_WS*
    ;

placeholder
  :  simplePlaceholder | blockPlaceholder
  ;

