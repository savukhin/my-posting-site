import axios from "axios"

export function register(username: string, email: string, password: string, password2: string) {

    return axios.post("/api/auth/register")
        .then(val => {
            console.log(val);
            return 1
        })
        .catch(err => {
            console.log(err);
            return "asdf"
        })
}