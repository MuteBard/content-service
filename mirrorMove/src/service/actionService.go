package mirrorMove

import (
    "time"
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

func (this *ActionService) CreateAction(actionCreate Dto.ActionCreate) ([]Dto.Action, error) {
    
    //create an action object
	newAction := Dto.Action{
		Name: actionCreate.Name,
        Description: actionCreate.Description,
        Seconds: actionCreate.Seconds,
        Token: actionCreate.Token,
        CreatedAt: time.Now().Format(time.RFC3339),
        UpdatedAt: time.Now().Format(time.RFC3339),
	}
    
    actions, err := this.repo.CreateAction(newAction);
    
    if err != nil {
        return nil, err
    }

    return actions, nil;
}

func (this *ActionService) UpdateAction(actionUpdate Dto.ActionUpdate) ([]Dto.Action, error) {
    
    //recreate an existing an action object
	existingAction := Dto.Action {
        Id: actionUpdate.Id,
		Name: actionUpdate.Name,
        Description: actionUpdate.Description,
        Seconds: actionUpdate.Seconds,
        Token: actionUpdate.Token,
        CreatedAt: time.Now().Format(time.RFC3339),
        UpdatedAt: time.Now().Format(time.RFC3339),
	}
    
    
    actions, err := this.repo.UpdateAction(existingAction);
    
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
