package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cantaloup struct{}

// AlwaysConsumable ...
func (Cantaloup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cantaloup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cantaloup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cantaloup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cantaloup) Edible() bool {
	return true
}

// EncodeItem ...
func (Cantaloup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cantaloup", 0
}

// Name ...
func (Cantaloup) Name() string {
	return "Cantaloup"
}

// Texture ...
func (Cantaloup) Texture() image.Image {
	return textureFromName("melon_cantaloupe_01")
}
