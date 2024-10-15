package playerStats

import "github.com/charmbracelet/bubbles/progress"

type ProgressBar struct {
	width    int
	color    string
	progress progress.Model
}

func NewProgressBar(width int, color string) ProgressBar {
	prg := progress.New(progress.WithWidth(width), progress.WithSolidFill(color))
	return ProgressBar{
		width:    width,
		color:    color,
		progress: prg,
	}
}

func (p *ProgressBar) Render(c, t int) string {
	return p.progress.ViewAs(calculateRatio(c, t))
}

func calculateRatio(current, max int) float64 {
	return float64(current) / float64(max)
}
