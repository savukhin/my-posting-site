package mappers

import (
	"errors"
	kafka_messages "my-posting-site/backend/common/kafka"
	"my-posting-site/backend/common/models"
	pbPost "my-posting-site/backend/common/protobuf/golang/post"
)

func GrpcPostToKafkaPost(req *pbPost.CreatePostRequest, userId, postId int) *kafka_messages.Post {
	post := &kafka_messages.Post{
		UserID: userId,
		ID:     postId,
	}

	for _, elem := range req.Post.Elements {
		kafka_elem := &kafka_messages.PostElement{
			IsText: elem.IsText,
		}

		if elem.IsText {
			kafka_elem.Text = elem.Text
		} else {
			kafka_elem.Title = elem.PhotoTitle
			kafka_elem.Photo = &elem.Photo
		}

		post.Elements = append(post.Elements, *kafka_elem)
	}

	return post
}

func KafkaPostToPost(req *kafka_messages.Post, savePath string) (*models.Post, []*models.File, error) {
	post := &models.Post{
		AuthorID: req.UserID,
	}

	files := make([]*models.File, 0)

	for _, elem := range req.Elements {
		file := &models.File{
			Owner: models.PostOwner,
		}

		if elem.IsText && elem.Text != nil {
			file.Type = models.TextFile
			filename, err := saveText(*elem.Text, savePath)
			if err != nil {
				return nil, nil, err
			}

			file.Path = filename

		} else if !elem.IsText && elem.Photo != nil && elem.Title != nil {
			file.Type = models.PhotoFile
			file.Title = *elem.Title
			filename, err := savePhoto(*elem.Photo, savePath)
			if err != nil {
				return nil, nil, err
			}

			file.Path = filename
		} else {
			return nil, nil, errors.New("error in passed parameters")
		}

		files = append(files, file)
	}

	return post, files, nil
}
