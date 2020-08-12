package defs

import "testing"

func TestToForm(t *testing.T) {

var in string = "ABCDEFGHIJKLMNOP"
var out string ="AB:CD:EF:GH:IJ:KL"

result:=toForm(in,":")

if result!=out {

  t.Errorf("toForm expected %s got %s",out,result)

}

}
