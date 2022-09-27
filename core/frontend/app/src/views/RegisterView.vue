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
                            <span>Email:</span>        
                        </td>
                        <td>
                            <input name="email" id="email" placeholder="Email"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Password:</span>
                        </td>
                        <td>
                            <input name="password" type="password" id="password" placeholder="Password"/>
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <span>Repeat password:</span>
                        </td>
                        <td>
                            <input name="password2" type="password" id="password2" placeholder="Password"/>
                        </td>
                    </tr>
                </table>
            </fieldset>
            <p>Already has an account? <RouterLink to="/login">Login</RouterLink> </p>
            <button type="submit" class="btn btn-green">Register</button>
        </form>
    </Block>
</template>
