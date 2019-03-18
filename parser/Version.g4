grammar Version;

versionLock
    : version
    | cmp version
    | version interval version
    ;

cmp
    : '<' | '>' | '>=' | '<='
    ;

 interval
    : '~'
    ;


version
    : Number'.'Number'.'Number
    ;

Number
    : '0' | [1-9][0-9]*
    ;