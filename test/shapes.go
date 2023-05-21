package main

import "github.com/fogleman/gg"

type Circle struct {
	X, Y   float64
	Radius float64
	Stroke Color
	Fill   Color
	Depth  int
}

func (c *Circle) Draw(dc *gg.Context) {
	dc.SetRGB(c.Fill.R, c.Fill.G, c.Fill.B)
	dc.DrawCircle(c.X, c.Y, c.Radius)
	dc.Fill()

	dc.SetRGB(c.Stroke.R, c.Stroke.G, c.Stroke.B)
	dc.DrawCircle(c.X, c.Y, c.Radius)
	dc.Stroke()
}

type Triangle struct {
	X1, Y1, X2, Y2, X3, Y3 float64
	Stroke                 Color
	Fill                   Color
	Depth                  int
}

func (t *Triangle) Draw(dc *gg.Context) {
	dc.SetRGB(t.Fill.R, t.Fill.G, t.Fill.B)
	dc.MoveTo(t.X1, t.Y1)
	dc.LineTo(t.X2, t.Y2)
	dc.LineTo(t.X3, t.Y3)
	dc.ClosePath()
	dc.Fill()

	dc.SetRGB(t.Stroke.R, t.Stroke.G, t.Stroke.B)
	dc.MoveTo(t.X1, t.Y1)
	dc.LineTo(t.X2, t.Y2)
	dc.LineTo(t.X3, t.Y3)
	dc.ClosePath()
	dc.Stroke()
}

type Square struct {
	X, Y   float64
	Width  float64
	Stroke Color
	Fill   Color
	Depth  int
}

func (s *Square) Draw(dc *gg.Context) {
	dc.SetRGB(s.Fill.R, s.Fill.G, s.Fill.B)
	dc.DrawRectangle(s.X, s.Y, s.Width, s.Width)
	dc.Fill()

	dc.SetRGB(s.Stroke.R, s.Stroke.G, s.Stroke.B)
	dc.DrawRectangle(s.X, s.Y, s.Width, s.Width)
	dc.Stroke()
}

type Polygon struct {
	Points []Point
	Stroke Color
	Fill   Color
	Depth  int
}

type Point struct {
	X, Y float64
}

func (p *Polygon) Draw(dc *gg.Context) {
	dc.SetRGB(p.Fill.R, p.Fill.G, p.Fill.B)
	for i, point := range p.Points {
		if i == 0 {
			dc.MoveTo(point.X, point.Y)
		} else {
			dc.LineTo(point.X, point.Y)
		}
	}
	dc.ClosePath()
	dc.Fill()

	dc.SetRGB(p.Stroke.R, p.Stroke.G, p.Stroke.B)
	for i, point := range p.Points {
		if i == 0 {
			dc.MoveTo(point.X, point.Y)
		} else {
			dc.LineTo(point.X, point.Y)
		}
	}
	dc.ClosePath()
	dc.Stroke()
}

type Line struct {
	X1, Y1, X2, Y2 float64
	Stroke         Color
	Depth          int
}

func (l *Line) Draw(dc *gg.Context) {
	dc.SetRGB(l.Stroke.R, l.Stroke.G, l.Stroke.B)
	dc.MoveTo(l.X1, l.Y1)
	dc.LineTo(l.X2, l.Y2)
	dc.Stroke()
}

type Ellipse struct {
	X, Y    float64
	RadiusX float64
	RadiusY float64
	Stroke  Color
	Fill    Color
	Depth   int
}

func (e *Ellipse) Draw(dc *gg.Context) {
	dc.SetRGB(e.Fill.R, e.Fill.G, e.Fill.B)
	dc.DrawEllipse(e.X, e.Y, e.RadiusX, e.RadiusY)
	dc.Fill()

	dc.SetRGB(e.Stroke.R, e.Stroke.G, e.Stroke.B)
	dc.DrawEllipse(e.X, e.Y, e.RadiusX, e.RadiusY)
	dc.Stroke()
}
