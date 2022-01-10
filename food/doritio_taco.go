package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type DoritioTaco struct{}

// AlwaysConsumable ...
func (DoritioTaco) AlwaysConsumable() bool {
	return false
}

// Category ...
func (DoritioTaco) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (DoritioTaco) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (DoritioTaco) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (DoritioTaco) Edible() bool {
	return true
}

// EncodeItem ...
func (DoritioTaco) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:doritio_taco", 0
}

// Name ...
func (DoritioTaco) Name() string {
	return "DoritioTaco"
}

// Texture ...
func (DoritioTaco) Texture() image.Image {
	return textureFromName("taco2")
}
