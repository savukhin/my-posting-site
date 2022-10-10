package dto

import pbPost "my-posting-site/backend/common/protobuf/golang/post"

type PostElement struct {
	Text          string `json:"text,omitempty"`
	Title         string `json:"title,omitempty"`
	PhotoFileName string `json:"photo_file_name,omitempty"`
	IsText        bool   `json:"is_text"`
}

type Post struct {
	Elements []PostElement `json:"elements"`
}

type CreatePostResponse struct {
	Has_err bool   `json:"has_error"`
	Error   string `json:"error,omitempty"`
	PostId  int32  `json:"post_id,omitempty"`
}

func FromPostDefaultResponse(src *pbPost.DefaultResponse) *CreatePostResponse {
	return &CreatePostResponse{
		Has_err: !src.Success,
		Error:   src.Error,
		PostId:  src.PostId,
	}
}

func (post *Post) ToProtobuf(photos map[string][]byte) *pbPost.Post {
	pbElements := make([]*pbPost.PostElement, 0)
	for _, elem := range post.Elements {
		pbElem := &pbPost.PostElement{}
		if elem.IsText {
			pbElem.IsText = true
			pbElem.Text = &elem.Text
		} else {
			pbElem.IsText = false
			pbElem.PhotoTitle = &elem.Title
			pbElem.Photo = photos[elem.PhotoFileName]
		}

		pbElements = append(pbElements, pbElem)
	}

	result := &pbPost.Post{
		Elements: pbElements,
	}

	return result
}
