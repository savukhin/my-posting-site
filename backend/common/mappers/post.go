package mappers

import (
	"errors"
	"my-posting-site/backend/common/models"
	pbPost "my-posting-site/backend/common/protobuf/golang/post"
	"os"
	"path/filepath"
	"strconv"
)

func saveText(text string, folder string) (string, error) {
	path := filepath.Join("uploads", folder)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	filename := filepath.Join(path, strconv.Itoa(len(files)))

	err = os.WriteFile(filename, []byte(text), os.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func savePhoto(photo []byte, folder string) (string, error) {
	path := filepath.Join("uploads", folder)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	filename := filepath.Join(path, strconv.Itoa(len(files)))

	err = os.WriteFile(filename, photo, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func CreatePostRequestToPost(req *pbPost.CreatePostRequest, userId int, savePath string) (*models.Post, []*models.File, error) {
	post := &models.Post{
		AuthorID: userId,
	}

	files := make([]*models.File, 0)

	for _, elem := range req.Post.Elements {
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

		} else if !elem.IsText && elem.Photo != nil && elem.PhotoTitle != nil {
			file.Type = models.PhotoFile
			file.Title = *elem.PhotoTitle
			filename, err := savePhoto(elem.Photo, savePath)
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

func PostModelToPostResponse(post *models.Post, files []*models.File) (*pbPost.PostResponse, error) {
	if post == nil || files == nil {
		return nil, errors.New("passed params is nil")
	}

	result := &pbPost.PostResponse{
		Success:  true,
		Elements: make([]*pbPost.PostResponseElement, 0),
	}

	for _, file := range files {
		elem := &pbPost.PostResponseElement{}
		if file.Type == models.TextFile {
			elem.IsText = true
			byteArray, err := os.ReadFile(file.Path)
			if err != nil {
				continue
			}
			s := string(byteArray)
			elem.Text = &s

		} else {
			elem.IsText = false
			elem.PhotoTitle = &file.Title
			elem.PhotoUrl = &file.Path
		}

		result.Elements = append(result.Elements, elem)
	}

	return result, nil
}
