# Dice

[![Static Badge](https://img.shields.io/badge/project%20use%20codesystem-green?link=https%3A%2F%2Fgithub.com%2Fgofast-pkg%2Fcodesystem)](https://github.com/gofast-pkg/codesystem)
![Build](https://github.com/winning-number/dice/actions/workflows/ci.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/winning-number/dice.svg)](https://pkg.go.dev/github.com/winning-number/dice)
[![codecov](https://codecov.io/gh/winning-number/dice/branch/main/graph/badge.svg?token=7TCE3QB21E)](https://codecov.io/gh/winning-number/dice)
[![Release](https://img.shields.io/github/release/winning-number/dice?style=flat-square)](https://github.com/winning-number/dice/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/winning-number/dice)](https://goreportcard.com/report/github.com/winning-number/dice)
[![codebeat badge](https://codebeat.co/badges/12a947eb-2c06-4f39-8e68-6dd316c471da)](https://codebeat.co/projects/github-com-winning-number-dice-main)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwinning-number%2Fdice.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fwinning-number%2Fdice?ref=badge_shield)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/winning-number/dice/blob/main/LICENSE)

Dice package allow to explore the virtual dice behavior

## Install

``` bash
$> go get github.com/winning-number/dice@latest
```

## Usage

``` Golang
import github.com/winning-number/dice

func main() {
  var err error
  var dice dice.Dice

  if dice, err = dice.New(20); err != nil {
    panic(err)
  }
  fmt.Printf("result for the roll of the dice %s", dice.Throw())
}
```

Too, Dice provide relevant informations about statistics and history.
Check the [go documentation](https://pkg.go.dev/github.com/winning-number/dice) for more details.

## Contributing

&nbsp;:grey_exclamation:&nbsp; Use issues for everything

Read more informations with the [CONTRIBUTING_GUIDE](./.github/CONTRIBUTING.md)

For all changes, please update the CHANGELOG.txt file by replacing the existant content.

Thank you &nbsp;:pray:&nbsp;&nbsp;:+1:&nbsp;

<a href="https://github.com/winning-number/dice/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=winning-number/dice" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## Licence

[MIT](https://github.com/winning-number/dice/blob/main/LICENSE)
