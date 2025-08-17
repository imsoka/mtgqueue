package game

type GameRepositoryInterface interface {
  Save(g *Game) error
  Update(g *Game) error
  Delete(id string) error

  //Gets
  GetAll(limit int) ([]*Game, error)
  GetById(id string) (*Game, error)
  GetByPlayer(pid string, limit int) ([]*Game, error)
  GetByStatus(s Status, limit int) ([]*Game, error)
  GetByHost(pid string, limit int) ([]*Game, error)
  GetByFormat(int Format) ([]*Game, error)

  //Gets paged
  GetByPlayerPaged(pid string, l, o int) ([]*Game, error)
  GetByStatusPaged(s Status, l, o int) ([]*Game, error)
  GetByHostPaged(pid string, l, o int) ([]*Game, error)
  GetByFormatPaged(f Format, l, o int) ([]*Game, error)

  //Utilities
  Exists(id string) (bool, error)
}

