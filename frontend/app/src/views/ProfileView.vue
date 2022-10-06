<script setup lang="ts">
import { ErrorResponse } from "@/dto/defaultResponses";
import type { Profile } from "@/dto/Profile";
import { User } from "@/dto/User";
import router from "@/router";
import { onMounted, ref } from "@vue/runtime-dom";
import { useRoute } from "vue-router";
import Avatar from "../components/Avatar.vue";
import Block from "../components/Block.vue";
import { getProfile } from "../hooks/user"

const profile = ref<Profile | undefined>(undefined)
const isLoading = ref(true)

const props = defineProps({
    user: {
        type: User,
        required: false,
    }
})

function notFoundUser() {
    isLoading.value = false
    profile.value = undefined
}

const route = useRoute()
const profileId = route.params.profile_id
if (typeof profileId != "string") {
    notFoundUser()
}

if (profileId == undefined) {
    console.log(props.user);
    
    if (props.user != undefined) {
        router.push("/profile/" + props.user.id.toString())
    } else {
        notFoundUser()
    }
}

getProfile(+profileId).then((value) => {
    if (value == undefined || value instanceof ErrorResponse) {
        notFoundUser()
        return
    }

    isLoading.value = false
    profile.value = value
})

</script>

<template>
    <Block class="form-centered-vertically">
        <div v-if="!isLoading && profile != undefined">
            <h1>{{ profile?.username }}</h1>
            <Avatar v-if="!isLoading" :avatarURL="profile?.avatar_url" />
        </div>
        <div v-else-if="!isLoading">
            <h1> No user with this id :( </h1>
        </div>
        <div v-else>
            <h1>Loading...</h1>
        </div>
    </Block>
</template>
