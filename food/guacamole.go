package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Guacamole struct{}

// AlwaysConsumable ...
func (Guacamole) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Guacamole) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Guacamole) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Guacamole) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Guacamole) Edible() bool {
	return true
}

// EncodeItem ...
func (Guacamole) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:guacamole", 0
}

// Name ...
func (Guacamole) Name() string {
	return "Guacamole"
}

// Texture ...
func (Guacamole) Texture() image.Image {
	return textureFromName("guac")
}
