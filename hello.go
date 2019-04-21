package testdep

import "github.com/WinChua/testmod"

func DepHello() string {
    return testmod.Hello("World")
}
