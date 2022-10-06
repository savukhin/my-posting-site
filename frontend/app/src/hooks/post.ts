import type { Login, Register } from "@/dto/Auth";
import { GenerateError, type DefaultJWTResponse, ErrorResponse } from "@/dto/defaultResponses";
import type { Post } from "@/dto/Post";
import axios, { AxiosError } from "axios"

export function createPost(post: Post) {
    let token = localStorage.getItem("token");
    if (token == null) {
        return false
    }

    let formData = post.toFormData()
    

    return axios.post<DefaultJWTResponse>("/api/post/create_post", formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
                'Authorization': token
            }
        })
        .then(response => {
            if (response.status != 200) {
                return undefined
            }
            
            return response.data as DefaultJWTResponse
        })
        .catch((err: AxiosError) => {
            if (err.response == undefined) {
                return undefined
            }
            
            let response = err.response.data as DefaultJWTResponse
            if (!response.has_error) {
                return undefined
            }
            return new ErrorResponse(response.msg)
        })
}
