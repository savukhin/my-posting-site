<script setup lang="ts">
import { ref } from "vue";
import Block from "../components/Block.vue";
import { register } from "../hooks/auth"
import { GenerateError, type DefaultJWTResponse, ErrorResponse } from "@/dto/defaultResponses";
import router from "@/router";

const username = ref<HTMLInputElement>()
const email = ref<HTMLInputElement>()
const password = ref<HTMLInputElement>()
const password2 = ref<HTMLInputElement>()
const errors = ref<string>()

function registerClick(event: Event) {
    event.preventDefault()
    register(
        username.value!.value,
        email.value!.value,
        password.value!.value,
        password2.value!.value,
    ).then((msg) => {
        console.log(msg);
        if (msg == undefined)
            return

        if (msg instanceof ErrorResponse) {
            errors.value = (msg as ErrorResponse).msg
            return
        }

        router.push("/profile/1")
    })
}

</script>

<template>
    <Block>
        <form name="reg-form" method="post" class="form-centered-vertically">
            <fieldset form="reg-form">
                <table>
                    <tr>
                        <td>
                            <span>Username:</span>
                        </td>
                        <td>
                            <input ref="username" name="username" id="username" placeholder="Username"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Email:</span>        
                        </td>
                        <td>
                            <input ref="email" name="email" id="email" placeholder="Email"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Password:</span>
                        </td>
                        <td>
                            <input ref="password" name="password" type="password" id="password" placeholder="Password"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Repeat password:</span>
                        </td>
                        <td>
                            <input ref="password2" name="password2" type="password" id="password2" placeholder="Password"/>
                        </td>
                    </tr>
                </table>
            </fieldset>
            <p class="warning" :v-if="errors">{{errors}}</p>
            <p>Already has an account? <RouterLink to="/login">Login</RouterLink> </p>
            <button type="submit" class="btn btn-green" @click="registerClick">Register</button>
        </form>
    </Block>
</template>
