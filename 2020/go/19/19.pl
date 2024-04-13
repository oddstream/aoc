#!/usr/bin/perl

use 5.028;

use strict;
use warnings;
no  warnings 'syntax';

use experimental 'signatures';
use experimental 'lexical_subs';

my $input = shift // "input";
open my $fh, "<", $input or die "open: $input";

#
# Turn a rule from the input in the Perl regexp rule
#
sub make_rule_definition ($rule) {
    $rule =~ s [^(\d+p?):] [    (?<RULE_$1>(?:]a;
    $rule =~ s [\b(\d+)\b] [(?&RULE_$1)]ag;
    $rule =~ s [\|]        [)|(?:]g;
    $rule =~ s ["]         []g;
    $rule .= "))";
    $rule;
}

my $pat1 = "(?(DEFINE)";  # Start off with a DEFINE block.

#
# Iterate over the first section, adding each rule to the DEFINE block
#
while (<$fh>) {
    last unless /\S/;
    chomp;
    $pat1 .= "\n" . make_rule_definition $_;
}

#
# For part2, make alternative rules for rules 8 and 11
#
$pat1 .= "\n" . make_rule_definition "8p: 42 | 42 8";
$pat1 .= "\n" . make_rule_definition "11p: 42 31 | 42 11 31";

#
# Close off the DEFINE block.
#
$pat1 .= "\n)";

#
# Create the DEFINE block of the second pattern by replacing
# each *call* to RULE_8 and RULE_11 to calls to RULE_8p and RULE_11p
# (We leave the rules RULE_8 and RULE_11 as is; they won't be called).
#
my $pat2 = $pat1 =~ s/\(\?&RULE_(8||11)\)/(?&RULE_${1}p)/gr;

#
# Full patterns: DEFINE block, then calling RULE_0
#
my $pattern1 = qr /$pat1^(?&RULE_0)$/x;
my $pattern2 = qr /$pat2^(?&RULE_0)$/x;


#
# Count matches and report on them.
#
chomp (my @lines = <$fh>);

say "Solution 1: " . grep {/$pattern1/} @lines;
say "Solution 2: " . grep {/$pattern2/} @lines;


__END__