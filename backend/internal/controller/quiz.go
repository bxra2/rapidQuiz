package controller

import (
	"github.com/bxra2/rapidQuiz/internal/service"
	"github.com/gofiber/fiber/v2"
)

type QuizController struct {
	quizService *service.QuizService
}

func Quiz(quizService *service.QuizService) QuizController {
	return QuizController{
		quizService: quizService,
	}
}

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}
	return ctx.JSON(quizzes)
}
