package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CoconutCreamPie struct{}

// AlwaysConsumable ...
func (CoconutCreamPie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CoconutCreamPie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CoconutCreamPie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CoconutCreamPie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CoconutCreamPie) Edible() bool {
	return true
}

// EncodeItem ...
func (CoconutCreamPie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:coconut_cream_pie", 0
}

// Name ...
func (CoconutCreamPie) Name() string {
	return "CoconutCreamPie"
}

// Texture ...
func (CoconutCreamPie) Texture() image.Image {
	return textureFromName("23_cheesecake_dish")
}
