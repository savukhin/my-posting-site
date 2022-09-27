<script setup lang="ts">
import { Post, PostItem, PostItemPhoto } from "@/dto/Post";
import { onMounted, ref } from "@vue/runtime-dom";
import { useRoute } from "vue-router";
import Block from "../components/Block.vue";
import Dropdown from "../components/Dropdown.vue";
import IconCamera from "../components/icons/IconCamera.vue";
import IconText from "../components/icons/IconText.vue";

const route = useRoute()
const postId = route.params.postId

const post = ref(new Post)

setTimeout(() => {
    post.value = new Post()

    post.value.items.push(
        new PostItem(1, "text").setText("Paragraph 1"),
        new PostItem(2, "text").setText("text 1"),
        new PostItem(3, "text").setText("Paragraph 2"),
        new PostItem(4, "photo").setPhoto("Google logo", "https://www.google.ru/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"),
        new PostItem(5, "text").setText("text 2"),
    )
}, 300)

</script>

<template>
    <Block>
        <h1>Post #{{ postId }}</h1>

        <ul id="post-content" class="post-ul">
            <li v-for="(item, index) in post.items" :key="item.key">
                <div v-if="item.type == 'text'" class="post-item">
                    <span> {{ item.itemText.text }} </span>
                </div>
                <div v-else class="post-item">
                    <div class="post-item-photo">
                        <label :for="'photo-'+item.key">
                            <div class="upload-photo-img">
                                <strong v-if="!item.itemPhoto.photoURL">+</strong>
                                <img v-else :src="item.itemPhoto.photoURL" />
                            </div>
                        </label>
                        <p> {{ item.itemPhoto.title }} </p>
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
