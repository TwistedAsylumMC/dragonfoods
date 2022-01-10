package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type OatmailCookie struct{}

// AlwaysConsumable ...
func (OatmailCookie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (OatmailCookie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (OatmailCookie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (OatmailCookie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (OatmailCookie) Edible() bool {
	return true
}

// EncodeItem ...
func (OatmailCookie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:oatmail_cookie", 0
}

// Name ...
func (OatmailCookie) Name() string {
	return "OatmailCookie"
}

// Texture ...
func (OatmailCookie) Texture() image.Image {
	return textureFromName("cookie")
}
