import type { Login, Register } from "@/dto/Auth";
import { GenerateError, type DefaultJWTResponse, ErrorResponse } from "@/dto/DefaultResponses";
import axios, { AxiosError } from "axios"

export function register(username: string, email: string, password: string, password2: string) {
    let request: Register = {username, email, password, password2}

    return axios.post<DefaultJWTResponse>("/api/auth/register", request)
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

export function login(username: string, password: string) {
    let request: Login = {username, password}

    return axios.post<DefaultJWTResponse>("/api/auth/login", request)
        .then(response => {
            if (response.status != 200) {
                return undefined
            }
            let data = response.data as DefaultJWTResponse

            localStorage.setItem("token", data.token)
            
            return data
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