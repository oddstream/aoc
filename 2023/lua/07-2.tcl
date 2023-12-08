#!/usr/bin/env tclsh
# vim:sts=4:sw=4:tw=80:et:ft=tcl

namespace path [list ::tcl::mathop ::tcl::mathfunc]

set ranks {15 7 24 6 23 5 33 4 32 3 42 2 51 1}
set ren {2 a 3 b 4 c 5 d 6 e 7 f 8 g 9 h T i J J Q k K l A m}

set jranks { 421 7 322 7 331 6 223 7 232 6 241 4 124 7 133 6 132 5 142 4 151 2 }

proc score {hand} {
    global ranks ren jranks
    set new [string map $ren $hand]
    foreach c  [split $new {}] {
        dict incr cards $c
    }
    set v [lsort -dec [dict values $cards]]
    lassign $v m
    set l [llength $v]
    if {![dict exists $cards J]} {
        set s [dict get $ranks $l$m]
    } else {
        set j [dict get $cards J]
        if {$l == 1} {
            set s 7
        } else {
            if {$j == $m} {set m [lindex $v 1]}
            set s [dict get $jranks $j$l$m]
        }
    }
    return $s$new
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
