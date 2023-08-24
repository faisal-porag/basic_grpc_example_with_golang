package service_middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc_sample/utils"
	"strings"
)

var allowedTokenTypes = []string{
	"jwt",
	"bearer",
}

func CheckAuthorizationMiddleware(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization metadata not found")
	}

	// Check the authorization token in the metadata
	authTokens := md.Get("authorization")
	if len(authTokens) == 0 {
		return status.Errorf(codes.Unauthenticated, "Authorization token missing")
	}

	// Validate the JWT token
	serviceAuthToken := authTokens[0]
	tokenParts := strings.Split(serviceAuthToken, " ")
	if len(tokenParts) != 2 {
		return status.Errorf(codes.PermissionDenied, "Invalid authorization token format")
	} else {
		if !slices.Contains(allowedTokenTypes, strings.ToLower(tokenParts[0])) {
			return status.Errorf(codes.PermissionDenied, "Invalid authorization token format")
		}
	}

	tokenSecret := utils.JwtSecretKey // Replace with your secret key from .env
	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil || !token.Valid {
		return status.Errorf(codes.PermissionDenied, "Invalid authorization token")
	}

	return nil
}
