#!/usr/bin/perl
use strict;
use warnings;

use Data::Dumper;

my @funcs  = `grep -h "^func" resolver-*.go  | sort |uniq`;
my %f;
foreach my $x (@funcs) {
    # chomp $x;
    $f{$x} = 1;
}
my @types = `grep -h "^type" resolver-*.go |sed -e 's/struct.*/struct\{\}/'  | sort |uniq`;
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
        delete $f{$resolver[$c]};
        $c += 2;
    } elsif (exists $t{$resolver[$c]}) {
        print "skipping $resolver[$c]";
        delete $t{$resolver[$c]};
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

print "Orphan functions\n";
foreach my $k (keys %f){
    print "$k\n";
}

print "Orphan types\n";
foreach my $k (keys %t){
    print "$k\n";
}