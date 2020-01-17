#!/usr/bin/perl
use strict;
use warnings;

use Data::Dumper;

my @funcs  = `grep -h "^func" *.go  | sort |uniq -d`;
my %f;
foreach my $x (@funcs) {
    # chomp $x;
    $f{$x} = 1;
}
my @types = `grep -h "^type" *.go |sed -e 's/struct.*/struct\{\}/'  | sort |uniq -d`;
my %t;
foreach my $x (@types) {
    # chomp $x;
    $t{$x} = 1;
}

open(my $fh, "<", "resolver.go")
	or die "Can't open < resolver.go: $!";

my @resolver = <$fh>;
close $fh;
my @fixed;

my $c = 0;

while ($c <= $#resolver) {
    chomp $c;
    if (exists $f{$resolver[$c]}) {
        print "skipping $resolver[$c]$resolver[$c+1]$resolver[$c+2]";
        $c += 2;
    } elsif (exists $t{$resolver[$c]}) {
        print "skipping $resolver[$c]";
    } elsif ( $resolver[$c] =~ /import/) {
        $c++ while( $resolver[$c] !~ /\)/);
        print "skipping imports\n";
        push @fixed, "import ()\n";

    } else {
        push @fixed, $resolver[$c];
    }
    $c++
}

open(my $fh2, ">", "resolver.go")
	or die "Can't open > resolver.go: $!";

foreach my $l (@fixed){
    print $fh2 $l;
}