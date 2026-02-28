package ui

type Size string

const (
	SizeXS Size = "xs"
	SizeSM Size = "sm"
	SizeMD Size = "md"
	SizeLG Size = "lg"
	SizeXL Size = "xl"
)

type Color string

const (
	ColorPrimary   Color = "primary"
	ColorSecondary Color = "secondary"
	ColorAccent    Color = "accent"
	ColorNeutral   Color = "neutral"

	ColorBase100 Color = "base-100"
	ColorBase200 Color = "base-200"
	ColorBase300 Color = "base-300"

	ColorInfo    Color = "info"
	ColorSuccess Color = "success"
	ColorWarning Color = "warning"
	ColorError   Color = "error"
)

type ContentColor string

const (
	ContentPrimary   ContentColor = "primary-content"
	ContentSecondary ContentColor = "secondary-content"
	ContentAccent    ContentColor = "accent-content"
	ContentNeutral   ContentColor = "neutral-content"
	ContentBase      ContentColor = "base-content"
	ContentInfo      ContentColor = "info-content"
	ContentSuccess   ContentColor = "success-content"
	ContentWarning   ContentColor = "warning-content"
	ContentError     ContentColor = "error-content"
)

type Shape string

const (
	ShapeSquircle Shape = "mask-squircle"
	ShapeHeart    Shape = "mask-heart"
	ShapeHexagon  Shape = "mask-hexagon"
	ShapeRounded  Shape = "rounded-xl"
)

type Style string

const (
	StyleOutline Style = "outline"
	StyleDash    Style = "dash"
	StyleSoft    Style = "soft"
	StyleGhost   Style = "ghost"
	StyleGlass   Style = "glass"
	StyleLink    Style = "link"
	StyleRotate  Style = "rotate"
	StyleFlip    Style = "flip"
	StyleBorder  Style = "border"
)

type Behavior string

const (
	BehaviorActive   Behavior = "active"
	BehaviorDisabled Behavior = "disabled"
)

type Modifier string

const (
	ModifierWide   Modifier = "wide"
	ModifierBlock  Modifier = "block"
	ModifierSquare Modifier = "square"
	ModifierCircle Modifier = "circle"
	ModifierHover  Modifier = "hover"
	ModifierOpen   Modifier = "open"
	ModifierClose  Modifier = "close"
	ModifierActive Modifier = "active"
)

type Placement string

const (
	PlacementTop    Placement = "top"
	PlacementBot    Placement = "bot"
	PlacementLeft   Placement = "left"
	PlacementRight  Placement = "right"
	PlacementStart  Placement = "start"
	PlacementEnd    Placement = "end"
	PlacementCenter Placement = "center"
)

type MaxWidth string

const (
	MaxWidthProse MaxWidth = "max-w-prose" // Parfait pour du texte (blog)
	MaxWidthMD    MaxWidth = "max-w-3xl"   // Pour des petites sections
	MaxWidthLG    MaxWidth = "max-w-5xl"   // Pour des grilles de cards
	MaxWidthXL    MaxWidth = "max-w-7xl"   // Pour les Hero sections
	MaxWidthFull  MaxWidth = "max-w-full"
)

type Ratio string

const (
	RatioVideo    Ratio = "aspect-video"  // 16:9 (Youtube, Hero)
	RatioSquare   Ratio = "aspect-square" // 1:1 (Avatars, Instagram)
	RatioPortrait Ratio = "aspect-[3/4]"  // Pour les photos d'équipe/agents
	RatioWide     Ratio = "aspect-[21/9]" // Bannières panoramiques
)
