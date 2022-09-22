<script setup lang="ts">
import { Profile } from "@/dto/Profile";
import { User } from "@/dto/User";
import { onMounted, ref } from "@vue/runtime-dom";
import Avatar from "../components/Avatar.vue";
import Block from "../components/Block.vue";

const profile = ref(new Profile())
const isLoading = ref(true)

const props = defineProps({
    user: {
        type: User,
        required: true,
    }
})

setTimeout(() => { 
    isLoading.value = false
    profile.value = new Profile()
    profile.value.avatarURL = undefined
    profile.value.id = 0
    profile.value.username = "saveliy"

    console.log("Profile generated")
 }, 1)

</script>

<template>
    <Block class="form-centered-vertically">
        <h1>{{ profile.username }}</h1>
        <Avatar v-if="!isLoading" :avatarURL="profile.avatarURL" />
    </Block>
</template>
