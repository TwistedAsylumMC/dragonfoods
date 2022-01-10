package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CornedBeefSandwich struct{}

// AlwaysConsumable ...
func (CornedBeefSandwich) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CornedBeefSandwich) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CornedBeefSandwich) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CornedBeefSandwich) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CornedBeefSandwich) Edible() bool {
	return true
}

// EncodeItem ...
func (CornedBeefSandwich) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:corned_beef_sandwich", 0
}

// Name ...
func (CornedBeefSandwich) Name() string {
	return "CornedBeefSandwich"
}

// Texture ...
func (CornedBeefSandwich) Texture() image.Image {
	return textureFromName("sandwich4")
}
