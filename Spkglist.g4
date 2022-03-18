lexer grammar Spkglist;

WS : [ \t]+ -> skip ; // skip whitespace
NL : ('\r'? '\n')+ ; // newline

// a '#' followed by any non-newline character
COMMENT : '#' ~[\r\n]* -> skip ;

QUOTED_PACKAGE: BACKTICK (~[\u0060])* BACKTICK ;
BARE_PACKAGE : (PKG_BARE ' '*)+ ;

// characters valid in bare package words
fragment PKG_BARE : (ALPHA | DIGIT | SYMBOL) ;

fragment BACKTICK: '\u0060' ;
fragment DIGIT : [0-9] ;
fragment ALPHA : [A-Za-z] ;

// most symbols, except '#' (used for comments)
fragment SYMBOL:
'@' | '!' | '$' | '%' | '^' | '&' | '*' | '(' | ')' | '_' |
'-' | '+' | '=' | '[' | ']' | '{' | '}' | ':' | ';' |
'"' | '\'' | '>' | '<' | '.' | ',' | '?' | '|' | '/' | '~'
;
