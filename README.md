# DragonFoods

A library that implements over 500 food items in to [dragonfly](https://github.com/df-mc/dragonfly).

## Usage

This library makes use of the [Food's Plus](https://mcpedl.com/foods-plus-addon-1/) addon creates
by [Dreampixel](https://twitter.com/dreampixel7). Due to copyright issues the assets are not on this repository, but you
can download their resource pack and paste their textures in to a directory called `assets` in the root directory.

Once you have got all the assets, you can either use the provided `main.go` file or register the items in your own code.

```go
package main

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/creative"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/twistedasylummc/dragonfoods/food"
)

func main() {
	for _, x := range food.All() {
		world.RegisterItem(x)
		creative.RegisterItem(item.NewStack(x, 1))
	}
	
	// Start dragonfly etc...
}
```

> The items must be registered before the server is started otherwise they will not be sent to the player.