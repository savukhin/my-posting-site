<script setup lang="ts">
    import { Post, PostItem, PostItemPhoto } from "@/dto/Post";
    import { ref } from "@vue/runtime-dom";
    import Block from "../components/Block.vue";
    import Dropdown from "../components/Dropdown.vue";
    import IconCamera from "../components/icons/IconCamera.vue";
    import IconText from "../components/icons/IconText.vue";
    
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
    
    function photoUpload(event: InputEvent, item: PostItem) {
        const photoItem = item.getContent()
        const target = (event.target as HTMLInputElement)
        if (!(photoItem instanceof PostItemPhoto) || !target || !target.files || target.files.length == 0)
            return
        
        const newPhoto = target.files[0]
        const url = URL.createObjectURL(newPhoto)
        photoItem.photo = newPhoto
        photoItem.photoURL = url
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
                    <div v-else class="post-item">
                        <div class="post-item-photo">
                            <input 
                                type="file" 
                                :id="'photo-'+item.key" 
                                accept="image/png, image/gif, image/jpeg"
                                @change="(event) => {photoUpload(event, item)}"
                            />
    
                            <label :for="'photo-'+item.key">
                                <div class="upload-photo-img">
                                    <strong v-if="!item.itemPhoto.photoURL">+</strong>
                                    <img v-else :src="item.itemPhoto.photoURL" />
                                </div>
                            </label>
                            <input type="text" name="" :id="'title-'+item.key"/>
                        </div>
                        <div class="btn-hollow" @click="() => { changeItemType(item, 'text') }">
                            <IconText/>
                        </div>
                    </div>
                </li>
            </ul>
    
            <button class="btn btn-hollow full-width" @click="createItem"> + </button>
            <br/>
            <br/>
            <button class="btn btn-green full-width"> Submit </button>
        </Block>
    </template>
    
    <style>
    .post-item-photo {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    
    .post-item-photo > input[type='file'] {
        display: none;
    }
    
    .post-item-photo > label {
        margin-bottom: 10px;
    }
    
    .upload-photo-img {
        width: 100px;
        height: 100px;
        background-color: var(--text-link-color);
        font-size: xxx-large;
        color: var(--vt-c-text-dark-2);
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
        border-radius: 13px;
    }
    
    .upload-photo-img > img {
        width: 100%;
        height: 100%;
        object-fit: scale-down;
    }
    
    .upload-photo-img:hover {
        background-color: var(--background-green);
    }
    
    .post-item > .btn-hollow {
        display: flex;
        align-items: center;
        justify-content: center;
    
        padding: 0 2px;
        height: 22px;
    }
    
    .post-item > textarea {
        resize: none;
        overflow: hidden;
    }
    </style>
    