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
}
