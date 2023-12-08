#!/usr/bin/env tclsh
# vim:sts=4:sw=4:tw=80:et:ft=tcl

namespace path [list ::tcl::mathop ::tcl::mathfunc]

set ranks {15 7 24 6 23 5 33 4 32 3 42 2 51 1}
set ren {2 a 3 b 4 c 5 d 6 e 7 f 8 g 9 h T i J j Q k K l A m}

proc score {hand} {
    global ranks ren
    foreach c  [split $hand {}] {
        dict incr cards $c
    }
    set v [lsort -dec [dict values $cards]]
    lassign $v m
    set l [llength $v]
    set s [dict get $ranks $l$m]
    return $s[string map $ren $hand]
}

proc run {} {

    set data [read -nonewline stdin]

    foreach l [split $data \n] {
        lassign $l hand bid
        set score [score $hand]
        lappend cards [list $score $bid]
    }

    set cards [lsort -index 0 $cards]

    set n 1
    foreach card $cards {
        incr sum [* $n [lindex $card 1]]
        incr n
    }
    puts $sum

}

run
