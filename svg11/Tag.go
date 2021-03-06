package svg11

import (
	"encoding/xml"
	"github.com/iamGreedy/dom/svg11/presentation"
	"fmt"
)

type Element interface {
	Parent() Element
	Childrun() []Element

	createElement(name xml.Name) Element
}

type commonTree struct {
	parent Element
	childrun []Element
}
func (s *commonTree) Parent() Element {
	return s.parent
}
func (s *commonTree) Childrun() []Element {
	return s.childrun
}

type Core struct {
	ID    *Name `xml:"id,attr"`
	Base  string `xml:"base,attr"`
	Lang  string `xml:"lang,attr"`
	Space string `xml:"space,attr"`
}
type ConditionalProcess struct {
	RequiredFeatures   string `xml:"requiredFeatures,attr"`
	RequiredExtensions string `xml:"requiredExtensions,attr"`
	SystemLanguage     string `xml:"systemLanguage,attr"`
}
type DocumentEvent struct {
	OnUnload string `xml:"onunload,attr"`
	OnAbort  string `xml:"onabort,attr"`
	OnError  string `xml:"onerror,attr"`
	OnResize string `xml:"onresize,attr"`
	OnScroll string `xml:"onscroll,attr"`
	OnZoom   string `xml:"onzoom,attr"`
}
type GraphicalEvent struct {
	OnFocusIn   string `xml:"onfocusin,attr"`
	OnFocusOut  string `xml:"onfocusout,attr"`
	OnActivate  string `xml:"onactivate,attr"`
	OnClick     string `xml:"onclick,attr"`
	OnMouseDown string `xml:"onmousedown,attr"`
	OnMouseUp   string `xml:"onmouseup,attr"`
	OnMouseOver string `xml:"onmouseover,attr"`
	OnMouseMove string `xml:"onmousemove,attr"`
	OnMouseOut  string `xml:"onmouseout,attr"`
	OnLoad      string `xml:"onload,attr"`
}
type Presentation map[presentation.Key]interface{}

func (s *Presentation) alignedarglist(b *FuncStyle, keys ...presentation.Key) {
	if len(keys) == 0{
		keys = presentation.Order
	}
	for _, v := range keys {
		if val, ok := (*s)[v];ok {
			b.Map(string(v), fmt.Sprint(val))
		}
	}
}


