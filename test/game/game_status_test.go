package test

import (
	"fmt"
	"soka/mtgqueue/internal/domain/game"
	"testing"
)

func TestGameSatus_Constants(t *testing.T) {
	tests := []struct {
		name     string
		status   game.Status
		expected string
	}{
		{"Created status", game.Created, "CREATED"},
		{"WaitingForPlayers status", game.WaitingForPlayers, "WAITING_FOR_PLAYERS"},
		{"Cancelled status", game.Cancelled, "CANCELLED"},
		{"Started status", game.Started, "STARTED"},
		{"Abandoned status", game.Abandoned, "ABANDONED"},
		{"Finished status", game.Finished, "FINISHED"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.status.String() != tt.expected {
				t.Errorf("Expected %s got %s", tt.expected, tt.status.String())
			}
		})
	}
}

func TestGameStatus_Constructor(t *testing.T) {
	tests := []struct {
		name           string
		status         string
		expected       game.Status
		expecetedError bool
	}{
		{"Create created status", "CREATED", game.Created, false},
		{"Create created status", "CrEaTeD", game.Created, false},
		{"Create waiting status", "WAITING_FOR_PLAYERS", game.WaitingForPlayers, false},
		{"Create cancelled status", "CANCELLED", game.Cancelled, false},
		{"Create started status", "STARTED", game.Started, false},
		{"Create abandoned status", "ABANDONED", game.Abandoned, false},
		{"Create finished status", "FINISHED", game.Finished, false},
		{"Create invalid status", "INVALID", game.Status(0), true},
		{"invalid empty string", "", game.Status(0), true},
		{"invalid number", "123", game.Status(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := game.NewStatusFromString(tt.status)

			if tt.expecetedError {
				if err == nil {
					t.Errorf("Expected error for input %s, but got none", tt.status)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %s: %v", tt.status, err)
				}

				if s != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected.String(), s.String())
				}
			}
		})
	}
}

func TestGameStatus_String(t *testing.T) {
	tests := []struct {
		name     string
		status   game.Status
		expected string
	}{
		{"Created to string", game.Created, "CREATED"},
		{"WaitingForPlayers to string", game.WaitingForPlayers, "WAITING_FOR_PLAYERS"},
		{"Cancelled to string", game.Cancelled, "CANCELLED"},
		{"Started to string", game.Started, "STARTED"},
		{"Abandoned to string", game.Abandoned, "ABANDONED"},
		{"Finished to string", game.Finished, "FINISHED"},
		{"Invalid status", game.Status(69), "UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.String()
			if result != tt.expected {
				t.Errorf("Expected %s got %s", tt.expected, result)
			}
		})
	}
}

func TestStatus_Equals(t *testing.T) {
	tests := []struct {
		name     string
		status1  game.Status
		status2  game.Status
		expected bool
	}{
		// Casos de igualdad - todos los estados consigo mismos
		{"Created equals Created", game.Created, game.Created, true},
		{"WaitingForPlayers equals WaitingForPlayers", game.WaitingForPlayers, game.WaitingForPlayers, true},
		{"Cancelled equals Cancelled", game.Cancelled, game.Cancelled, true},
		{"Started equals Started", game.Started, game.Started, true},
		{"Abandoned equals Abandoned", game.Abandoned, game.Abandoned, true},
		{"Finished equals Finished", game.Finished, game.Finished, true},

		// Casos de desigualdad - estados diferentes
		{"Created not equals WaitingForPlayers", game.Created, game.WaitingForPlayers, false},
		{"Created not equals Cancelled", game.Created, game.Cancelled, false},
		{"Created not equals Started", game.Created, game.Started, false},
		{"Created not equals Abandoned", game.Created, game.Abandoned, false},
		{"Created not equals Finished", game.Created, game.Finished, false},

		{"WaitingForPlayers not equals Created", game.WaitingForPlayers, game.Created, false},
		{"WaitingForPlayers not equals Cancelled", game.WaitingForPlayers, game.Cancelled, false},
		{"WaitingForPlayers not equals Started", game.WaitingForPlayers, game.Started, false},
		{"WaitingForPlayers not equals Abandoned", game.WaitingForPlayers, game.Abandoned, false},
		{"WaitingForPlayers not equals Finished", game.WaitingForPlayers, game.Finished, false},

		{"Cancelled not equals Created", game.Cancelled, game.Created, false},
		{"Cancelled not equals WaitingForPlayers", game.Cancelled, game.WaitingForPlayers, false},
		{"Cancelled not equals Started", game.Cancelled, game.Started, false},
		{"Cancelled not equals Abandoned", game.Cancelled, game.Abandoned, false},
		{"Cancelled not equals Finished", game.Cancelled, game.Finished, false},

		{"Started not equals Created", game.Started, game.Created, false},
		{"Started not equals WaitingForPlayers", game.Started, game.WaitingForPlayers, false},
		{"Started not equals Cancelled", game.Started, game.Cancelled, false},
		{"Started not equals Abandoned", game.Started, game.Abandoned, false},
		{"Started not equals Finished", game.Started, game.Finished, false},

		{"Abandoned not equals Created", game.Abandoned, game.Created, false},
		{"Abandoned not equals WaitingForPlayers", game.Abandoned, game.WaitingForPlayers, false},
		{"Abandoned not equals Cancelled", game.Abandoned, game.Cancelled, false},
		{"Abandoned not equals Started", game.Abandoned, game.Started, false},
		{"Abandoned not equals Finished", game.Abandoned, game.Finished, false},

		{"Finished not equals Created", game.Finished, game.Created, false},
		{"Finished not equals WaitingForPlayers", game.Finished, game.WaitingForPlayers, false},
		{"Finished not equals Cancelled", game.Finished, game.Cancelled, false},
		{"Finished not equals Started", game.Finished, game.Started, false},
		{"Finished not equals Abandoned", game.Finished, game.Abandoned, false},

		// Casos edge con valores inválidos/desconocidos
		{"Unknown status equals itself", game.Status(99), game.Status(99), true},
		{"Unknown status not equals Created", game.Status(99), game.Created, false},
		{"Created not equals unknown status", game.Created, game.Status(99), false},
		{"Different unknown statuses", game.Status(99), game.Status(100), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status1.Equals(tt.status2)

			if result != tt.expected {
				t.Errorf("Expected %s.Equals(%s) to be %v, but got %v",
					tt.status1.String(), tt.status2.String(), tt.expected, result)
			}
		})
	}
}

// Test adicional para verificar la simetría del método Equals
func TestStatus_Equals_Symmetry(t *testing.T) {
	statuses := []game.Status{
		game.Created,
		game.WaitingForPlayers,
		game.Cancelled,
		game.Started,
		game.Abandoned,
		game.Finished,
	}

	for i, status1 := range statuses {
		for j, status2 := range statuses {
			t.Run(fmt.Sprintf("Symmetry_%s_%s", status1.String(), status2.String()), func(t *testing.T) {
				result1 := status1.Equals(status2)
				result2 := status2.Equals(status1)

				if result1 != result2 {
					t.Errorf("Equals method is not symmetric: %s.Equals(%s) = %v, but %s.Equals(%s) = %v",
						status1.String(), status2.String(), result1,
						status2.String(), status1.String(), result2)
				}

				// También verificamos que la reflexividad funcione
				if i == j && !result1 {
					t.Errorf("Status should equal itself: %s.Equals(%s) should be true",
						status1.String(), status2.String())
				}
			})
		}
	}
}
