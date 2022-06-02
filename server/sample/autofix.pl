use strict;

local @ARGV = "ent/mutation.go";
local $^I = '';
while (<>) {
    s{("github.com/cybozu-go/scim/server/sample/ent/name")}{entname $1};
    s{name\.Field}{entname.Field}g;
    s{name\.Edge}{entname.Edge}g;
    print;
}
