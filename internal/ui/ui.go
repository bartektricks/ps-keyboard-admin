package ui

import (
	"fmt"
	"strings"

	"github.com/bartektricks/ps-keyboard-admin/internal/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle   = lipgloss.NewStyle().Margin(1, 2)
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)
	itemsStyle         = lipgloss.NewStyle().PaddingLeft(2)
	cursorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))
	selectedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
	idStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Italic(true)
	requestedTagsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("87"))
	prevTagsStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type RequestItem struct {
	ID            string
	Name          string
	PrevTags      []string
	RequestedTags []string
	Selected      bool
}

type Model struct {
	Choices  []RequestItem
	Cursor   int
	Selected map[int]RequestItem
	Quitting bool
}

func InitialModel(requests []model.Request) Model {
	items := make([]RequestItem, len(requests))

	for i, req := range requests {
		items[i] = RequestItem{
			ID:            req.ID,
			Name:          req.Name,
			PrevTags:      []string(req.NotVerifiedTags),
			RequestedTags: []string(req.VerifiedTags),
		}
	}

	return Model{
		Choices:  items,
		Selected: make(map[int]RequestItem),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
				m.Choices[m.Cursor].Selected = false
			} else {
				m.Selected[m.Cursor] = m.Choices[m.Cursor]
				m.Choices[m.Cursor].Selected = true
			}
		case "enter":
			if len(m.Selected) > 0 {
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.Quitting {
		return ""
	}

	s := titleStyle.Render("Verification Requests") + "\n\n"

	if len(m.Choices) == 0 {
		return s + itemsStyle.Render("No requests available.") + "\n"
	}

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = cursorStyle.Render(">")
		}

		checkbox := "[ ]"
		if _, ok := m.Selected[i]; ok {
			checkbox = selectedStyle.Render("[✓]")
		}

		nameText := choice.Name

		prevTagsText := ""
		if len(choice.PrevTags) > 0 {
			prevTagsText = prevTagsStyle.Render(fmt.Sprintf("Prev tags: %s", strings.Join(choice.PrevTags, ", ")))
		}

		requestedTagsText := ""
		if len(choice.RequestedTags) > 0 {
			requestedTagsText = requestedTagsStyle.Render(fmt.Sprintf("Requested tags: %s", strings.Join(choice.RequestedTags, ", ")))
		}

		line := fmt.Sprintf("%s %s %s\n  %s %s", cursor, checkbox, nameText, prevTagsText, requestedTagsText)
		s += itemsStyle.Render(line)
	}

	s += "\n" + docStyle.Render("↑/↓: navigate • space: select • enter: confirm • q: quit")

	return s
}

func (m Model) GetSelectedRequests() []RequestItem {
	selected := make([]RequestItem, 0, len(m.Selected))

	for _, item := range m.Selected {
		selected = append(selected, item)
	}

	return selected
}

func RunUI(requests []model.Request) ([]RequestItem, error) {
	model := InitialModel(requests)

	p := tea.NewProgram(model, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := finalModel.(Model); ok {
		return m.GetSelectedRequests(), nil
	}

	return nil, fmt.Errorf("could not get selected requests")
}
