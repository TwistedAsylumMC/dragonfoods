package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedChocolatePie struct{}

// AlwaysConsumable ...
func (BakedChocolatePie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedChocolatePie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedChocolatePie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedChocolatePie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedChocolatePie) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedChocolatePie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_chocolate_pie", 0
}

// Name ...
func (BakedChocolatePie) Name() string {
	return "BakedChocolatePie"
}

// Texture ...
func (BakedChocolatePie) Texture() image.Image {
	return textureFromName("bakedchocolatepie")
}
