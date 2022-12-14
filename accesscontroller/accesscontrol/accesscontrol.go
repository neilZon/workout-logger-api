package accesscontrol

import (
	"errors"

	"github.com/neilZon/workout-logger-api/accesscontroller"
	"github.com/neilZon/workout-logger-api/database"
	"github.com/neilZon/workout-logger-api/utils"
	"gorm.io/gorm"
)

type AccessController struct {
	DB *gorm.DB
}

// CanAccessExercise implements accesscontroller.AccessControllerService
func (*AccessController) CanAccessExercise(userId string, exerciseId string) error {
	panic("unimplemented")
}

func (ac *AccessController) CanAccessWorkoutRoutine(userId string, workoutRoutineId string) error {
	workoutRoutine, err := database.GetWorkoutRoutine(ac.DB, workoutRoutineId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if utils.UIntToString(workoutRoutine.UserID) != userId {
		return errors.New("Access Denied")
	}
	return nil
}

func (ac *AccessController) CanAccessWorkoutSession(userId string, workoutSessionId string) error {
	workoutSession, err := database.GetWorkoutSession(ac.DB, workoutSessionId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if utils.UIntToString(workoutSession.UserID) != userId {
		return errors.New("Access Denied")
	}
	return nil
}

func (ac *AccessController) CanAccessExerciseRoutine(userId string, exerciseId string) error {
	panic("unimplemented")
}

func (ac *AccessController) CanAccessSetEntry(userId string, exerciseId string) error {
	panic("unimplemented")
}

func NewAccessControllerService(db *gorm.DB) accesscontroller.AccessControllerService {
	return &AccessController{
		DB: db,
	}
}
