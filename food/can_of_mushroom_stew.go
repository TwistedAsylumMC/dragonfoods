package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CanOfMushroomStew struct{}

// AlwaysConsumable ...
func (CanOfMushroomStew) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CanOfMushroomStew) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CanOfMushroomStew) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CanOfMushroomStew) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CanOfMushroomStew) Edible() bool {
	return true
}

// EncodeItem ...
func (CanOfMushroomStew) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:can_of_mushroom_stew", 0
}

// Name ...
func (CanOfMushroomStew) Name() string {
	return "Can Of Mushroom Stew"
}

// Texture ...
func (CanOfMushroomStew) Texture() image.Image {
	return textureFromName("can11")
}
