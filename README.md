# openfoodfacts-go

[![Build Status](https://travis-ci.org/openfoodfacts/openfoodfacts-go.svg?branch=master)](https://travis-ci.org/openfoodfacts/openfoodfacts-go)

## About

This library is for accessing [OpenFoodFacts](http://world.openfoodfacts.org/) data for food product, ingredients and nutritional data from within your go application. You can also get the Nutri-Score and the Eco-Score

Use of the OpenFoodFacts database is subject to the [OpenFoodFacts Terms of Use](http://world.openfoodfacts.org/terms-of-use), please read them before using this library in your application.

This library is copyright © 2019-2021 OpenFoodFacts. All rights reserved. Use of this library is governed by the MIT license which can be found in the [LICENSE.txt](LICENSE.txt) file.

## Documentation

[Package documentation](https://godoc.org/github.com/j1mb0b/openfoodfacts-go)

[OpenFoodFacts API details](http://en.wiki.openfoodfacts.org/Project:API)

[Go API wiki page](http://en.wiki.openfoodfacts.org/API/Go)

## Usage details

This is a very simple example how to use the library in you go program.

~~~go
package main

import (
	"github.com/j1mb0b/openfoodfacts-go"
)

func main() {
	api := openfoodfacts.NewClient("world", "", "")
	product, err := api.Product("0737628064502")
}

~~~

There are runnable examples in the `examples` subdirectory.

## Authors
Ken Allan

## Third party applications
If you use this SDK, feel free to open a PR to add your application in this list.

