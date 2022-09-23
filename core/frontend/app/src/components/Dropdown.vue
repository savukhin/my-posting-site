<script setup lang="ts">import { onMounted, ref } from '@vue/runtime-dom';

defineProps({
    options: Array<String>
})

const emit = defineEmits<{
    (event: "change", chosen: string) : void
}>()

// const opened = ref(false)
const opened = ref(true)
const dropdownContent = ref<HTMLDivElement>()
const chosen = ref<String>()

function toggleDrowndown() {
    opened.value = !opened.value
}

onMounted(() => {
    if (!dropdownContent.value)
        return

    dropdownContent.value.style.top = -dropdownContent.value.scrollHeight / 2 + 10 + "px" 
})

function chose(option: string) {
    opened.value = !opened.value
    chosen.value = option
    emit("chane", chosen.value)
}

</script>

<template>
    <div>
        <button ref="button" class="btn-dropdown" @click="toggleDrowndown">
            ...
        </button>

        <div class="dropdown-content-wrapper">
            <div ref="dropdownContent" v-if="opened" class="dropdown-content">
                <ul>
                    <li v-for="option in options" class="dropdown-item btn-item">
                        <div @click="() => { chose(option) }"> {{ option }} </div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<style>
.dropdown-content-wrapper {
    position: absolute;
    left: 40px;
    top: 0px;
}

.dropdown-content {
    /* position: relative; */
    /* position: absolute;
    left: 40px;
    top: 0px; */
    background-color: var(--vt-c-black-mute);
    color: var(--vt-c-text-dark-2);

    border-radius: 3px;
    padding: 5px 0 5px 0px;
}

.dropdown-content > ul {
    padding: 0px
}

.dropdown-content li {
    list-style: none;
    /* margin: 0px 20px 0px 5px; */
}

.dropdown-content li > div {
    /* list-style: none; */
    margin: 0px 20px 0px 5px;
}

.dropdown-content::after {
    content: "";
    width: 0;
    height: 0;
    border-style: solid;
    border-width: 4px 8px 4px 0;
    border-width: 5px;
    border-color: transparent var(--vt-c-black-mute) transparent transparent;
    position: absolute;
    margin-top: -50%;
    margin-left: -10px;
}
</style>
