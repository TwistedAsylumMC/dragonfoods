package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CornChowder struct{}

// AlwaysConsumable ...
func (CornChowder) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CornChowder) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CornChowder) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CornChowder) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CornChowder) Edible() bool {
	return true
}

// EncodeItem ...
func (CornChowder) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:corn_chowder", 0
}

// Name ...
func (CornChowder) Name() string {
	return "CornChowder"
}

// Texture ...
func (CornChowder) Texture() image.Image {
	return textureFromName("cornchowder")
}
