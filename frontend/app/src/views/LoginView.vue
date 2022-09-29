<script setup lang="ts">
import Block from "../components/Block.vue";

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
                            <input name="login" id="login" placeholder="Login"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Password:</span>
                        </td>
                        <td>
                            <input name="password" type="password" id="login" placeholder="Password"/>
                        </td>
                    </tr>
                </table>
            </fieldset>
            <p>Has no account? <RouterLink to="/register">Register</RouterLink> </p>
            <button type="submit" class="btn btn-green">Login</button>
        </form>
    </Block>
</template>
