package abuse

import godaddy "github.com/alyx/go-daddy"

// Ticket defines the content of a ticket generated from GetTicket()
type Ticket struct {
	Closed    bool
	ClosedAt  string
	CreatedAt string
	DomainIP  string
	Reporter  string
	Source    string
	Target    string
	TicketID  string
	Type      string
}

type TicketCreate struct {
	Info        string
	InfoUrl     string
	Intentional bool
	Proxy       string
	Source      string
	Target      string
	Type        string
}

type TicketID struct {
	TicketID string
}

type TicketList struct {
	Pagination godaddy.Pagination
	TicketIDs  []string
}
