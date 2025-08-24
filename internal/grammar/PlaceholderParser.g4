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
  : OPEN placeholderName=ID attribute* CLOSE
  ;

blockPlaceholder
  : blockPlaceholderStart
    blockPlaceholderContent
    blockPlaceholderEnd
  ;

blockPlaceholderStart:
    BLOCK_OPEN INS_WS? placeholderName=ID INS_WS? attribute* INS_WS? CLOSE
    ;
blockPlaceholderEnd:
    BLOCK_END INS_WS? placeholderName=ID INS_WS? CLOSE
    ;
blockPlaceholderContent:
    INS_WS? content INS_WS?
    ;

placeholder
  :  simplePlaceholder | blockPlaceholder
  ;
