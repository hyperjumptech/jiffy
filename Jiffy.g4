grammar Jiffy;

root
    : timeexpr EOF
    ;

timeexpr
    : timeatom timeatom*
    ;

timeatom
    : microsecond
    | millisecond
    | second
    | minute
    | hour
    | day
    | month
    | year
    ;

microsecond
    : decimalLiteral ( US | MICROSECOND | MICROSECONDS )
    | realLiteral ( US | MICROSECOND | MICROSECONDS )
    ;

millisecond
    : decimalLiteral ( MS | MILLISECOND | MILLISECONDS )
    | realLiteral ( MS | MILLISECOND | MILLISECONDS )
    ;

second
    : decimalLiteral ( S | SEC | SECOND | SECONDS )
    | realLiteral ( S | SEC | SECOND | SECONDS )
    ;

minute
    : decimalLiteral ( M | MIN | MINUTE | MINUTES )
    | realLiteral ( M | MIN | MINUTE | MINUTES )
    ;

hour
    : decimalLiteral ( H | HRS | HOUR | HOURS )
    | realLiteral ( H | HRS | HOUR | HOURS )
    ;

day
    : decimalLiteral ( D | DAY | DAYS )
    | realLiteral ( D | DAY | DAYS )
    ;

month
    : decimalLiteral ( MON | MONTH | MONTHS )
    | realLiteral ( MON | MONTH | MONTHS )
    ;

year
    : decimalLiteral ( Y | YEAR | YEARS )
    | realLiteral ( Y | YEAR | YEARS )
    ;

decimalLiteral
    : MINUS? DECIMAL_LITERAL
    ;

realLiteral
    : MINUS? REAL_LITERAL
    ;

// LEXER HERE

fragment DEC_DIGIT          : [0-9];
fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
D                           : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
H                           : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
M                           : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
S                           : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
Y                           : [yY] ;
fragment Z                  : [zZ] ; 
fragment EXPONENT_NUM_PART  : 'E' '-'? DEC_DIGIT+;

MINUS                       : '-' ;

US                          : U S ;
MICROSECOND                 : M I C R O S E C O N D  ;
MICROSECONDS                : M I C R O S E C O N D S  ;
MS                          : M S ;
MILLISECOND                 : M I L L I S E C O N D  ;
MILLISECONDS                : M I L L I S E C O N D S  ;
SEC                         : S E C  ;
SECOND                      : S E C O N D  ;
SECONDS                     : S E C O N D S  ;
MIN                         : M I N  ;
MINUTE                      : M I N U T E  ;
MINUTES                     : M I N U T E S  ;
HRS                         : H R S  ;
HOUR                        : H O U R  ;
HOURS                       : H O U R S  ;
DAY                         : D A Y  ;
DAYS                        : D A Y S  ;
MON                         : M O N  ;
MONTH                       : M O N T H  ;
MONTHS                      : M O N T H S  ;
YEAR                        : Y E A R  ;
YEARS                       : Y E A R S  ;
AND                         : A N D {l.Skip()} ;


DECIMAL_LITERAL             : DEC_DIGIT+;

REAL_LITERAL                : (DEC_DIGIT+)? '.' DEC_DIGIT+
                            | DEC_DIGIT+ '.' EXPONENT_NUM_PART
                            | (DEC_DIGIT+)? '.' (DEC_DIGIT+ EXPONENT_NUM_PART)
                            | DEC_DIGIT+ EXPONENT_NUM_PART
                            ;


SPACE                       : [ \t,]+ {l.Skip()} ;