package Tickets

import (
	"errors"
	"fmt"
)

type Ticket struct {
	ID          string
	Title       string
	Description string
	Priority    int
	AssigneeID  string
	Status      string
}

type TicketStore struct {
	Items map[string]Ticket
}

func NewTicketStore() *TicketStore {
	return &TicketStore{
		Items: make(map[string]Ticket),
	}
}

func (ts *TicketStore) Create(t Ticket) error {
	if len(t.ID) == 0 {
		return errors.New("ID must ne non-empty")
	} else if len(t.Title) == 0 {
		return errors.New("Title must ne non-empty")
	}

	for key := range ts.Items {
		if key == t.ID {
			return errors.New("ID must be unique")
		}
	}

	if t.Priority < 1 && t.Priority > 3 {
		return errors.New("Priority must be between 1 and 3")
	}

	if t.Status != "OPEN" {
		return errors.New("Status must start as 'OPEN'")
	}

	fmt.Println("Ticket successfully added to Ticket Store")
	ts.Items[t.ID] = t

	return nil
}

func (ts *TicketStore) Assign(ticketID string, assigneeID string) error {
	ticket, ok := ts.Items[ticketID]

	if !ok {
		return errors.New("Ticket does not exist")
	}

	if ticket.Status != "OPEN" {
		return errors.New("Ticket status is not OPEN")
	}

	if len(assigneeID) == 0 {
		return errors.New("AssigneeID must ne non-empty")
	}

	fmt.Println("Ticket successfully assigned")

	ticket.AssigneeID = assigneeID
	ts.Items[ticketID] = ticket

	return nil
}

func (ts *TicketStore) Resolve(ticketID string) error {
	ticket, ok := ts.Items[ticketID]

	if !ok {
		return errors.New("Ticket does not exist")
	}

	if ticket.Status != "OPEN" {
		return errors.New("Ticket status is not OPEN")
	}

	ticket.Status = "DONE"
	ts.Items[ticketID] = ticket

	return nil
}

func (ts *TicketStore) ListAll() []Ticket {
	tickets := make([]Ticket, 0)

	for _, value := range ts.Items {
		fmt.Printf("ID: %s, Title: %s, Description: %s, Priority: %d, AssigneeID: %s, Status: %s\n", value.ID, value.Title, value.Description, value.Priority, value.AssigneeID, value.Status)
		tickets = append(tickets, value)
	}

	return tickets
}

func (ts *TicketStore) ListByStatus(status string) []Ticket {
	tickets := make([]Ticket, 0)

	for _, value := range ts.Items {
		if value.Status == status {
			fmt.Printf("ID: %s, Title: %s, Description: %s, Priority: %d, AssigneeID: %s, Status: %s\n", value.ID, value.Title, value.Description, value.Priority, value.AssigneeID, value.Status)
			tickets = append(tickets, value)
		}
	}

	return tickets
}

func (ts *TicketStore) ListUnassigned() []Ticket {
	tickets := make([]Ticket, 0)

	for _, value := range ts.Items {
		if value.Status == "OPEN" && len(value.AssigneeID) == 0 {
			fmt.Printf("ID: %s, Title: %s, Description: %s, Priority: %d, AssigneeID: %s, Status: %s\n", value.ID, value.Title, value.Description, value.Priority, value.AssigneeID, value.Status)
			tickets = append(tickets, value)
		}
	}

	return tickets
}
