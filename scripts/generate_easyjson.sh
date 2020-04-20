#!/usr/bin/env sh

basepath=$(git rev-parse --show-toplevel)


install_easyjson()
{
   go get  github.com/mailru/easyjson/...
}


install_easyjson

rm -rf $basepath/vendor

#easyjson -all $basepath/model/model.go



