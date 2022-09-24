<script setup lang="ts">
import { ref } from "@vue/runtime-dom";
import Block from "../components/Block.vue";
import Dropdown from "../components/Dropdown.vue";
import IconCamera from "../components/icons/IconCamera.vue";
import IconText from "../components/icons/IconText.vue";

class PostItem {
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
}

class PostItemText {
    public text: string = ""
}

class PostItemPhoto {
    public photo?: File
    public title: string = ""
}

class Post {
    items: PostItem[] = []
}

const post = ref(new Post)

function createItemAfter(index: number) {
    let newID = post.value.items.length

    post.value.items =
    [
        ...post.value.items.slice(0, index + 1),
        new PostItem(newID, "text"),
        ...post.value.items.slice(index + 1)
    ]

    return newID
}


function createItem() {
    let newID = post.value.items.length

    post.value.items.push(
        new PostItem(newID, "text")
    )

    return newID
}

function onTextareaChange(event: KeyboardEvent, afterIndex: number) {
    if (event.key === "Enter") {
        event.preventDefault()
        let key = createItemAfter(afterIndex)

        setTimeout(() => {
             document.getElementById("textarea-" + key)?.focus()
        })
        
        return
    }

    const target = event.target as HTMLTextAreaElement
    setTimeout(() => {
        target.style.height = "0px";
        target.style.height = (target.scrollHeight) + 1 + "px";
    }, 0)   
}

function changeItemType(item: PostItem, newType: "text" | "photo") {
    item.type = newType
}

</script>

<template>
    <Block>
        <h1>Create post</h1>

        <ul id="post-content" class="post-ul">
            <li v-for="(item, index) in post.items" :key="item.key">
                <div v-if="item.type == 'text'" class="post-item">
                    <textarea rows="1" @keydown="(event) => { onTextareaChange(event, index) }" :id="'textarea-' + item.key" />
                    <div class="btn-hollow" @click="() => { changeItemType(item, 'photo') }">
                        <IconCamera/>
                    </div>
                </div>
                <div v-else>
                    <div class="btn-hollow" @click="() => { changeItemType(item, 'text') }">
                        <IconText/>
                    </div>
                </div>
            </li>
        </ul>

        <button class="btn btn-hollow full-width" @click="createItem"> + </button>
    </Block>
</template>
