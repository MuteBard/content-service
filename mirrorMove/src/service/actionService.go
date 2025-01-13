package mirrorMove

import (
    Dto "mirrorMove/src/dtos"
    Repository "mirrorMove/src/repository"
)


type ActionService struct {
    repo *Repository.ActionRepository
}

func NewActionService(repo *Repository.ActionRepository) *ActionService {
    return &ActionService{repo: repo}
}

func (s *ActionService) GetAllActions() ([]Dto.Action, error) {
    // Call the repository method to retrieve all actions
    actions, err := s.repo.GetAllActions();
    
    if err != nil {
        return nil, err
    }

    return actions, nil;
}

func (s *ActionService) GetTest() string {
    return "abc"
}