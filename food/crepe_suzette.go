package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CrepeSuzette struct{}

// AlwaysConsumable ...
func (CrepeSuzette) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CrepeSuzette) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CrepeSuzette) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CrepeSuzette) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CrepeSuzette) Edible() bool {
	return true
}

// EncodeItem ...
func (CrepeSuzette) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:crepe_suzette", 0
}

// Name ...
func (CrepeSuzette) Name() string {
	return "Crepe Suzette"
}

// Texture ...
func (CrepeSuzette) Texture() image.Image {
	return textureFromName("crepesuzette")
}
