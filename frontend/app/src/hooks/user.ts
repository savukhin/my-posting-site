import type { Profile } from "@/dto/Profile";
import { ErrorResponse } from "@/dto/defaultResponses";
import axios, { AxiosError } from "axios"

export function getProfile(id: number) {
    return axios.get<Profile>("/api/user/get_profile/" + id)
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
            
            let response = err.response.data as Profile
            if (!response.has_error) {
                return undefined
            }
            return new ErrorResponse(response.msg)
        })
}