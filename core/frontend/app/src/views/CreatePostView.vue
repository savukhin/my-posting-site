<script setup lang="ts">
import { ref } from "@vue/runtime-dom";
import Block from "../components/Block.vue";
import Dropdown from "../components/Dropdown.vue";

class PostItem {
    public key: number

    constructor(key: number) {
        this.key = key
    }
}

class PostItemText extends PostItem {
    public text: string = ""

    constructor(key: number) {
        super(key)
    }
}

class PostItemPhoto extends PostItem {
    public photo?: File
    public title: string = ""

    constructor(key: number) {
        super(key)
    }
}

class Post {
    items: PostItem[] = []
}

const post = ref(new Post)

function createItem() {
    let newID = post.value.items.length

    post.value.items.push(
        new PostItemText(newID)
    )

    return newID
}

function onTextareaChange(event: KeyboardEvent) {
    if (event.key === "Enter") {
        event.preventDefault()
        let key = createItem()

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

</script>

<template>
    <Block>
        <h1>Create post</h1>

        <ul id="post-content" class="post-ul">
            <li v-for="item in post.items" :key="item.key" class="post-item">
                <textarea rows="1" @keydown="onTextareaChange" :id="'textarea-' + item.key" />
                <!-- <button class="btn btn-hollow">...</button> -->
                <Dropdown :options="['Text', 'Photo']" />
            </li>
        </ul>

        <button class="btn btn-hollow full-width" @click="createItem"> + </button>
    </Block>
</template>
