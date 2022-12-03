package clients

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/genproto/article_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/genproto/author_service"
	"google.golang.org/grpc"
)

type GrpcClients struct {
	Article article_service.ArticleServiceClient
	Author  author_service.AuthorServiceClient
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {

	connArticle, err := grpc.Dial(cfg.ArticleServiceGrpcHost+cfg.ArticleServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	articleC := article_service.NewArticleServiceClient(connArticle)

	connAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	authorC := author_service.NewAuthorServiceClient(connAuthor)

	return &GrpcClients{
		Article: articleC,
		Author:  authorC,
	}, nil
}
