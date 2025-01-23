package mirrorMove

import (
    Repository "mirrorMove/src/repository"
	Dto "mirrorMove/src/dtos"

)


type MoveService struct {
    repo *Repository.MoveRepository
}

func NewMoveService(repo *Repository.MoveRepository) *MoveService {
    return &MoveService{repo: repo}
}

// func (this *MoveService) GetMove(id string) ([]Dto.Move, error) {
//     moves, err := this.repo.GetMove(id);
    
//     if err != nil {
//         return nil, err
//     }

//     return moves, nil;
// }

func (this *MoveService) CreateMove(moveCreate Dto.MoveCreate) ([]Dto.Move, error) {
    moves, err := this.repo.CreateMove(moveCreate);
    
    if err != nil {
        return nil, err
    }

    return moves, nil;
}