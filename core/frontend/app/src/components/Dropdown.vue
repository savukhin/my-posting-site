<script setup lang="ts">import { onMounted, ref } from '@vue/runtime-dom';

const props = defineProps({
    options: {
        type: Array<String>,
        required: true
    }
})

let items: {option: String, chosen: boolean}[] = []
props.options.forEach((option) => {
    items.push({ option: option, chosen: false })
})

const emit = defineEmits<{
    (event: "change", chosen: string | undefined) : void
}>()

const opened = ref(false)
const button = ref<HTMLButtonElement>()
const dropdownContent = ref<HTMLDivElement>()
const chosen = ref<string>()

function fixContentPosition() {
    setTimeout( () => {
        if (dropdownContent.value)
            dropdownContent.value.style.top = -dropdownContent.value.scrollHeight / 2 + 10 + "px" 
    })
}

function toggleDrowndown() {
    opened.value = !opened.value
    fixContentPosition()
}

onMounted(() => {
    fixContentPosition()
})

function chose(option: string) {
    opened.value = !opened.value
    chosen.value = option
    emit("change", chosen.value)
}

document.onclick = (event: MouseEvent) => {
    if (!button.value || !dropdownContent.value ) {
        opened.value = false
        return
    }

    console.log(event.target);
    console.log(button.value);
    console.log(dropdownContent.value);
    

    if (event.target != button.value && event.target != dropdownContent.value) {
        opened.value = false
        return
    }
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
                    <li v-for="(item, index) in items" class="dropdown-item btn-item">
                        <div @click="() => { chose(index) }"> {{ item.option }} </div>
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
