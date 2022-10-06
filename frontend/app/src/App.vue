<script setup lang="ts">
import { ref } from "vue";
import { RouterLink, RouterView } from "vue-router";
import { ErrorResponse } from "./dto/defaultResponses";
import { User } from "./dto/User";
import { checkJWT } from "./hooks/user";

  const user = ref<User>()

  const isLoading = ref(true)
  const isLogged = ref(false)

  // setTimeout(() => {
  //   isLoading.value = false
  //   isLogged.value = true
  //   user.value = new User()
  //   user.value.avatarURL = ""
  //   user.value.id = 0
  //   user.value.username = "saveliy"

  //   console.log("User generated");
    
  // }, 1)

  function userNotFound() {
    isLoading.value = false
    isLogged.value = false
    user.value = undefined
  }

  let checking = checkJWT()
  if (checking == false) {
    userNotFound()
  } else {
    checking.then(value => {
      if (value == undefined || value instanceof ErrorResponse) {
        userNotFound()
        return
      } 

      isLogged.value = true
      user.value = new User()
      user.value.id = value.id
      
      setTimeout(() => {isLoading.value = false})
      
    })
  }
</script>

<template>
  <div v-if="isLoading"> Loading </div>
  <div v-else>
    <header>
      <img
        alt="Vue logo"
        class="logo"
        src="@/assets/logo.svg"
        width="125"
        height="var(--navbar-height)"
      />

      <div class="wrapper">
        <nav>
          <RouterLink class="btn-hollow" to="/register">Register</RouterLink>
          <RouterLink class="btn-hollow" to="/login">Login</RouterLink>
          <RouterLink class="btn-hollow" to="/profile">Profile</RouterLink>
          <RouterLink class="btn-hollow" to="/create-post">Create Post</RouterLink>
          <RouterLink class="btn-hollow" to="/post/1">Post</RouterLink>
        </nav>
      </div>
    </header>
    <div class="post-header"></div>

    <div id="main-content">
      <RouterView :user="user"/>
    </div>
  </div>

</template>

<style scoped>
header {
  z-index: 4;
  user-select: all;
  line-height: 1.5;

  position: fixed !important;
  height: var(--navbar-height);
  width: 100vw;
  top: 0;
  left: 0;
  display: block;
  
  background-color: var(--color-heading);
}

header .wrapper {
  display: flex;
  place-items: flex-start;
  flex-wrap: wrap;
  height: 100%;
  align-items: center;
}

.post-header {
  user-select: none;
  height: calc(var(--navbar-height) + 20px);
}

.logo {
  /* display: block;
  margin: 0 auto 2rem; */
  display: none;
}

nav {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: center;
    
  width: 100%;
  font-size: 12px;
  text-align: center;
}

nav a.router-link-exact-active {
  /* color: var(--color-text); */
  color: var(--text-link-color-choosen);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  margin: auto 5px auto 5px;
  display: inline-block;
  padding: 0 1rem;
}

@media (min-width: 1024px) {
  .logo {
    margin: 0 2rem 0 0;
  }

  nav {
    text-align: left;
    font-size: 1rem;
  }
}
</style>
