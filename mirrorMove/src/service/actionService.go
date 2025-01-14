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

func (this *ActionService) SearchActions(actionApiArgs Dto.ActionApiArguments) ([]Dto.Action, error) {
    actions, err := this.repo.SearchActions(actionApiArgs);
    
    if err != nil {
        return nil, err
    }

    return actions, nil;
}

func (this *ActionService) GetAction(id string) ([]Dto.Action, error) {
    actions, err := this.repo.GetAction(id);
    
    if err != nil {
        return nil, err
    }

    return actions, nil;
}

func (this *ActionService) CreateAction(action Dto.Action) ([]Dto.Action, error) {
    actions, err := this.repo.CreateAction(action);
    
    if err != nil {
        return nil, err
    }

    return actions, nil;
}

func (this *ActionService) HideAction(id string) ([]Dto.Action, error) {
    actions, err := this.repo.HideAction(id);
    
    if err != nil {
        return nil, err
    }

    return actions, nil;
}
