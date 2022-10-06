export interface PostElementDTO {
    text: string
	title: string
	photo_file_name: string
	is_text: boolean
}

export interface PostDTO {
    elements: PostElementDTO[]
}

export class PostItem {
    public key: number
    public type: "text" | "photo"
    private itemText = new PostItemText()
    private itemPhoto = new PostItemPhoto()

    constructor(key: number, type: "text" | "photo") {
        this.key = key
        this.type = type
    }

    public getContent() {
        if (this.type == "photo") 
            return this.itemPhoto
        return this.itemText
    }

    public setText(text: string) {
        this.itemText = new PostItemText(text)
        return this
    }

    public setPhoto(title?: string, photo?: File | string) {
        this.itemPhoto = new PostItemPhoto(title, photo)
        return this
    }

    public setTitle(title: string) {
        this.itemPhoto.title = title
        return this
    }

    public GetPhotoName() {
        return "photo_number_" + this.key
    }

    public getPhoto() {
        return this.itemPhoto.photo
    }

    public toDTO(): PostElementDTO {
        let element: PostElementDTO = { 
            text: this.itemText.text, 
            title: this.itemPhoto.title, 
            is_text: this.type == "text", 
            photo_file_name: this.GetPhotoName()
        } 

        return element 
    }

    public toFormData(formData: FormData, prefix: string): FormData {
        formData.append(prefix + "isText", String(this.type == "text"))

        if (this.type == "text") {
            formData.append(prefix + "text", this.itemText.text)
            
        } else {
            formData.append(prefix + "title", this.itemPhoto.title)

            if (this.itemPhoto.photo)
                formData.append(prefix + "photo", this.itemPhoto.photo)
        }

        return formData
    }
}

export class PostItemText {
    public text: string = ""

    constructor(text?: string) {
        if (text)
            this.text = text
    }
}

export class PostItemPhoto {
    public photo?: File
    public photoURL?: string
    public title: string = ""

    constructor(title?: string, photo?: File | string) {
        if (title)
            this.title = title

        if (photo) {
            if (photo instanceof File) {
                this.photo = photo
                this.photoURL = URL.createObjectURL(this.photo)
            } else { 
                this.photoURL = photo
            }
        }
    }
}

export class Post {
    items: PostItem[] = []

    toDTO(): PostDTO {
        let result: PostDTO = { elements: [] }

        for (let item of this.items) {
            result.elements.push(item.toDTO())
        }

        return result
    }

    toFormData(): FormData {
        let formData = new FormData()
        formData.append("count", this.items.length.toString())

        for (let i = 0; i < this.items.length; i++) {
            formData = this.items[i].toFormData(formData, i + "_")
        }

        return formData
    }
}

