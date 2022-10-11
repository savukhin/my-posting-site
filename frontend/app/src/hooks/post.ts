import type { Login, Register } from "@/dto/Auth";
import { GenerateError, type DefaultResponse, ErrorResponse } from "@/dto/DefaultResponses";
import type { PostResponse } from "@/dto/GetPost";
import type { Post } from "@/dto/Post";
import axios, { AxiosError } from "axios"

export function createPost(post: Post) {
    let token = localStorage.getItem("token");
    if (token == null) {
        return false
    }

    let formData = post.toFormData()

    return axios.post<DefaultResponse>("/api/post/create_post", formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
                'Authorization': token
            }
        })
        .then(response => {
            if (response.status != 200) {
                return undefined
            }
            
            return response.data as DefaultResponse
        })
        .catch((err: AxiosError) => {
            if (err.response == undefined) {
                return undefined
            }
            
            let response = err.response.data as DefaultResponse
            if (!response.has_error) {
                return undefined
            }
            return new ErrorResponse(response.msg)
        })
}

export function getPost(id: number) {
    return axios.get<PostResponse>("/api/post/get_post/" + id)
        .then(response => {
            if (response.status != 200) {
                return undefined
            }
            
            return response.data
        })
        .catch((err: AxiosError) => {
            if (err.response == undefined) {
                return undefined
            }
            
            let response = err.response.data as PostResponse
            if (response.success) {
                return undefined
            }
            return new ErrorResponse(response.error)
        })
}
