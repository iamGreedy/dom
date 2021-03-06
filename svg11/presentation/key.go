package presentation

type Key string

const (
	AlignmentBaseline          Key = "alignment-baseline"
	BaselineShift              Key = "baseline-shift"
	Clip                       Key = "clip"
	ClipPath                   Key = "clip-path"
	ClipRule                   Key = "clip-rule"
	Color                      Key = "color"
	ColorInterpolation         Key = "color-interpolation"
	ColorInterpolationFilters  Key = "color-interpolation-filters"
	ColorProfile               Key = "color-profile"
	ColorRendering             Key = "color-rendering"
	Cursor                     Key = "cursor"
	Direction                  Key = "direction"
	Display                    Key = "display"
	DominantBaseline           Key = "dominant-baseline"
	EnableBackground           Key = "enable-background"
	Fill                       Key = "fill"
	FillOpacity                Key = "fill-opacity"
	FillRule                   Key = "fill-rule"
	Filter                     Key = "filter"
	FloodColor                 Key = "flood-color"
	FloodOpacity               Key = "flood-opacity"
	FontFamily                 Key = "font-family"
	FontSize                   Key = "font-size"
	FontSizeAdjust             Key = "font-size-adjust"
	FontStretch                Key = "font-stretch"
	FontStyle                  Key = "font-style"
	FontVariant                Key = "font-variant"
	FontWeight                 Key = "font-weight"
	GlyphOrientationHorizontal Key = "glyph-orientation-horizontal"
	GlyphOrientationVertical   Key = "glyph-orientation-vertical"
	ImageRendering             Key = "image-rendering"
	Kerning                    Key = "kerning"
	LetterSpacing              Key = "letter-spacing"
	LightingColor              Key = "lighting-color"
	MarkerEnd                  Key = "marker-end"
	MarkerMid                  Key = "marker-mid"
	MarkerStart                Key = "marker-start"
	Mask                       Key = "mask"
	Opacity                    Key = "opacity"
	Overflow                   Key = "overflow"
	PointerEvents              Key = "pointer-events"
	ShapeRendering             Key = "shape-rendering"
	StopColor                  Key = "stop-color"
	StopOpacity                Key = "stop-opacity"
	Stroke                     Key = "stroke"
	StrokeDasharray            Key = "stroke-dasharray"
	StrokeDashoffset           Key = "stroke-dashoffset"
	StrokeLinecap              Key = "stroke-linecap"
	StrokeLinejoin             Key = "stroke-linejoin"
	StrokeMiterlimit           Key = "stroke-miterlimit"
	StrokeOpacity              Key = "stroke-opacity"
	StrokeWidth                Key = "stroke-width"
	TextAnchor                 Key = "text-anchor"
	TextDecoration             Key = "text-decoration"
	TextRendering              Key = "text-rendering"
	UnicodeBidi                Key = "unicode-bidi"
	Visibility                 Key = "visibility"
	WordSpacing                Key = "word-spacing"
	WritingMode                Key = "writing-mode"
)

var Order = []Key{
	AlignmentBaseline,
	BaselineShift,
	Clip,
	ClipPath,
	ClipRule,
	Color,
	ColorInterpolation,
	ColorInterpolationFilters,
	ColorProfile,
	ColorRendering,
	Cursor,
	Direction,
	Display,
	DominantBaseline,
	EnableBackground,
	Fill,
	FillOpacity,
	FillRule,
	Filter,
	FloodColor,
	FloodOpacity,
	FontFamily,
	FontSize,
	FontSizeAdjust,
	FontStretch,
	FontStyle,
	FontVariant,
	FontWeight,
	GlyphOrientationHorizontal,
	GlyphOrientationVertical,
	ImageRendering,
	Kerning,
	LetterSpacing,
	LightingColor,
	MarkerEnd,
	MarkerMid,
	MarkerStart,
	Mask,
	Opacity,
	Overflow,
	PointerEvents,
	ShapeRendering,
	StopColor,
	StopOpacity,
	Stroke,
	StrokeDasharray,
	StrokeDashoffset,
	StrokeLinecap,
	StrokeLinejoin,
	StrokeMiterlimit,
	StrokeOpacity,
	StrokeWidth,
	TextAnchor,
	TextDecoration,
	TextRendering,
	UnicodeBidi,
	Visibility,
	WordSpacing,
	WritingMode,
}
