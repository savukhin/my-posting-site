export interface PostResponseElement {
    text: string
    photo_url: string
    photo_title: string
    is_text: boolean
}

export interface PostResponse {
    success: boolean
    error: string
    elements: PostResponseElement[]
}
