<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps({
    avatarURL: {
        type: String,
    },
    size: {
        type: String,
    }
})

const linearSize = ref("200px")
if (props.size === "small")
    linearSize.value = "20px"

</script>

<template>
    <img 
        v-if="!avatarURL || avatarURL === ''" 
        src="@/img/blank-profile.webp" 
        :style="{height: linearSize, width: linearSize}"
    />
</template>

<style>
    img {
        height: l;
    }
</style>
