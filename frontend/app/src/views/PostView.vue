<script setup lang="ts">
// import { Post, PostItem, PostItemPhoto } from "@/dto/GPost";
import { onMounted, ref } from "@vue/runtime-dom";
import { useRoute } from "vue-router";
import Block from "../components/Block.vue";
import Dropdown from "../components/Dropdown.vue";
import IconCamera from "../components/icons/IconCamera.vue";
import IconText from "../components/icons/IconText.vue";
import type { PostResponse } from "@/dto/GetPost";
import { getPost } from "@/hooks/post";
import { ErrorResponse } from "@/dto/defaultResponses";
import { User } from "@/dto/User";

const props = defineProps({
    user: {
        type: User,
        required: false,
    }
})

const post = ref<PostResponse>()
const isFinished = ref(false)
const isLoading = ref(true)
const error = ref<string>()

const route = useRoute()
function notFoundPost() {
    isLoading.value = false
    post.value = undefined
}

const postId = route.params.postId
if (typeof postId != "string") {
    notFoundPost()
} else if (postId == undefined) {
    notFoundPost()
} else {
    getPost(+postId).then(value => {
        console.log(value);
        
        isLoading.value = false

        if (value == undefined) {
            error.value = "This post doesn't exists!"
        } else if (value instanceof ErrorResponse) {
            error.value = ( value.msg == "" ? "Unknown error" : value.msg )
        } else if (!value.finished) {
            isFinished.value = false
        } else {
            post.value = value
            isFinished.value = true
        }
    })
}

</script>

<template>
    <Block v-if="!isLoading && error">
        <h1> {{ error }}</h1>
    </Block>
    <Block v-else-if="!isLoading && !isFinished">
        <h1> Sorry this post is processing now </h1>
    </Block>
    <Block v-else-if="!isLoading && post && isFinished">
        <h1>Post #{{ postId }}</h1>

        <ul id="post-content" class="post-ul">
            <li v-for="(item, index) in post.elements" :key="index">
                <div v-if="item.is_text" class="post-item">
                    <span> {{ item.text }} </span>
                </div>
                <div v-else class="post-item">
                    <div class="post-item-photo">
                        <label :for="'photo-'+index">
                            <div class="upload-photo-img">
                                <!-- <strong v-if="!item.photo_url">+</strong> -->
                                <!-- <img v-else :src="item.itemPhoto.photoURL" /> -->
                                <img :src="'http://127.0.0.1:3000/api/post/get_file/' +item.photo_url" />
                            </div>
                        </label>
                        <p> {{ item.photo_title }} </p>
                    </div>
                </div>
            </li>
        </ul>

    </Block>
</template>

<style>
.post-item-photo {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.post-item-photo > p {
    color: var(--vt-c-text-light-2);
}

.post-item-photo > label {
    margin-bottom: 3px;
}

.upload-photo-img {
    width: 100px;
    height: 100px;
    border: solid 1px var(--text-link-color);
    font-size: xxx-large;
    color: var(--vt-c-text-dark-2);
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 13px;
}

.upload-photo-img > img {
    width: 100%;
    height: 100%;
    object-fit: scale-down;
}

.post-item > .btn-hollow {
    display: flex;
    align-items: center;
    justify-content: center;

    padding: 0 2px;
    height: 22px;
}

</style>
