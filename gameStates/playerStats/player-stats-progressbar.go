package playerStats

import "github.com/charmbracelet/bubbles/progress"

type ProgressBars struct {
	width    int
	color    string
	progress progress.Model
}

func NewProgressBars(width int, color string) ProgressBars {
	prg := progress.New(progress.WithWidth(width), progress.WithSolidFill(color))
	return ProgressBars{
		width:    width,
		color:    color,
		progress: prg,
	}
}

func (p *ProgressBars) Render(c, t int) string {
	return p.progress.ViewAs(calculateRatio(c, t))
}

func calculateRatio(current, max int) float64 {
	return float64(current) / float64(max)
}
