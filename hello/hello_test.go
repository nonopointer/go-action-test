package hello

import "testing"

func TestGreet(t *testing.T){
    res := Greet()
    if res != "Hello Github Actions" {
        t.Errorf("Greet() = %s;Expected Hello Github Actions",res)
    }
}
