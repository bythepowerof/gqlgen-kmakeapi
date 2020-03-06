#!/usr/bin/perl
use strict;
use warnings;

use Data::Dumper;

my @funcs  = `grep "^func" resolver-*.go  | sort |uniq`;
my %f;
foreach my $x (@funcs) {
    my @a = split(':', $x);
    $f{$a[1]} = $a[0];
}
my @types = `grep  "^type" resolver-*.go |sed -e 's/struct.*/struct\{\}/'  | sort |uniq`;
my %t;
foreach my $x (@types) {
    my @a = split(':', $x);
    $t{$a[1]} = $a[0];
}

open(my $fh, "<", "resolver.go")
	or die "Can't open < resolver.go: $!";

my @resolver = <$fh>;
close $fh;
my @fixed;

my $c = 0;

my %tests;

while ($c <= $#resolver) {
    chomp $c;
    if (exists $f{$resolver[$c]}) {
        print "skipping $resolver[$c]$resolver[$c+1]$resolver[$c+2]";
        $tests{$f{$resolver[$c]}} = [] unless ($tests{$f{$resolver[$c]}});
        push @{$tests{$f{$resolver[$c]}}}, $resolver[$c];

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
    print "$k";
}

print "Orphan types\n";
foreach my $k (keys %t){
    print "$k";
}

my $test_file = q|
package gqlgen_kmakeapi

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
)

var _ = Describe("Fake client", func() {
	var k k8sclient.Client
	var fo *k8s.FakeObjects
	var r KmakeResolver

	BeforeEach(func() {

		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: &controller.KubernetesController{
				Client: k,
			},
		}
		r = res.%s()
	})

	Describe("with %s method", func() {
		Context("should be able to get", func() {
        //+ Methods Here

		})
	})
})
|;

foreach (keys %tests) {
	my ($k) = s/.go/_test.go/;
	if (-e $_) {
		print "$_ exists\n";
	} else {
		print "creating $_\n";
		open f, ">$_" or die "cannot create $_: $!\n";
		printf f $test_file, "XXX", "XXX";
		close f;
	}
}