package handler

import "context"

type Service interface {
	CreateToken(context.Context, string, string) (string, string, error)
	RefreshToken(context.Context, string, string, string) (string, string, error)
	
}
