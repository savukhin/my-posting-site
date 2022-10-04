export class ErrorResponse {
	msg: string
	constructor (msg: string = "") {
		this.msg = msg
	}
}

export function GenerateError(msg: string) {
    let err: ErrorResponse = { msg }
    return err
}

export interface DefaultResponse {
	has_error: boolean
	msg: string
}

export interface DefaultJWTResponse {
	has_error: boolean
	msg: string
	token: string
}