func (s *Presentation) xmlAttrs(attrs ...xml.Attr) (err error) {
	if *s == nil {
		*s = make(Presentation)
	}
	for _, attr := range attrs {
		err = s.xmlAttr(attr)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Presentation) xmlAttr(attr xml.Attr) error {
	switch k := presentation.Key(attr.Name.Local); k {
	case presentation.AlignmentBaseline:
		(*s)[k] = attr.Value
	case presentation.BaselineShift:
		(*s)[k] = attr.Value
	case presentation.Clip:
		(*s)[k] = attr.Value
	case presentation.ClipPath:
		(*s)[k] = attr.Value
	case presentation.ClipRule:
		(*s)[k] = attr.Value
	case presentation.Color:
		(*s)[k] = attr.Value
	case presentation.ColorInterpolation:
		(*s)[k] = attr.Value
	case presentation.ColorInterpolationFilters:
		(*s)[k] = attr.Value
	case presentation.ColorProfile:
		(*s)[k] = attr.Value
	case presentation.ColorRendering:
		(*s)[k] = attr.Value
	case presentation.Cursor:
		(*s)[k] = attr.Value
	case presentation.Direction:
		(*s)[k] = attr.Value
	case presentation.Display:
		(*s)[k] = attr.Value
	case presentation.DominantBaseline:
		(*s)[k] = attr.Value
	case presentation.EnableBackground:
		(*s)[k] = attr.Value
	case presentation.Fill:
		(*s)[k] = attr.Value
	case presentation.FillOpacity:
		(*s)[k] = attr.Value
	case presentation.FillRule:
		(*s)[k] = attr.Value
	case presentation.Filter:
		(*s)[k] = attr.Value
	case presentation.FloodColor:
		(*s)[k] = attr.Value
	case presentation.FloodOpacity:
		(*s)[k] = attr.Value
	case presentation.FontFamily:
		(*s)[k] = attr.Value
	case presentation.FontSize:
		(*s)[k] = attr.Value
	case presentation.FontSizeAdjust:
		(*s)[k] = attr.Value
	case presentation.FontStretch:
		(*s)[k] = attr.Value
	case presentation.FontStyle:
		(*s)[k] = attr.Value
	case presentation.FontVariant:
		(*s)[k] = attr.Value
	case presentation.FontWeight:
		(*s)[k] = attr.Value
	case presentation.GlyphOrientationHorizontal:
		(*s)[k] = attr.Value
	case presentation.GlyphOrientationVertical:
		(*s)[k] = attr.Value
	case presentation.ImageRendering:
		(*s)[k] = attr.Value
	case presentation.Kerning:
		(*s)[k] = attr.Value
	case presentation.LetterSpacing:
		(*s)[k] = attr.Value
	case presentation.LightingColor:
		(*s)[k] = attr.Value
	case presentation.MarkerEnd:
		(*s)[k] = attr.Value
	case presentation.MarkerMid:
		(*s)[k] = attr.Value
	case presentation.MarkerStart:
		(*s)[k] = attr.Value
	case presentation.Mask:
		(*s)[k] = attr.Value
	case presentation.Opacity:
		(*s)[k] = attr.Value
	case presentation.Overflow:
		(*s)[k] = attr.Value
	case presentation.PointerEvents:
		(*s)[k] = attr.Value
	case presentation.ShapeRendering:
		(*s)[k] = attr.Value
	case presentation.StopColor:
		(*s)[k] = attr.Value
	case presentation.StopOpacity:
		(*s)[k] = attr.Value
	case presentation.Stroke:
		(*s)[k] = attr.Value
	case presentation.StrokeDasharray:
		(*s)[k] = attr.Value
	case presentation.StrokeDashoffset:
		(*s)[k] = attr.Value
	case presentation.StrokeLinecap:
		(*s)[k] = attr.Value
	case presentation.StrokeLinejoin:
		(*s)[k] = attr.Value
	case presentation.StrokeMiterlimit:
		(*s)[k] = attr.Value
	case presentation.StrokeOpacity:
		(*s)[k] = attr.Value
	case presentation.StrokeWidth:
		(*s)[k] = attr.Value
	case presentation.TextAnchor:
		(*s)[k] = attr.Value
	case presentation.TextDecoration:
		(*s)[k] = attr.Value
	case presentation.TextRendering:
		(*s)[k] = attr.Value
	case presentation.UnicodeBidi:
		(*s)[k] = attr.Value
	case presentation.Visibility:
		(*s)[k] = attr.Value
	case presentation.WordSpacing:
		(*s)[k] = attr.Value
	case presentation.WritingMode:
		(*s)[k] = attr.Value
	}
	return nil
}

func createAnimatable(name xml.Name) Element {
	switch name.Local {
	case "animate":
		return nil
	case "animateColor":
		//TODO
		return nil
	case "animateMotion":
		//TODO
		return nil
	case "animateTransform":
		//TODO
		return nil
	case "set":
		//TODO
		return nil
	}
	return nil
}
func createDescriptive(name xml.Name) Element {
	switch name.Local {
	case "desc":
		return new(ElemDesc)
	case "metadata":
		//TODO
		return nil
	case "title":
		return new(ElemTitle)
	}
	return nil
}
func createShape(name xml.Name) Element {
	switch name.Local {
	case "circle":
		return new(ElemCircle)
	case "ellipse":
		return new(ElemEllipse)
	case "line":
		return new(ElemLine)
	case "path":
		return new(ElemPath)
	case "polygon":
		return new(ElemPolygon)
	case "polyline":
		return new(ElemPolyline)
	case "rect":
		return new(ElemRect)
	}
	return nil
}
func createStructural(name xml.Name) Element {
	switch name.Local {
	case "defs":
		return new(ElemDefine)
	case "g":
		return new(ElemGroup)
	case "svg":
		return new(ElemSVG)
	case "symbol":
		// TODO
		return nil
	case "use":
		// TODO
		return nil
	}
	return nil
}
func createGradient(name xml.Name) Element {
	switch name.Local {
	case "linearGradient":
		// TODO
		return nil
	case "radialGradient":
		// TODO
		return nil
	}
	return nil
}
func create(name xml.Name, fns ... func(name xml.Name) Element) Element {
	for _, fn := range fns {
		if e := fn(name); e != nil{
			return e
		}
	}
	return nil
}


