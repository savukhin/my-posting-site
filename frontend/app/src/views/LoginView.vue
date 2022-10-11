<script setup lang="ts">
import { ErrorResponse } from "@/dto/DefaultResponses";
import { login } from "@/hooks/auth";
import router from "@/router";
import { ref } from "vue";
import Block from "../components/Block.vue";

const username = ref<HTMLInputElement>()
const password = ref<HTMLInputElement>()
const errors = ref<string>()

function loginClick(event: Event) {
    event.preventDefault()
    login(
        username.value!.value,
        password.value!.value,
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
                            <span>Login:</span>
                        </td>
                        <td>
                            <input ref="username" name="username" id="username" placeholder="Username"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Password:</span>
                        </td>
                        <td>
                            <input ref="password" name="password" type="password" id="login" placeholder="Password"/>
                        </td>
                    </tr>
                </table>
            </fieldset>
            <p class="warning" :v-if="errors">{{errors}}</p>
            <p>Has no account? <RouterLink to="/register">Register</RouterLink> </p>
            <button type="submit" class="btn btn-green" @click="loginClick">Login</button>
        </form>
    </Block>
</template>
