parser grammar PlaceholderParser;

options { tokenVocab=PlaceholderLexer; }
@parser::header {

}
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


//placeholder
//  :
//     OPEN INS_WS? placeholderName=ID attributeList? INS_WS? CLOSE
//     (
//        content
//        OPEN END placeholderEndName=ID INS_WS?
//        { $placeholderName.GetText() == $placeholderEndName.GetText() }?
//        CLOSE
//     )?
//  ;
//
//
// Package-level helpers for the Go target

placeholder
  :  OPEN placeholderName=ID { pushTag($placeholderName.GetText()) } CLOSE
     (content OPEN_END placeholderEndName=ID {
      name := $placeholderEndName.GetText()
      tag := peekTag()
      if tag == name {
        popTag()
        // keep as END_TAG
      } else {
        goto errorExit
      }

     } CLOSE)?

  ;
